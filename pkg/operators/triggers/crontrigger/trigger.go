package crontrigger

import (
	"fmt"
	"sort"
	"sync"
	"time"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"

	nats "github.com/nats-io/go-nats"
	observer "github.com/refunc/go-observer"
	rfv1beta3 "github.com/refunc/refunc/pkg/apis/refunc/v1beta3"
	refunc "github.com/refunc/refunc/pkg/generated/clientset/versioned"
	informers "github.com/refunc/refunc/pkg/generated/informers/externalversions"
	operators "github.com/refunc/refunc/pkg/operators"
	"github.com/refunc/refunc/pkg/utils"
)

// Operator creates http based rpc trigger automatically for every funcs
type Operator struct {
	*operators.BaseOperator

	Offset time.Duration

	triggers sync.Map

	liveTasks operators.LiveTaskStore

	scheduled observer.Property
}

const (
	// Type name for rpc trigger
	Type = "crontrigger"

	triggerPrifix    = "cron-"
	labelAutoCreated = "crontrigger.refunc.io/auto-created"
)

// NewOperator creates a new rpc trigger operator
func NewOperator(
	cfg *rest.Config,
	conn *nats.Conn,
	rclient refunc.Interface,
	rfInformers informers.SharedInformerFactory,
) (*Operator, error) {
	base, err := operators.NewBaseOperator(cfg, conn, rclient, rfInformers)
	if err != nil {
		return nil, err
	}

	r := &Operator{
		BaseOperator: base,
		liveTasks:    operators.NewLiveTaskStore(),
		scheduled:    observer.NewProperty(nil),
	}

	return r, nil
}

// Run will not return until stopC is closed.
func (r *Operator) Run(stopC <-chan struct{}) {
	defer func() {
		if re := recover(); re != nil {
			utils.LogTraceback(re, 4, klog.V(1))
		}
	}()

	// add events emitter
	r.RefuncInformers.Refunc().V1beta3().Triggers().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    r.handleTriggerAdd,
		UpdateFunc: r.handleTriggerUpdate,
		DeleteFunc: r.handleTriggerDelete,
	})

	klog.Info("(crontrigger) waiting for listers to be fully synced")
	go r.ioLoop(stopC, r.scheduled.Observe())

	if !r.BaseOperator.WaitForCacheSync(stopC) {
		klog.Error("(crontrigger) cannot fully sync resources")
		return
	}

	<-stopC
	klog.Info("(crontrigger) shuting down cron trigger operator")
}

func (r *Operator) handleTriggerAdd(o interface{}) {
	trigger := o.(*rfv1beta3.Trigger)
	if trigger.Spec.Type != Type {
		// skip other triggers
		return
	}

	key := k8sKey(trigger)
	if trigger.Spec.Cron == nil {
		klog.Errorf("(crontrigger) %s cron is empty", key)
		return
	}
	_, loaded := r.triggers.LoadOrStore(key, &cronHandler{
		trKey:    key,
		ns:       trigger.Namespace,
		name:     trigger.Name,
		operator: r,
	})
	if !loaded {
		r.scheduleTasks()
	}
}

func (r *Operator) handleTriggerUpdate(oldObj, curObj interface{}) {
	old := oldObj.(*rfv1beta3.Trigger)
	cur := curObj.(*rfv1beta3.Trigger)

	// Periodic resync may resend the deployment without changes in-between.
	// Also breaks loops created by updating the resource ourselves.
	if old.ResourceVersion == cur.ResourceVersion {
		return
	}
	if old.Spec.Cron != nil && cur.Spec.Cron == nil {
		r.handleTriggerDelete(cur)
		return
	}
	if old.Spec.Cron == nil && cur.Spec.Cron != nil {
		r.handleTriggerAdd(cur)
		return
	}
	if (old.Spec.Cron != nil && cur.Spec.Cron != nil) && (old.Spec.Cron.Cron != cur.Spec.Cron.Cron) {
		// delete and reschedule
		key := k8sKey(cur)
		klog.V(3).Infof("(crontrigger) updating trigger %s", key)
		r.triggers.Delete(key)
		r.handleTriggerAdd(cur)
	}
}

