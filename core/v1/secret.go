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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	internal "github.com/ecordell/applyconfigurations/internal"
	v1 "github.com/ecordell/applyconfigurations/meta/v1"
)

// SecretApplyConfiguration represents an declarative configuration of the Secret type for use
// with apply.
type SecretApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Immutable                        *bool              `json:"immutable,omitempty"`
	Data                             map[string][]byte  `json:"data,omitempty"`
	StringData                       map[string]string  `json:"stringData,omitempty"`
	Type                             *corev1.SecretType `json:"type,omitempty"`
}

// Secret constructs an declarative configuration of the Secret type for use with
// apply.
func Secret(name, namespace string) *SecretApplyConfiguration {
	b := &SecretApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("Secret")
	b.WithAPIVersion("v1")
	return b
}

// ExtractSecret extracts the applied configuration owned by fieldManager from
// secret. If no managedFields are found in secret for fieldManager, a
// SecretApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// secret must be a unmodified Secret API object that was retrieved from the Kubernetes API.
// ExtractSecret provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractSecret(secret *corev1.Secret, fieldManager string) (*SecretApplyConfiguration, error) {
	return extractSecret(secret, fieldManager, "")
}

// ExtractSecretStatus is the same as ExtractSecret except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractSecretStatus(secret *corev1.Secret, fieldManager string) (*SecretApplyConfiguration, error) {
	return extractSecret(secret, fieldManager, "status")
}

func extractSecret(secret *corev1.Secret, fieldManager string, subresource string) (*SecretApplyConfiguration, error) {
	b := &SecretApplyConfiguration{}
	err := managedfields.ExtractInto(secret, internal.Parser().Type("io.k8s.api.core.v1.Secret"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(secret.Name)
	b.WithNamespace(secret.Namespace)

	b.WithKind("Secret")
	b.WithAPIVersion("v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *SecretApplyConfiguration) WithKind(value string) *SecretApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *SecretApplyConfiguration) WithAPIVersion(value string) *SecretApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *SecretApplyConfiguration) WithName(value string) *SecretApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *SecretApplyConfiguration) WithGenerateName(value string) *SecretApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *SecretApplyConfiguration) WithNamespace(value string) *SecretApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *SecretApplyConfiguration) WithUID(value types.UID) *SecretApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *SecretApplyConfiguration) WithResourceVersion(value string) *SecretApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *SecretApplyConfiguration) WithGeneration(value int64) *SecretApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *SecretApplyConfiguration) WithCreationTimestamp(value metav1.Time) *SecretApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *SecretApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *SecretApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *SecretApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *SecretApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *SecretApplyConfiguration) WithLabels(entries map[string]string) *SecretApplyConfiguration {
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
func (b *SecretApplyConfiguration) WithAnnotations(entries map[string]string) *SecretApplyConfiguration {
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
func (b *SecretApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *SecretApplyConfiguration {
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
func (b *SecretApplyConfiguration) WithFinalizers(values ...string) *SecretApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *SecretApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithImmutable sets the Immutable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Immutable field is set to the value of the last call.
func (b *SecretApplyConfiguration) WithImmutable(value bool) *SecretApplyConfiguration {
	b.Immutable = &value
	return b
}

// WithData puts the entries into the Data field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Data field,
// overwriting an existing map entries in Data field with the same key.
func (b *SecretApplyConfiguration) WithData(entries map[string][]byte) *SecretApplyConfiguration {
	if b.Data == nil && len(entries) > 0 {
		b.Data = make(map[string][]byte, len(entries))
	}
	for k, v := range entries {
		b.Data[k] = v
	}
	return b
}

// WithStringData puts the entries into the StringData field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the StringData field,
// overwriting an existing map entries in StringData field with the same key.
func (b *SecretApplyConfiguration) WithStringData(entries map[string]string) *SecretApplyConfiguration {
	if b.StringData == nil && len(entries) > 0 {
		b.StringData = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.StringData[k] = v
	}
	return b
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *SecretApplyConfiguration) WithType(value corev1.SecretType) *SecretApplyConfiguration {
	b.Type = &value
	return b
}
