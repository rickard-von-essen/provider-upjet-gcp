/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BootDiskInitializeParamsObservation struct {
}

type BootDiskInitializeParamsParameters struct {

	// The image from which this disk was initialised.
	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// A set of key/value label pairs assigned to the disk.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The size of the image in gigabytes.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The Google Compute Engine disk type. Such as pd-standard, pd-ssd or pd-balanced.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type InstanceFromTemplateAdvancedMachineFeaturesObservation struct {
}

type InstanceFromTemplateAdvancedMachineFeaturesParameters struct {

	// Whether to enable nested virtualization or not.
	// +kubebuilder:validation:Optional
	EnableNestedVirtualization *bool `json:"enableNestedVirtualization,omitempty" tf:"enable_nested_virtualization,omitempty"`

	// The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.
	// +kubebuilder:validation:Optional
	ThreadsPerCore *float64 `json:"threadsPerCore,omitempty" tf:"threads_per_core,omitempty"`
}

type InstanceFromTemplateAttachedDiskObservation struct {
}

type InstanceFromTemplateAttachedDiskParameters struct {

	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name"`

	// +kubebuilder:validation:Optional
	DiskEncryptionKeyRaw *string `json:"diskEncryptionKeyRaw,omitempty" tf:"disk_encryption_key_raw"`

	// +kubebuilder:validation:Optional
	DiskEncryptionKeySha256 *string `json:"diskEncryptionKeySha256,omitempty" tf:"disk_encryption_key_sha256"`

	// +kubebuilder:validation:Optional
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link"`

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode"`

	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source"`
}

type InstanceFromTemplateBootDiskObservation struct {

	// The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource.
	DiskEncryptionKeySha256 *string `json:"diskEncryptionKeySha256,omitempty" tf:"disk_encryption_key_sha256,omitempty"`
}

