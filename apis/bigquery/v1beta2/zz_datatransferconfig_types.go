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

type DataTransferConfigInitParameters struct {

	// The number of days to look back to automatically refresh the data.
	// For example, if dataRefreshWindowDays = 10, then every day BigQuery
	// reingests data for [today-10, today-1], rather than ingesting data for
	// just [today-1]. Only valid if the data source supports the feature.
	// Set the value to 0 to use the default value.
	DataRefreshWindowDays *float64 `json:"dataRefreshWindowDays,omitempty" tf:"data_refresh_window_days,omitempty"`

	// The data source id. Cannot be changed once the transfer config is created.
	DataSourceID *string `json:"dataSourceId,omitempty" tf:"data_source_id,omitempty"`

	// The BigQuery target dataset id.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta2.Dataset
	DestinationDatasetID *string `json:"destinationDatasetId,omitempty" tf:"destination_dataset_id,omitempty"`

	// Reference to a Dataset in bigquery to populate destinationDatasetId.
	// +kubebuilder:validation:Optional
	DestinationDatasetIDRef *v1.Reference `json:"destinationDatasetIdRef,omitempty" tf:"-"`

	// Selector for a Dataset in bigquery to populate destinationDatasetId.
	// +kubebuilder:validation:Optional
	DestinationDatasetIDSelector *v1.Selector `json:"destinationDatasetIdSelector,omitempty" tf:"-"`

	// When set to true, no runs are scheduled for a given transfer.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The user specified display name for the transfer config.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Email notifications will be sent according to these preferences to the
	// email address of the user who owns this transfer config.
	// Structure is documented below.
	EmailPreferences *EmailPreferencesInitParameters `json:"emailPreferences,omitempty" tf:"email_preferences,omitempty"`

	// The geographic location where the transfer config should reside.
	// Examples: US, EU, asia-northeast1. The default value is US.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Pub/Sub topic where notifications will be sent after transfer runs
	// associated with this transfer config finish.
	NotificationPubsubTopic *string `json:"notificationPubsubTopic,omitempty" tf:"notification_pubsub_topic,omitempty"`

	// Parameters specific to each data source. For more information see the bq tab in the 'Setting up a data transfer'
	// section for each data source. For example the parameters for Cloud Storage transfers are listed here:
	// https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
	// NOTE : If you are attempting to update a parameter that cannot be updated (due to api limitations) please force recreation of the resource.
	// +mapType=granular
	Params map[string]*string `json:"params,omitempty" tf:"params,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Data transfer schedule. If the data source does not support a custom
	// schedule, this should be empty. If it is empty, the default value for
	// the data source will be used. The specified times are in UTC. Examples
	// of valid format: 1st,3rd monday of month 15:30, every wed,fri of jan,
	// jun 13:15, and first sunday of quarter 00:00. See more explanation
	// about the format here:
	// https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
	// NOTE: The minimum interval time between recurring transfers depends
	// on the data source; refer to the documentation for your data source.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// Options customizing the data transfer schedule.
	// Structure is documented below.
	ScheduleOptions *ScheduleOptionsInitParameters `json:"scheduleOptions,omitempty" tf:"schedule_options,omitempty"`

	// Different parameters are configured primarily using the the params field on this
	// resource. This block contains the parameters which contain secrets or passwords so that they can be marked
	// sensitive and hidden from plan output. The name of the field, eg: secret_access_key, will be the key
	// in the params map in the api request.
	// Credentials may not be specified in both locations and will cause an error. Changing from one location
	// to a different credential configuration in the config will require an apply to update state.
	// Structure is documented below.
	SensitiveParams *SensitiveParamsInitParameters `json:"sensitiveParams,omitempty" tf:"sensitive_params,omitempty"`

	// Service account email. If this field is set, transfer config will
	// be created with this service account credentials. It requires that
	// requesting user calling this API has permissions to act as this service account.
	ServiceAccountName *string `json:"serviceAccountName,omitempty" tf:"service_account_name,omitempty"`
}

type DataTransferConfigObservation struct {

	// The number of days to look back to automatically refresh the data.
	// For example, if dataRefreshWindowDays = 10, then every day BigQuery
	// reingests data for [today-10, today-1], rather than ingesting data for
	// just [today-1]. Only valid if the data source supports the feature.
	// Set the value to 0 to use the default value.
	DataRefreshWindowDays *float64 `json:"dataRefreshWindowDays,omitempty" tf:"data_refresh_window_days,omitempty"`

	// The data source id. Cannot be changed once the transfer config is created.
	DataSourceID *string `json:"dataSourceId,omitempty" tf:"data_source_id,omitempty"`

	// The BigQuery target dataset id.
	DestinationDatasetID *string `json:"destinationDatasetId,omitempty" tf:"destination_dataset_id,omitempty"`

	// When set to true, no runs are scheduled for a given transfer.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The user specified display name for the transfer config.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Email notifications will be sent according to these preferences to the
	// email address of the user who owns this transfer config.
	// Structure is documented below.
	EmailPreferences *EmailPreferencesObservation `json:"emailPreferences,omitempty" tf:"email_preferences,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The geographic location where the transfer config should reside.
	// Examples: US, EU, asia-northeast1. The default value is US.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The resource name of the transfer config. Transfer config names have the
	// form projects/{projectId}/locations/{location}/transferConfigs/{configId}
	// or projects/{projectId}/transferConfigs/{configId},
	// where configId is usually a uuid, but this is not required.
	// The name is ignored when creating a transfer config.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Pub/Sub topic where notifications will be sent after transfer runs
	// associated with this transfer config finish.
	NotificationPubsubTopic *string `json:"notificationPubsubTopic,omitempty" tf:"notification_pubsub_topic,omitempty"`

	// Parameters specific to each data source. For more information see the bq tab in the 'Setting up a data transfer'
	// section for each data source. For example the parameters for Cloud Storage transfers are listed here:
	// https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
	// NOTE : If you are attempting to update a parameter that cannot be updated (due to api limitations) please force recreation of the resource.
	// +mapType=granular
	Params map[string]*string `json:"params,omitempty" tf:"params,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Data transfer schedule. If the data source does not support a custom
	// schedule, this should be empty. If it is empty, the default value for
	// the data source will be used. The specified times are in UTC. Examples
	// of valid format: 1st,3rd monday of month 15:30, every wed,fri of jan,
	// jun 13:15, and first sunday of quarter 00:00. See more explanation
	// about the format here:
	// https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
	// NOTE: The minimum interval time between recurring transfers depends
	// on the data source; refer to the documentation for your data source.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// Options customizing the data transfer schedule.
	// Structure is documented below.
	ScheduleOptions *ScheduleOptionsObservation `json:"scheduleOptions,omitempty" tf:"schedule_options,omitempty"`

	// Different parameters are configured primarily using the the params field on this
	// resource. This block contains the parameters which contain secrets or passwords so that they can be marked
	// sensitive and hidden from plan output. The name of the field, eg: secret_access_key, will be the key
	// in the params map in the api request.
	// Credentials may not be specified in both locations and will cause an error. Changing from one location
	// to a different credential configuration in the config will require an apply to update state.
	// Structure is documented below.
	SensitiveParams *SensitiveParamsParameters `json:"sensitiveParams,omitempty" tf:"sensitive_params,omitempty"`

	// Service account email. If this field is set, transfer config will
	// be created with this service account credentials. It requires that
	// requesting user calling this API has permissions to act as this service account.
	ServiceAccountName *string `json:"serviceAccountName,omitempty" tf:"service_account_name,omitempty"`
}

type DataTransferConfigParameters struct {

	// The number of days to look back to automatically refresh the data.
	// For example, if dataRefreshWindowDays = 10, then every day BigQuery
	// reingests data for [today-10, today-1], rather than ingesting data for
	// just [today-1]. Only valid if the data source supports the feature.
	// Set the value to 0 to use the default value.
	// +kubebuilder:validation:Optional
	DataRefreshWindowDays *float64 `json:"dataRefreshWindowDays,omitempty" tf:"data_refresh_window_days,omitempty"`

	// The data source id. Cannot be changed once the transfer config is created.
	// +kubebuilder:validation:Optional
	DataSourceID *string `json:"dataSourceId,omitempty" tf:"data_source_id,omitempty"`

	// The BigQuery target dataset id.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta2.Dataset
	// +kubebuilder:validation:Optional
	DestinationDatasetID *string `json:"destinationDatasetId,omitempty" tf:"destination_dataset_id,omitempty"`

	// Reference to a Dataset in bigquery to populate destinationDatasetId.
	// +kubebuilder:validation:Optional
	DestinationDatasetIDRef *v1.Reference `json:"destinationDatasetIdRef,omitempty" tf:"-"`

	// Selector for a Dataset in bigquery to populate destinationDatasetId.
	// +kubebuilder:validation:Optional
	DestinationDatasetIDSelector *v1.Selector `json:"destinationDatasetIdSelector,omitempty" tf:"-"`

	// When set to true, no runs are scheduled for a given transfer.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The user specified display name for the transfer config.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Email notifications will be sent according to these preferences to the
	// email address of the user who owns this transfer config.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EmailPreferences *EmailPreferencesParameters `json:"emailPreferences,omitempty" tf:"email_preferences,omitempty"`

	// The geographic location where the transfer config should reside.
	// Examples: US, EU, asia-northeast1. The default value is US.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Pub/Sub topic where notifications will be sent after transfer runs
	// associated with this transfer config finish.
	// +kubebuilder:validation:Optional
	NotificationPubsubTopic *string `json:"notificationPubsubTopic,omitempty" tf:"notification_pubsub_topic,omitempty"`

	// Parameters specific to each data source. For more information see the bq tab in the 'Setting up a data transfer'
	// section for each data source. For example the parameters for Cloud Storage transfers are listed here:
	// https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
	// NOTE : If you are attempting to update a parameter that cannot be updated (due to api limitations) please force recreation of the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Params map[string]*string `json:"params,omitempty" tf:"params,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Data transfer schedule. If the data source does not support a custom
	// schedule, this should be empty. If it is empty, the default value for
	// the data source will be used. The specified times are in UTC. Examples
	// of valid format: 1st,3rd monday of month 15:30, every wed,fri of jan,
	// jun 13:15, and first sunday of quarter 00:00. See more explanation
	// about the format here:
	// https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
	// NOTE: The minimum interval time between recurring transfers depends
	// on the data source; refer to the documentation for your data source.
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// Options customizing the data transfer schedule.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ScheduleOptions *ScheduleOptionsParameters `json:"scheduleOptions,omitempty" tf:"schedule_options,omitempty"`

	// Different parameters are configured primarily using the the params field on this
	// resource. This block contains the parameters which contain secrets or passwords so that they can be marked
	// sensitive and hidden from plan output. The name of the field, eg: secret_access_key, will be the key
	// in the params map in the api request.
	// Credentials may not be specified in both locations and will cause an error. Changing from one location
	// to a different credential configuration in the config will require an apply to update state.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SensitiveParams *SensitiveParamsParameters `json:"sensitiveParams,omitempty" tf:"sensitive_params,omitempty"`

	// Service account email. If this field is set, transfer config will
	// be created with this service account credentials. It requires that
	// requesting user calling this API has permissions to act as this service account.
	// +kubebuilder:validation:Optional
	ServiceAccountName *string `json:"serviceAccountName,omitempty" tf:"service_account_name,omitempty"`
}

