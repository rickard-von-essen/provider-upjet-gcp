// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	ujconversion "github.com/crossplane/upjet/pkg/controller/conversion"
	"github.com/crossplane/upjet/pkg/resource"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this AutoscalingPolicy to the hub type.
func (tr *AutoscalingPolicy) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the AutoscalingPolicy type.
func (tr *AutoscalingPolicy) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this Cluster to the hub type.
func (tr *Cluster) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the Cluster type.
func (tr *Cluster) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this Job to the hub type.
func (tr *Job) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the Job type.
func (tr *Job) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this MetastoreService to the hub type.
func (tr *MetastoreService) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the MetastoreService type.
func (tr *MetastoreService) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this WorkflowTemplate to the hub type.
func (tr *WorkflowTemplate) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the WorkflowTemplate type.
func (tr *WorkflowTemplate) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}