type InstanceFromTemplateBootDiskParameters struct {

	// Whether the disk will be auto-deleted when the instance is deleted.
	// +kubebuilder:validation:Optional
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`

	// Name with which attached disk will be accessible under /dev/disk/by-id/
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// A 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set.
	// +kubebuilder:validation:Optional
	DiskEncryptionKeyRawSecretRef *v1.SecretKeySelector `json:"diskEncryptionKeyRawSecretRef,omitempty" tf:"-"`

	// Parameters with which a disk was created alongside the instance.
	// +kubebuilder:validation:Optional
	InitializeParams []BootDiskInitializeParamsParameters `json:"initializeParams,omitempty" tf:"initialize_params,omitempty"`

	// The self_link of the encryption key that is stored in Google Cloud KMS to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set.
	// +kubebuilder:validation:Optional
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`

	// Read/write mode for the disk. One of "READ_ONLY" or "READ_WRITE".
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name or self_link of the disk attached to this instance.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type InstanceFromTemplateConfidentialInstanceConfigObservation struct {
}

type InstanceFromTemplateConfidentialInstanceConfigParameters struct {

	// Defines whether the instance should have confidential compute enabled.
	// +kubebuilder:validation:Required
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute" tf:"enable_confidential_compute,omitempty"`
}

type InstanceFromTemplateGuestAcceleratorObservation struct {
}

type InstanceFromTemplateGuestAcceleratorParameters struct {

	// +kubebuilder:validation:Optional
	Count *float64 `json:"count,omitempty" tf:"count"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type InstanceFromTemplateNetworkInterfaceObservation struct {

	// An array of IPv6 access configurations for this interface. Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig specified, then this instance will have no external IPv6 Internet access.
	// +kubebuilder:validation:Optional
	IPv6AccessConfig []NetworkInterfaceIPv6AccessConfigObservation `json:"ipv6AccessConfig,omitempty" tf:"ipv6_access_config,omitempty"`

	// One of EXTERNAL, INTERNAL to indicate whether the IP can be accessed from the Internet. This field is always inherited from its subnetwork.
	IPv6AccessType *string `json:"ipv6AccessType,omitempty" tf:"ipv6_access_type,omitempty"`

	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	// The name of the interface
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type InstanceFromTemplateNetworkInterfaceParameters struct {

	// Access configurations, i.e. IPs via which this instance can be accessed via the Internet.
	// +kubebuilder:validation:Optional
	AccessConfig []NetworkInterfaceAccessConfigParameters `json:"accessConfig,omitempty" tf:"access_config,omitempty"`

	// An array of alias IP ranges for this network interface.
	// +kubebuilder:validation:Optional
	AliasIPRange []NetworkInterfaceAliasIPRangeParameters `json:"aliasIpRange,omitempty" tf:"alias_ip_range,omitempty"`

	// An array of IPv6 access configurations for this interface. Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig specified, then this instance will have no external IPv6 Internet access.
	// +kubebuilder:validation:Optional
	IPv6AccessConfig []NetworkInterfaceIPv6AccessConfigParameters `json:"ipv6AccessConfig,omitempty" tf:"ipv6_access_config,omitempty"`

	// The name or self_link of the network attached to this interface.
	// +crossplane:generate:reference:type=Network
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The private IP address assigned to the instance.
	// +kubebuilder:validation:Optional
	NetworkIP *string `json:"networkIp,omitempty" tf:"network_ip,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The type of vNIC to be used on this interface. Possible values:GVNIC, VIRTIO_NET
	// +kubebuilder:validation:Optional
	NicType *string `json:"nicType,omitempty" tf:"nic_type,omitempty"`

	// The networking queue count that's specified by users for the network interface. Both Rx and Tx queues will be set to this number. It will be empty if not specified.
	// +kubebuilder:validation:Optional
	QueueCount *float64 `json:"queueCount,omitempty" tf:"queue_count,omitempty"`

	// The stack type for this network interface to identify whether the IPv6 feature is enabled or not. If not specified, IPV4_ONLY will be used.
	// +kubebuilder:validation:Optional
	StackType *string `json:"stackType,omitempty" tf:"stack_type,omitempty"`

	// The name or self_link of the subnetwork attached to this interface.
	// +crossplane:generate:reference:type=Subnetwork
	// +kubebuilder:validation:Optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// The project in which the subnetwork belongs.
	// +kubebuilder:validation:Optional
	SubnetworkProject *string `json:"subnetworkProject,omitempty" tf:"subnetwork_project,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetworkRef *v1.Reference `json:"subnetworkRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetworkSelector *v1.Selector `json:"subnetworkSelector,omitempty" tf:"-"`
}

type InstanceFromTemplateObservation struct {

	// The boot disk for the instance.
	// +kubebuilder:validation:Optional
	BootDisk []InstanceFromTemplateBootDiskObservation `json:"bootDisk,omitempty" tf:"boot_disk,omitempty"`

	// The CPU platform used by this instance.
	CPUPlatform *string `json:"cpuPlatform,omitempty" tf:"cpu_platform,omitempty"`

	// Current status of the instance.
	CurrentStatus *string `json:"currentStatus,omitempty" tf:"current_status,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The server-assigned unique identifier of this instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The unique fingerprint of the labels.
	LabelFingerprint *string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty"`

	// The unique fingerprint of the metadata.
	MetadataFingerprint *string `json:"metadataFingerprint,omitempty" tf:"metadata_fingerprint,omitempty"`

	// The networks attached to the instance.
	// +kubebuilder:validation:Optional
	NetworkInterface []InstanceFromTemplateNetworkInterfaceObservation `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// The unique fingerprint of the tags.
	TagsFingerprint *string `json:"tagsFingerprint,omitempty" tf:"tags_fingerprint,omitempty"`
}

type InstanceFromTemplateParameters struct {

	// Controls for advanced machine-related behavior features.
	// +kubebuilder:validation:Optional
	AdvancedMachineFeatures []InstanceFromTemplateAdvancedMachineFeaturesParameters `json:"advancedMachineFeatures,omitempty" tf:"advanced_machine_features,omitempty"`

	// If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires stopping the instance without setting this field, the update will fail.
	// +kubebuilder:validation:Optional
	AllowStoppingForUpdate *bool `json:"allowStoppingForUpdate,omitempty" tf:"allow_stopping_for_update,omitempty"`

	// List of disks attached to the instance
	// +kubebuilder:validation:Optional
	AttachedDisk []InstanceFromTemplateAttachedDiskParameters `json:"attachedDisk,omitempty" tf:"attached_disk,omitempty"`

	// The boot disk for the instance.
	// +kubebuilder:validation:Optional
	BootDisk []InstanceFromTemplateBootDiskParameters `json:"bootDisk,omitempty" tf:"boot_disk,omitempty"`

	// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
	// +kubebuilder:validation:Optional
	CanIPForward *bool `json:"canIpForward,omitempty" tf:"can_ip_forward,omitempty"`

	// The Confidential VM config being used by the instance.  on_host_maintenance has to be set to TERMINATE or this will fail to create.
	// +kubebuilder:validation:Optional
	ConfidentialInstanceConfig []InstanceFromTemplateConfidentialInstanceConfigParameters `json:"confidentialInstanceConfig,omitempty" tf:"confidential_instance_config,omitempty"`

	// Whether deletion protection is enabled on this instance.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// A brief description of the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Desired status of the instance. Either "RUNNING" or "TERMINATED".
	// +kubebuilder:validation:Optional
	DesiredStatus *string `json:"desiredStatus,omitempty" tf:"desired_status,omitempty"`

	// Whether the instance has virtual displays enabled.
	// +kubebuilder:validation:Optional
	EnableDisplay *bool `json:"enableDisplay,omitempty" tf:"enable_display,omitempty"`

	// List of the type and count of accelerator cards attached to the instance.
	// +kubebuilder:validation:Optional
	GuestAccelerator []InstanceFromTemplateGuestAcceleratorParameters `json:"guestAccelerator,omitempty" tf:"guest_accelerator,omitempty"`

	// A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// A set of key/value label pairs assigned to the instance.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The machine type to create.
	// +kubebuilder:validation:Optional
	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	// Metadata key/value pairs made available within the instance.
	// +kubebuilder:validation:Optional
	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Metadata startup scripts made available within the instance.
	// +kubebuilder:validation:Optional
	MetadataStartupScript *string `json:"metadataStartupScript,omitempty" tf:"metadata_startup_script,omitempty"`

	// The minimum CPU platform specified for the VM instance.
	// +kubebuilder:validation:Optional
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty" tf:"min_cpu_platform,omitempty"`

	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	// The name of the instance. One of name or self_link must be provided.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The networks attached to the instance.
	// +kubebuilder:validation:Optional
	NetworkInterface []InstanceFromTemplateNetworkInterfaceParameters `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither self_link nor project are provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Specifies the reservations that this instance can consume from.
	// +kubebuilder:validation:Optional
	ReservationAffinity []InstanceFromTemplateReservationAffinityParameters `json:"reservationAffinity,omitempty" tf:"reservation_affinity,omitempty"`

	// A list of short names or self_links of resource policies to attach to the instance. Currently a max of 1 resource policy is supported.
	// +kubebuilder:validation:Optional
	ResourcePolicies []*string `json:"resourcePolicies,omitempty" tf:"resource_policies,omitempty"`

	// The scheduling strategy being used by the instance.
	// +kubebuilder:validation:Optional
	Scheduling []InstanceFromTemplateSchedulingParameters `json:"scheduling,omitempty" tf:"scheduling,omitempty"`

	// The scratch disks attached to the instance.
	// +kubebuilder:validation:Optional
	ScratchDisk []InstanceFromTemplateScratchDiskParameters `json:"scratchDisk,omitempty" tf:"scratch_disk,omitempty"`

	// The service account to attach to the instance.
	// +kubebuilder:validation:Optional
	ServiceAccount []InstanceFromTemplateServiceAccountParameters `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// The shielded vm config being used by the instance.
	// +kubebuilder:validation:Optional
	ShieldedInstanceConfig []InstanceFromTemplateShieldedInstanceConfigParameters `json:"shieldedInstanceConfig,omitempty" tf:"shielded_instance_config,omitempty"`

	// Name or self link of an instance
	// template to create the instance based on.
	// Name or self link of an instance template to create the instance based on.
	// +crossplane:generate:reference:type=InstanceTemplate
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-gcp/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SourceInstanceTemplate *string `json:"sourceInstanceTemplate,omitempty" tf:"source_instance_template,omitempty"`

	// +kubebuilder:validation:Optional
	SourceInstanceTemplateRef *v1.Reference `json:"sourceInstanceTemplateRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SourceInstanceTemplateSelector *v1.Selector `json:"sourceInstanceTemplateSelector,omitempty" tf:"-"`

	// The list of tags attached to the instance.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The zone that the machine should be created in. If not
	// set, the provider zone is used.
	// The zone of the instance. If self_link is provided, this value is ignored. If neither self_link nor zone are provided, the provider zone is used.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type InstanceFromTemplateReservationAffinityObservation struct {
}

type InstanceFromTemplateReservationAffinityParameters struct {

	// Specifies the label selector for the reservation to use.
	// +kubebuilder:validation:Optional
	SpecificReservation []ReservationAffinitySpecificReservationParameters `json:"specificReservation,omitempty" tf:"specific_reservation,omitempty"`

	// The type of reservation from which this instance can consume resources.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type InstanceFromTemplateSchedulingObservation struct {
}

type InstanceFromTemplateSchedulingParameters struct {

	// Specifies if the instance should be restarted if it was terminated by Compute Engine (not a user).
	// +kubebuilder:validation:Optional
	AutomaticRestart *bool `json:"automaticRestart,omitempty" tf:"automatic_restart,omitempty"`

	// +kubebuilder:validation:Optional
	MinNodeCpus *float64 `json:"minNodeCpus,omitempty" tf:"min_node_cpus,omitempty"`

	// Specifies node affinities or anti-affinities to determine which sole-tenant nodes your instances and managed instance groups will use as host systems.
	// +kubebuilder:validation:Optional
	NodeAffinities []SchedulingNodeAffinitiesParameters `json:"nodeAffinities,omitempty" tf:"node_affinities,omitempty"`

	// Describes maintenance behavior for the instance. One of MIGRATE or TERMINATE,
	// +kubebuilder:validation:Optional
	OnHostMaintenance *string `json:"onHostMaintenance,omitempty" tf:"on_host_maintenance,omitempty"`

	// Whether the instance is preemptible.
	// +kubebuilder:validation:Optional
	Preemptible *bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`

	// Whether the instance is spot. If this is set as SPOT.
	// +kubebuilder:validation:Optional
	ProvisioningModel *string `json:"provisioningModel,omitempty" tf:"provisioning_model,omitempty"`
}

type InstanceFromTemplateScratchDiskObservation struct {
}

type InstanceFromTemplateScratchDiskParameters struct {

	// +kubebuilder:validation:Optional
	Interface *string `json:"interface,omitempty" tf:"interface"`
}

type InstanceFromTemplateServiceAccountObservation struct {
}

type InstanceFromTemplateServiceAccountParameters struct {

	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email"`

	// +kubebuilder:validation:Optional
	Scopes []*string `json:"scopes,omitempty" tf:"scopes"`
}

type InstanceFromTemplateShieldedInstanceConfigObservation struct {
}

type InstanceFromTemplateShieldedInstanceConfigParameters struct {

	// Whether integrity monitoring is enabled for the instance.
	// +kubebuilder:validation:Optional
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty" tf:"enable_integrity_monitoring,omitempty"`

	// Whether secure boot is enabled for the instance.
	// +kubebuilder:validation:Optional
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty" tf:"enable_secure_boot,omitempty"`

	// Whether the instance uses vTPM.
	// +kubebuilder:validation:Optional
	EnableVtpm *bool `json:"enableVtpm,omitempty" tf:"enable_vtpm,omitempty"`
}

type NetworkInterfaceAccessConfigObservation struct {
}

type NetworkInterfaceAccessConfigParameters struct {

	// +kubebuilder:validation:Optional
	NATIP *string `json:"natIp,omitempty" tf:"nat_ip"`

	// +kubebuilder:validation:Optional
	NetworkTier *string `json:"networkTier,omitempty" tf:"network_tier"`

	// +kubebuilder:validation:Optional
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty" tf:"public_ptr_domain_name"`
}

type NetworkInterfaceAliasIPRangeObservation struct {
}

type NetworkInterfaceAliasIPRangeParameters struct {

	// +kubebuilder:validation:Optional
	IPCidrRange *string `json:"ipCidrRange,omitempty" tf:"ip_cidr_range"`

	// +kubebuilder:validation:Optional
	SubnetworkRangeName *string `json:"subnetworkRangeName,omitempty" tf:"subnetwork_range_name"`
}

type NetworkInterfaceIPv6AccessConfigObservation struct {

	// The first IPv6 address of the external IPv6 range associated with this instance, prefix length is stored in externalIpv6PrefixLength in ipv6AccessConfig. The field is output only, an IPv6 address from a subnetwork associated with the instance will be allocated dynamically.
	ExternalIPv6 *string `json:"externalIpv6,omitempty" tf:"external_ipv6,omitempty"`

	// The prefix length of the external IPv6 range.
	ExternalIPv6PrefixLength *string `json:"externalIpv6PrefixLength,omitempty" tf:"external_ipv6_prefix_length,omitempty"`
}

type NetworkInterfaceIPv6AccessConfigParameters struct {

	// The service-level to be provided for IPv6 traffic when the subnet has an external subnet. Only PREMIUM tier is valid for IPv6
	// +kubebuilder:validation:Required
	NetworkTier *string `json:"networkTier" tf:"network_tier,omitempty"`

	// The domain name to be used when creating DNSv6 records for the external IPv6 ranges.
	// +kubebuilder:validation:Optional
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty" tf:"public_ptr_domain_name,omitempty"`
}

type ReservationAffinitySpecificReservationObservation struct {
}

type ReservationAffinitySpecificReservationParameters struct {

	// Corresponds to the label key of a reservation resource. To target a SPECIFIC_RESERVATION by name, specify compute.googleapis.com/reservation-name as the key and specify the name of your reservation as the only value.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Corresponds to the label values of a reservation resource.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type SchedulingNodeAffinitiesObservation struct {
}

type SchedulingNodeAffinitiesParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

// InstanceFromTemplateSpec defines the desired state of InstanceFromTemplate
type InstanceFromTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceFromTemplateParameters `json:"forProvider"`
}

// InstanceFromTemplateStatus defines the observed state of InstanceFromTemplate.
type InstanceFromTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceFromTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceFromTemplate is the Schema for the InstanceFromTemplates API. Manages a VM instance resource within GCE.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type InstanceFromTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceFromTemplateSpec   `json:"spec"`
	Status            InstanceFromTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceFromTemplateList contains a list of InstanceFromTemplates
type InstanceFromTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceFromTemplate `json:"items"`
}

// Repository type metadata.
var (
	InstanceFromTemplate_Kind             = "InstanceFromTemplate"
	InstanceFromTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceFromTemplate_Kind}.String()
	InstanceFromTemplate_KindAPIVersion   = InstanceFromTemplate_Kind + "." + CRDGroupVersion.String()
	InstanceFromTemplate_GroupVersionKind = CRDGroupVersion.WithKind(InstanceFromTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceFromTemplate{}, &InstanceFromTemplateList{})
}
