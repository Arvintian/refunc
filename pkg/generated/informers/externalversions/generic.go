/*
Copyright 2018 The refunc Authors

TODO: choose a opensource licence.
*/

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1beta3 "github.com/refunc/refunc/pkg/apis/refunc/v1beta3"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=refunc.refunc.io, Version=v1beta3
	case v1beta3.SchemeGroupVersion.WithResource("funcdeves"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Refunc().V1beta3().Funcdeves().Informer()}, nil
	case v1beta3.SchemeGroupVersion.WithResource("funcinsts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Refunc().V1beta3().Funcinsts().Informer()}, nil
	case v1beta3.SchemeGroupVersion.WithResource("triggers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Refunc().V1beta3().Triggers().Informer()}, nil
	case v1beta3.SchemeGroupVersion.WithResource("xenvs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Refunc().V1beta3().Xenvs().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
