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

type PasswordPolicyInitParameters struct {

	// Number of failed attempts allowed before the user get locked.
	AllowedFailedAttempts *float64 `json:"allowedFailedAttempts,omitempty" tf:"allowed_failed_attempts,omitempty"`

	// If true, the check that will lock user after too many failed login attempts will be enabled.
	EnableFailedAttemptsCheck *bool `json:"enableFailedAttemptsCheck,omitempty" tf:"enable_failed_attempts_check,omitempty"`

	// If true, the user must specify the current password before changing the password. This flag is supported only for MySQL.
	EnablePasswordVerification *bool `json:"enablePasswordVerification,omitempty" tf:"enable_password_verification,omitempty"`

	// Password expiration duration with one week grace period.
	PasswordExpirationDuration *string `json:"passwordExpirationDuration,omitempty" tf:"password_expiration_duration,omitempty"`
}

type PasswordPolicyObservation struct {

	// Number of failed attempts allowed before the user get locked.
	AllowedFailedAttempts *float64 `json:"allowedFailedAttempts,omitempty" tf:"allowed_failed_attempts,omitempty"`

	// If true, the check that will lock user after too many failed login attempts will be enabled.
	EnableFailedAttemptsCheck *bool `json:"enableFailedAttemptsCheck,omitempty" tf:"enable_failed_attempts_check,omitempty"`

	// If true, the user must specify the current password before changing the password. This flag is supported only for MySQL.
	EnablePasswordVerification *bool `json:"enablePasswordVerification,omitempty" tf:"enable_password_verification,omitempty"`

	// Password expiration duration with one week grace period.
	PasswordExpirationDuration *string `json:"passwordExpirationDuration,omitempty" tf:"password_expiration_duration,omitempty"`

	Status []StatusObservation `json:"status,omitempty" tf:"status,omitempty"`
}

type PasswordPolicyParameters struct {

	// Number of failed attempts allowed before the user get locked.
	// +kubebuilder:validation:Optional
	AllowedFailedAttempts *float64 `json:"allowedFailedAttempts,omitempty" tf:"allowed_failed_attempts,omitempty"`

	// If true, the check that will lock user after too many failed login attempts will be enabled.
	// +kubebuilder:validation:Optional
	EnableFailedAttemptsCheck *bool `json:"enableFailedAttemptsCheck,omitempty" tf:"enable_failed_attempts_check,omitempty"`

	// If true, the user must specify the current password before changing the password. This flag is supported only for MySQL.
	// +kubebuilder:validation:Optional
	EnablePasswordVerification *bool `json:"enablePasswordVerification,omitempty" tf:"enable_password_verification,omitempty"`

	// Password expiration duration with one week grace period.
	// +kubebuilder:validation:Optional
	PasswordExpirationDuration *string `json:"passwordExpirationDuration,omitempty" tf:"password_expiration_duration,omitempty"`
}

type SQLServerUserDetailsInitParameters struct {
}

type SQLServerUserDetailsObservation struct {
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	ServerRoles []*string `json:"serverRoles,omitempty" tf:"server_roles,omitempty"`
}

type SQLServerUserDetailsParameters struct {
}

type StatusInitParameters struct {
}

type StatusObservation struct {

	// (read only) If true, user does not have login privileges.
	Locked *bool `json:"locked,omitempty" tf:"locked,omitempty"`

	// (read only) Password expiration duration with one week grace period.
	PasswordExpirationTime *string `json:"passwordExpirationTime,omitempty" tf:"password_expiration_time,omitempty"`
}

type StatusParameters struct {
}

