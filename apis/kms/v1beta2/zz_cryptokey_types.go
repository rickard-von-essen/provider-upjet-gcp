// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CryptoKeyInitParameters struct {

	// The resource name of the backend environment associated with all CryptoKeyVersions within this CryptoKey.
	// The resource name is in the format "projects//locations//ekmConnections/*" and only applies to "EXTERNAL_VPC" keys.
	CryptoKeyBackend *string `json:"cryptoKeyBackend,omitempty" tf:"crypto_key_backend,omitempty"`

	// The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
	// If not specified at creation time, the default duration is 30 days.
	DestroyScheduledDuration *string `json:"destroyScheduledDuration,omitempty" tf:"destroy_scheduled_duration,omitempty"`

	// Whether this key may contain imported versions only.
	ImportOnly *bool `json:"importOnly,omitempty" tf:"import_only,omitempty"`

	// Labels with user-defined metadata to apply to this resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The immutable purpose of this CryptoKey. See the
	// purpose reference
	// for possible inputs.
	// Default value is "ENCRYPT_DECRYPT".
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`

	// Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
	// The first rotation will take place after the specified period. The rotation period has
	// the format of a decimal number with up to 9 fractional digits, followed by the
	// letter s (seconds). It must be greater than a day (ie, 86400).
	RotationPeriod *string `json:"rotationPeriod,omitempty" tf:"rotation_period,omitempty"`

	// If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
	// You must use the google_kms_crypto_key_version resource to create a new CryptoKeyVersion
	// or google_kms_key_ring_import_job resource to import the CryptoKeyVersion.
	SkipInitialVersionCreation *bool `json:"skipInitialVersionCreation,omitempty" tf:"skip_initial_version_creation,omitempty"`

	// A template describing settings for new crypto key versions.
	// Structure is documented below.
	VersionTemplate *VersionTemplateInitParameters `json:"versionTemplate,omitempty" tf:"version_template,omitempty"`
}

type CryptoKeyObservation struct {

	// The resource name of the backend environment associated with all CryptoKeyVersions within this CryptoKey.
	// The resource name is in the format "projects//locations//ekmConnections/*" and only applies to "EXTERNAL_VPC" keys.
	CryptoKeyBackend *string `json:"cryptoKeyBackend,omitempty" tf:"crypto_key_backend,omitempty"`

	// The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
	// If not specified at creation time, the default duration is 30 days.
	DestroyScheduledDuration *string `json:"destroyScheduledDuration,omitempty" tf:"destroy_scheduled_duration,omitempty"`

	// for all of the labels present on the resource.
	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// an identifier for the resource with format {{key_ring}}/cryptoKeys/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether this key may contain imported versions only.
	ImportOnly *bool `json:"importOnly,omitempty" tf:"import_only,omitempty"`

	// The KeyRing that this key belongs to.
	// Format: 'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'.
	KeyRing *string `json:"keyRing,omitempty" tf:"key_ring,omitempty"`

	// Labels with user-defined metadata to apply to this resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A copy of the primary CryptoKeyVersion that will be used by cryptoKeys.encrypt when this CryptoKey is given in EncryptRequest.name.
	// Keys with purpose ENCRYPT_DECRYPT may have a primary. For other keys, this field will be unset.
	// Structure is documented below.
	Primary []PrimaryObservation `json:"primary,omitempty" tf:"primary,omitempty"`

	// The immutable purpose of this CryptoKey. See the
	// purpose reference
	// for possible inputs.
	// Default value is "ENCRYPT_DECRYPT".
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`

	// Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
	// The first rotation will take place after the specified period. The rotation period has
	// the format of a decimal number with up to 9 fractional digits, followed by the
	// letter s (seconds). It must be greater than a day (ie, 86400).
	RotationPeriod *string `json:"rotationPeriod,omitempty" tf:"rotation_period,omitempty"`

	// If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
	// You must use the google_kms_crypto_key_version resource to create a new CryptoKeyVersion
	// or google_kms_key_ring_import_job resource to import the CryptoKeyVersion.
	SkipInitialVersionCreation *bool `json:"skipInitialVersionCreation,omitempty" tf:"skip_initial_version_creation,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`

	// A template describing settings for new crypto key versions.
	// Structure is documented below.
	VersionTemplate *VersionTemplateObservation `json:"versionTemplate,omitempty" tf:"version_template,omitempty"`
}

