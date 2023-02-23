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

package v1alpha1

import (
	networkingv1alpha1 "k8s.io/api/networking/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	internal "github.com/client-go/applyconfigurations/internal"
	v1 "github.com/client-go/applyconfigurations/meta/v1"
)

// ClusterCIDRApplyConfiguration represents an declarative configuration of the ClusterCIDR type for use
// with apply.
type ClusterCIDRApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *ClusterCIDRSpecApplyConfiguration `json:"spec,omitempty"`
}

// ClusterCIDR constructs an declarative configuration of the ClusterCIDR type for use with
// apply.
func ClusterCIDR(name string) *ClusterCIDRApplyConfiguration {
	b := &ClusterCIDRApplyConfiguration{}
	b.WithName(name)
	b.WithKind("ClusterCIDR")
	b.WithAPIVersion("networking.k8s.io/v1alpha1")
	return b
}

// ExtractClusterCIDR extracts the applied configuration owned by fieldManager from
// clusterCIDR. If no managedFields are found in clusterCIDR for fieldManager, a
// ClusterCIDRApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// clusterCIDR must be a unmodified ClusterCIDR API object that was retrieved from the Kubernetes API.
// ExtractClusterCIDR provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractClusterCIDR(clusterCIDR *networkingv1alpha1.ClusterCIDR, fieldManager string) (*ClusterCIDRApplyConfiguration, error) {
	return extractClusterCIDR(clusterCIDR, fieldManager, "")
}

// ExtractClusterCIDRStatus is the same as ExtractClusterCIDR except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractClusterCIDRStatus(clusterCIDR *networkingv1alpha1.ClusterCIDR, fieldManager string) (*ClusterCIDRApplyConfiguration, error) {
	return extractClusterCIDR(clusterCIDR, fieldManager, "status")
}

func extractClusterCIDR(clusterCIDR *networkingv1alpha1.ClusterCIDR, fieldManager string, subresource string) (*ClusterCIDRApplyConfiguration, error) {
	b := &ClusterCIDRApplyConfiguration{}
	err := managedfields.ExtractInto(clusterCIDR, internal.Parser().Type("io.k8s.api.networking.v1alpha1.ClusterCIDR"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(clusterCIDR.Name)

	b.WithKind("ClusterCIDR")
	b.WithAPIVersion("networking.k8s.io/v1alpha1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ClusterCIDRApplyConfiguration) WithKind(value string) *ClusterCIDRApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ClusterCIDRApplyConfiguration) WithAPIVersion(value string) *ClusterCIDRApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ClusterCIDRApplyConfiguration) WithName(value string) *ClusterCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *ClusterCIDRApplyConfiguration) WithGenerateName(value string) *ClusterCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ClusterCIDRApplyConfiguration) WithNamespace(value string) *ClusterCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ClusterCIDRApplyConfiguration) WithUID(value types.UID) *ClusterCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ClusterCIDRApplyConfiguration) WithResourceVersion(value string) *ClusterCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *ClusterCIDRApplyConfiguration) WithGeneration(value int64) *ClusterCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *ClusterCIDRApplyConfiguration) WithCreationTimestamp(value metav1.Time) *ClusterCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *ClusterCIDRApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *ClusterCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *ClusterCIDRApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *ClusterCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ClusterCIDRApplyConfiguration) WithLabels(entries map[string]string) *ClusterCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *ClusterCIDRApplyConfiguration) WithAnnotations(entries map[string]string) *ClusterCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *ClusterCIDRApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *ClusterCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *ClusterCIDRApplyConfiguration) WithFinalizers(values ...string) *ClusterCIDRApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *ClusterCIDRApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *ClusterCIDRApplyConfiguration) WithSpec(value *ClusterCIDRSpecApplyConfiguration) *ClusterCIDRApplyConfiguration {
	b.Spec = value
	return b
}