type UserInitParameters struct {

	// The deletion policy for the user.
	// Setting ABANDON allows the resource to be abandoned rather than deleted. This is useful
	// for Postgres, where users cannot be deleted from the API if they have been granted SQL roles.
	DeletionPolicy *string `json:"deletionPolicy,omitempty" tf:"deletion_policy,omitempty"`

	// The host the user can connect from. This is only supported
	// for BUILT_IN users in MySQL instances. Don't set this field for PostgreSQL and SQL Server instances.
	// Can be an IP address. Changing this forces a new resource to be created.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The name of the Cloud SQL instance. Changing this
	// forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/sql/v1beta2.DatabaseInstance
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// Reference to a DatabaseInstance in sql to populate instance.
	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// Selector for a DatabaseInstance in sql to populate instance.
	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	PasswordPolicy *PasswordPolicyInitParameters `json:"passwordPolicy,omitempty" tf:"password_policy,omitempty"`

	// The password for the user. Can be updated. For Postgres
	// instances this is a Required field, unless type is set to either CLOUD_IAM_USER
	// or CLOUD_IAM_SERVICE_ACCOUNT. Don't set this field for CLOUD_IAM_USER
	// and CLOUD_IAM_SERVICE_ACCOUNT user types for any Cloud SQL instance.
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The user type. It determines the method to authenticate the
	// user during login. The default is the database's built-in user type. Flags
	// include "BUILT_IN", "CLOUD_IAM_USER", and "CLOUD_IAM_SERVICE_ACCOUNT" for both
	// Postgres and MySQL.
	// MySQL also includes "CLOUD_IAM_GROUP", "CLOUD_IAM_GROUP_USER" and "CLOUD_IAM_GROUP_SERVICE_ACCOUNT".
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type UserObservation struct {

	// The deletion policy for the user.
	// Setting ABANDON allows the resource to be abandoned rather than deleted. This is useful
	// for Postgres, where users cannot be deleted from the API if they have been granted SQL roles.
	DeletionPolicy *string `json:"deletionPolicy,omitempty" tf:"deletion_policy,omitempty"`

	// The host the user can connect from. This is only supported
	// for BUILT_IN users in MySQL instances. Don't set this field for PostgreSQL and SQL Server instances.
	// Can be an IP address. Changing this forces a new resource to be created.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Cloud SQL instance. Changing this
	// forces a new resource to be created.
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	PasswordPolicy *PasswordPolicyObservation `json:"passwordPolicy,omitempty" tf:"password_policy,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	SQLServerUserDetails []SQLServerUserDetailsObservation `json:"sqlServerUserDetails,omitempty" tf:"sql_server_user_details,omitempty"`

	// The user type. It determines the method to authenticate the
	// user during login. The default is the database's built-in user type. Flags
	// include "BUILT_IN", "CLOUD_IAM_USER", and "CLOUD_IAM_SERVICE_ACCOUNT" for both
	// Postgres and MySQL.
	// MySQL also includes "CLOUD_IAM_GROUP", "CLOUD_IAM_GROUP_USER" and "CLOUD_IAM_GROUP_SERVICE_ACCOUNT".
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type UserParameters struct {

	// The deletion policy for the user.
	// Setting ABANDON allows the resource to be abandoned rather than deleted. This is useful
	// for Postgres, where users cannot be deleted from the API if they have been granted SQL roles.
	// +kubebuilder:validation:Optional
	DeletionPolicy *string `json:"deletionPolicy,omitempty" tf:"deletion_policy,omitempty"`

	// The host the user can connect from. This is only supported
	// for BUILT_IN users in MySQL instances. Don't set this field for PostgreSQL and SQL Server instances.
	// Can be an IP address. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The name of the Cloud SQL instance. Changing this
	// forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/sql/v1beta2.DatabaseInstance
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// Reference to a DatabaseInstance in sql to populate instance.
	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// Selector for a DatabaseInstance in sql to populate instance.
	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PasswordPolicy *PasswordPolicyParameters `json:"passwordPolicy,omitempty" tf:"password_policy,omitempty"`

	// The password for the user. Can be updated. For Postgres
	// instances this is a Required field, unless type is set to either CLOUD_IAM_USER
	// or CLOUD_IAM_SERVICE_ACCOUNT. Don't set this field for CLOUD_IAM_USER
	// and CLOUD_IAM_SERVICE_ACCOUNT user types for any Cloud SQL instance.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The user type. It determines the method to authenticate the
	// user during login. The default is the database's built-in user type. Flags
	// include "BUILT_IN", "CLOUD_IAM_USER", and "CLOUD_IAM_SERVICE_ACCOUNT" for both
	// Postgres and MySQL.
	// MySQL also includes "CLOUD_IAM_GROUP", "CLOUD_IAM_GROUP_USER" and "CLOUD_IAM_GROUP_SERVICE_ACCOUNT".
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
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
	InitProvider UserInitParameters `json:"initProvider,omitempty"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// User is the Schema for the Users API. Creates a new SQL user in Google Cloud SQL.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserSpec   `json:"spec"`
	Status            UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
