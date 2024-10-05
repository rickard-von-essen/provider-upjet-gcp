// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *CertificateMapEntry) ResolveReferences( // ResolveReferences of this CertificateMapEntry.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("certificatemanager.gcp.upbound.io", "v1beta2", "Certificate", "CertificateList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Certificates),
			Extract:       resource.ExtractResourceID(),
			References:    mg.Spec.ForProvider.CertificatesRefs,
			Selector:      mg.Spec.ForProvider.CertificatesSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Certificates")
	}
	mg.Spec.ForProvider.Certificates = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.CertificatesRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("certificatemanager.gcp.upbound.io", "v1beta1", "CertificateMap", "CertificateMapList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Map),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.MapRef,
			Selector:     mg.Spec.ForProvider.MapSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Map")
	}
	mg.Spec.ForProvider.Map = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MapRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("certificatemanager.gcp.upbound.io", "v1beta2", "Certificate", "CertificateList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Certificates),
			Extract:       resource.ExtractResourceID(),
			References:    mg.Spec.InitProvider.CertificatesRefs,
			Selector:      mg.Spec.InitProvider.CertificatesSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Certificates")
	}
	mg.Spec.InitProvider.Certificates = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.CertificatesRefs = mrsp.ResolvedReferences

	return nil
}
