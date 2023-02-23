/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "github.com/client-go/applyconfigurations/core/v1"
)

// IngressBackendApplyConfiguration represents an declarative configuration of the IngressBackend type for use
// with apply.
type IngressBackendApplyConfiguration struct {
	Service  *IngressServiceBackendApplyConfiguration            `json:"service,omitempty"`
	Resource *corev1.TypedLocalObjectReferenceApplyConfiguration `json:"resource,omitempty"`
}

// IngressBackendApplyConfiguration constructs an declarative configuration of the IngressBackend type for use with
// apply.
func IngressBackend() *IngressBackendApplyConfiguration {
	return &IngressBackendApplyConfiguration{}
}

// WithService sets the Service field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Service field is set to the value of the last call.
func (b *IngressBackendApplyConfiguration) WithService(value *IngressServiceBackendApplyConfiguration) *IngressBackendApplyConfiguration {
	b.Service = value
	return b
}

// WithResource sets the Resource field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resource field is set to the value of the last call.
func (b *IngressBackendApplyConfiguration) WithResource(value *corev1.TypedLocalObjectReferenceApplyConfiguration) *IngressBackendApplyConfiguration {
	b.Resource = value
	return b
}