type EmailPreferencesInitParameters struct {

	// If true, email notifications will be sent on transfer run failures.
	EnableFailureEmail *bool `json:"enableFailureEmail,omitempty" tf:"enable_failure_email,omitempty"`
}

type EmailPreferencesObservation struct {

	// If true, email notifications will be sent on transfer run failures.
	EnableFailureEmail *bool `json:"enableFailureEmail,omitempty" tf:"enable_failure_email,omitempty"`
}

type EmailPreferencesParameters struct {

	// If true, email notifications will be sent on transfer run failures.
	// +kubebuilder:validation:Optional
	EnableFailureEmail *bool `json:"enableFailureEmail" tf:"enable_failure_email,omitempty"`
}

type ScheduleOptionsInitParameters struct {

	// If true, automatic scheduling of data transfer runs for this
	// configuration will be disabled. The runs can be started on ad-hoc
	// basis using transferConfigs.startManualRuns API. When automatic
	// scheduling is disabled, the TransferConfig.schedule field will
	// be ignored.
	DisableAutoScheduling *bool `json:"disableAutoScheduling,omitempty" tf:"disable_auto_scheduling,omitempty"`

	// Defines time to stop scheduling transfer runs. A transfer run cannot be
	// scheduled at or after the end time. The end time can be changed at any
	// moment. The time when a data transfer can be triggered manually is not
	// limited by this option.
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies time to start scheduling transfer runs. The first run will be
	// scheduled at or after the start time according to a recurrence pattern
	// defined in the schedule string. The start time can be changed at any
	// moment. The time when a data transfer can be triggered manually is not
	// limited by this option.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type ScheduleOptionsObservation struct {

	// If true, automatic scheduling of data transfer runs for this
	// configuration will be disabled. The runs can be started on ad-hoc
	// basis using transferConfigs.startManualRuns API. When automatic
	// scheduling is disabled, the TransferConfig.schedule field will
	// be ignored.
	DisableAutoScheduling *bool `json:"disableAutoScheduling,omitempty" tf:"disable_auto_scheduling,omitempty"`

	// Defines time to stop scheduling transfer runs. A transfer run cannot be
	// scheduled at or after the end time. The end time can be changed at any
	// moment. The time when a data transfer can be triggered manually is not
	// limited by this option.
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies time to start scheduling transfer runs. The first run will be
	// scheduled at or after the start time according to a recurrence pattern
	// defined in the schedule string. The start time can be changed at any
	// moment. The time when a data transfer can be triggered manually is not
	// limited by this option.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type ScheduleOptionsParameters struct {

	// If true, automatic scheduling of data transfer runs for this
	// configuration will be disabled. The runs can be started on ad-hoc
	// basis using transferConfigs.startManualRuns API. When automatic
	// scheduling is disabled, the TransferConfig.schedule field will
	// be ignored.
	// +kubebuilder:validation:Optional
	DisableAutoScheduling *bool `json:"disableAutoScheduling,omitempty" tf:"disable_auto_scheduling,omitempty"`

	// Defines time to stop scheduling transfer runs. A transfer run cannot be
	// scheduled at or after the end time. The end time can be changed at any
	// moment. The time when a data transfer can be triggered manually is not
	// limited by this option.
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies time to start scheduling transfer runs. The first run will be
	// scheduled at or after the start time according to a recurrence pattern
	// defined in the schedule string. The start time can be changed at any
	// moment. The time when a data transfer can be triggered manually is not
	// limited by this option.
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type SensitiveParamsInitParameters struct {

	// The Secret Access Key of the AWS account transferring data from.
	// Note: This property is sensitive and will not be displayed in the plan.
	SecretAccessKeySecretRef v1.SecretKeySelector `json:"secretAccessKeySecretRef" tf:"-"`
}

type SensitiveParamsObservation struct {
}

type SensitiveParamsParameters struct {

	// The Secret Access Key of the AWS account transferring data from.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	SecretAccessKeySecretRef v1.SecretKeySelector `json:"secretAccessKeySecretRef" tf:"-"`
}

// DataTransferConfigSpec defines the desired state of DataTransferConfig
type DataTransferConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataTransferConfigParameters `json:"forProvider"`
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
	InitProvider DataTransferConfigInitParameters `json:"initProvider,omitempty"`
}

// DataTransferConfigStatus defines the observed state of DataTransferConfig.
type DataTransferConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataTransferConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DataTransferConfig is the Schema for the DataTransferConfigs API. Represents a data transfer configuration.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type DataTransferConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataSourceId) || (has(self.initProvider) && has(self.initProvider.dataSourceId))",message="spec.forProvider.dataSourceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.params) || (has(self.initProvider) && has(self.initProvider.params))",message="spec.forProvider.params is a required parameter"
	Spec   DataTransferConfigSpec   `json:"spec"`
	Status DataTransferConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataTransferConfigList contains a list of DataTransferConfigs
type DataTransferConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataTransferConfig `json:"items"`
}

// Repository type metadata.
var (
	DataTransferConfig_Kind             = "DataTransferConfig"
	DataTransferConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataTransferConfig_Kind}.String()
	DataTransferConfig_KindAPIVersion   = DataTransferConfig_Kind + "." + CRDGroupVersion.String()
	DataTransferConfig_GroupVersionKind = CRDGroupVersion.WithKind(DataTransferConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&DataTransferConfig{}, &DataTransferConfigList{})
}
