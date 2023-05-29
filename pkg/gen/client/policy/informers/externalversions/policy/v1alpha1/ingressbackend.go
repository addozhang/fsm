/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	policyv1alpha1 "github.com/flomesh-io/fsm/pkg/apis/policy/v1alpha1"
	versioned "github.com/flomesh-io/fsm/pkg/gen/client/policy/clientset/versioned"
	internalinterfaces "github.com/flomesh-io/fsm/pkg/gen/client/policy/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/flomesh-io/fsm/pkg/gen/client/policy/listers/policy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// IngressBackendInformer provides access to a shared informer and lister for
// IngressBackends.
type IngressBackendInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.IngressBackendLister
}

type ingressBackendInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewIngressBackendInformer constructs a new informer for IngressBackend type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIngressBackendInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIngressBackendInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredIngressBackendInformer constructs a new informer for IngressBackend type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIngressBackendInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1alpha1().IngressBackends(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1alpha1().IngressBackends(namespace).Watch(context.TODO(), options)
			},
		},
		&policyv1alpha1.IngressBackend{},
		resyncPeriod,
		indexers,
	)
}

func (f *ingressBackendInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIngressBackendInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *ingressBackendInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&policyv1alpha1.IngressBackend{}, f.defaultInformer)
}

func (f *ingressBackendInformer) Lister() v1alpha1.IngressBackendLister {
	return v1alpha1.NewIngressBackendLister(f.Informer().GetIndexer())
}