type CryptoKeyParameters struct {

	// The resource name of the backend environment associated with all CryptoKeyVersions within this CryptoKey.
	// The resource name is in the format "projects//locations//ekmConnections/*" and only applies to "EXTERNAL_VPC" keys.
	// +kubebuilder:validation:Optional
	CryptoKeyBackend *string `json:"cryptoKeyBackend,omitempty" tf:"crypto_key_backend,omitempty"`

	// The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
	// If not specified at creation time, the default duration is 30 days.
	// +kubebuilder:validation:Optional
	DestroyScheduledDuration *string `json:"destroyScheduledDuration,omitempty" tf:"destroy_scheduled_duration,omitempty"`

	// Whether this key may contain imported versions only.
	// +kubebuilder:validation:Optional
	ImportOnly *bool `json:"importOnly,omitempty" tf:"import_only,omitempty"`

	// The KeyRing that this key belongs to.
	// Format: 'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta1.KeyRing
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyRing *string `json:"keyRing,omitempty" tf:"key_ring,omitempty"`

	// Reference to a KeyRing in kms to populate keyRing.
	// +kubebuilder:validation:Optional
	KeyRingRef *v1.Reference `json:"keyRingRef,omitempty" tf:"-"`

	// Selector for a KeyRing in kms to populate keyRing.
	// +kubebuilder:validation:Optional
	KeyRingSelector *v1.Selector `json:"keyRingSelector,omitempty" tf:"-"`

	// Labels with user-defined metadata to apply to this resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The immutable purpose of this CryptoKey. See the
	// purpose reference
	// for possible inputs.
	// Default value is "ENCRYPT_DECRYPT".
	// +kubebuilder:validation:Optional
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`

	// Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
	// The first rotation will take place after the specified period. The rotation period has
	// the format of a decimal number with up to 9 fractional digits, followed by the
	// letter s (seconds). It must be greater than a day (ie, 86400).
	// +kubebuilder:validation:Optional
	RotationPeriod *string `json:"rotationPeriod,omitempty" tf:"rotation_period,omitempty"`

	// If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
	// You must use the google_kms_crypto_key_version resource to create a new CryptoKeyVersion
	// or google_kms_key_ring_import_job resource to import the CryptoKeyVersion.
	// +kubebuilder:validation:Optional
	SkipInitialVersionCreation *bool `json:"skipInitialVersionCreation,omitempty" tf:"skip_initial_version_creation,omitempty"`

	// A template describing settings for new crypto key versions.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	VersionTemplate *VersionTemplateParameters `json:"versionTemplate,omitempty" tf:"version_template,omitempty"`
}

type PrimaryInitParameters struct {
}

type PrimaryObservation struct {

	// (Output)
	// The resource name for this CryptoKeyVersion.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Output)
	// The current state of the CryptoKeyVersion.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type PrimaryParameters struct {
}

type VersionTemplateInitParameters struct {

	// The algorithm to use when creating a version based on this template.
	// See the algorithm reference for possible inputs.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// The protection level to use when creating a version based on this template. Possible values include "SOFTWARE", "HSM", "EXTERNAL", "EXTERNAL_VPC". Defaults to "SOFTWARE".
	ProtectionLevel *string `json:"protectionLevel,omitempty" tf:"protection_level,omitempty"`
}

type VersionTemplateObservation struct {

	// The algorithm to use when creating a version based on this template.
	// See the algorithm reference for possible inputs.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// The protection level to use when creating a version based on this template. Possible values include "SOFTWARE", "HSM", "EXTERNAL", "EXTERNAL_VPC". Defaults to "SOFTWARE".
	ProtectionLevel *string `json:"protectionLevel,omitempty" tf:"protection_level,omitempty"`
}

type VersionTemplateParameters struct {

	// The algorithm to use when creating a version based on this template.
	// See the algorithm reference for possible inputs.
	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm" tf:"algorithm,omitempty"`

	// The protection level to use when creating a version based on this template. Possible values include "SOFTWARE", "HSM", "EXTERNAL", "EXTERNAL_VPC". Defaults to "SOFTWARE".
	// +kubebuilder:validation:Optional
	ProtectionLevel *string `json:"protectionLevel,omitempty" tf:"protection_level,omitempty"`
}

// CryptoKeySpec defines the desired state of CryptoKey
type CryptoKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CryptoKeyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider CryptoKeyInitParameters `json:"initProvider,omitempty"`
}

// CryptoKeyStatus defines the observed state of CryptoKey.
type CryptoKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CryptoKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CryptoKey is the Schema for the CryptoKeys API. A
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type CryptoKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CryptoKeySpec   `json:"spec"`
	Status            CryptoKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CryptoKeyList contains a list of CryptoKeys
type CryptoKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CryptoKey `json:"items"`
}

// Repository type metadata.
var (
	CryptoKey_Kind             = "CryptoKey"
	CryptoKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CryptoKey_Kind}.String()
	CryptoKey_KindAPIVersion   = CryptoKey_Kind + "." + CRDGroupVersion.String()
	CryptoKey_GroupVersionKind = CRDGroupVersion.WithKind(CryptoKey_Kind)
)

func init() {
	SchemeBuilder.Register(&CryptoKey{}, &CryptoKeyList{})
}
