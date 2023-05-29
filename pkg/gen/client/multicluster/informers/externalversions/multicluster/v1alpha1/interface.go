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
	internalinterfaces "github.com/flomesh-io/fsm/pkg/gen/client/multicluster/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// GlobalTrafficPolicies returns a GlobalTrafficPolicyInformer.
	GlobalTrafficPolicies() GlobalTrafficPolicyInformer
	// ServiceExports returns a ServiceExportInformer.
	ServiceExports() ServiceExportInformer
	// ServiceImports returns a ServiceImportInformer.
	ServiceImports() ServiceImportInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// GlobalTrafficPolicies returns a GlobalTrafficPolicyInformer.
func (v *version) GlobalTrafficPolicies() GlobalTrafficPolicyInformer {
	return &globalTrafficPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceExports returns a ServiceExportInformer.
func (v *version) ServiceExports() ServiceExportInformer {
	return &serviceExportInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceImports returns a ServiceImportInformer.
func (v *version) ServiceImports() ServiceImportInformer {
	return &serviceImportInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