func (r *Operator) handleTriggerDelete(o interface{}) {
	trigger, ok := o.(*rfv1beta3.Trigger)
	if !ok {
		// it's cache.DeletedFinalStateUnknown
		return
	}
	if trigger.Spec.Type != Type {
		// skip other triggers
		return
	}

	key := k8sKey(trigger)
	if _, ok := r.triggers.Load(key); ok {
		klog.V(3).Infof("(crontrigger) deleting trigger %s", key)
		r.triggers.Delete(key)
		r.scheduleTasks()
	}
}

type timeKeyPair struct {
	t   time.Time
	key string
}

func (tkp *timeKeyPair) String() string {
	return fmt.Sprintf("%s@%s", tkp.key, tkp.t.Format(time.RFC3339))
}

func (r *Operator) scheduleTasks() {
	var tkps []*timeKeyPair
	now := time.Now()
	now = now.Add(r.Offset)
	r.triggers.Range(func(k, v interface{}) bool {
		key := k.(string)
		h := v.(*cronHandler)
		next, err := h.Next(now)
		if err != nil {
			klog.Errorf("(crontrigger) %s failed to schedule, %v", key, err)
			return true
		}
		tkps = append(tkps, &timeKeyPair{next.Add(-r.Offset), key})
		return true
	})
	if len(tkps) > 0 {
		// sort by trigger time
		sort.Slice(tkps, func(i int, j int) bool {
			return tkps[i].t.Before(tkps[j].t)
		})
		klog.V(4).Infof("(crontrigger) schduled %v", tkps)
		r.scheduled.Update(tkps)
	}
}

func (r *Operator) ioLoop(stopC <-chan struct{}, events observer.Stream) {

	klog.Info("(crontrigger) ioloop started")
	defer klog.Info("(crontrigger) ioloop exited")

	blockC := make(chan time.Time)
	defer close(blockC) // avoid leak

	var (
		tickC          <-chan time.Time = blockC
		scheduledTasks []*timeKeyPair
		next           time.Time
	)

	for {
		select {
		case <-stopC:
			return
		case <-events.Changes():
			for events.HasNext() {
				events.Next()
			}
			val := events.Value()
			if val == nil {
				klog.Warning("(crontrigger) got empty updates")
				return
			}

			scheduledTasks = val.([]*timeKeyPair)
			// peek top
			first := scheduledTasks[0]
			delta := first.t.Sub(next)
			if delta == 0 {
				// no need update current ticker
				continue
			}
			// original trigger was deleted or updated
			if dur := first.t.Sub(time.Now()); dur > 0 {
				next = first.t
				tickC = time.After(dur)
				klog.Infof("(crontrigger) next %s@%v", first.key, next.Format(time.RFC3339))
				continue
			}

		case <-tickC:
		}

		now := time.Now()
		for _, tkp := range scheduledTasks {
			val, ok := r.triggers.Load(tkp.key)
			if !ok {
				continue
			}
			h := val.(*cronHandler)

			delta := now.Sub(tkp.t)
			if delta > 30*time.Second {
				klog.Warningf("(crontrigger) %s missed trigger, want %v, current %v", h.trKey, tkp.t, now.Truncate(time.Second))
				continue
			}
			if delta < 0 {
				break
			}

			// delta <= time.Minute, tirgger one task
			val.(*cronHandler).Run(tkp.t)
		}
		// the pendding tasks not ready, re-schedule
		r.scheduleTasks()
	}
}

func (r *Operator) getTriggerLister() (trlsiter func(labels.Selector) ([]*rfv1beta3.Trigger, error)) {
	if r.Namespace == apiv1.NamespaceAll {
		trlsiter = r.TriggerLister.List
		return
	}
	trlsiter = r.TriggerLister.Triggers(r.Namespace).List
	return
}

func k8sKey(o metav1.Object) string {
	return o.GetNamespace() + "/" + o.GetName()
}

func retryOnceOnError(fn func() error) error {
	for i := 0; ; i++ {
		err := fn()
		if err != nil {
			if i >= operators.MaxRetries {
				return err
			}
			continue
		}
		return nil
	}
}
