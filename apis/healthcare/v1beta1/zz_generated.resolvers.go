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

func (mg *ConsentStore) ResolveReferences( // ResolveReferences of this ConsentStore.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("healthcare.gcp.upbound.io", "v1beta1", "Dataset", "DatasetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Dataset),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DatasetRef,
			Selector:     mg.Spec.ForProvider.DatasetSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Dataset")
	}
	mg.Spec.ForProvider.Dataset = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatasetRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("healthcare.gcp.upbound.io", "v1beta1", "Dataset", "DatasetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Dataset),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.DatasetRef,
			Selector:     mg.Spec.InitProvider.DatasetSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Dataset")
	}
	mg.Spec.InitProvider.Dataset = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DatasetRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Dataset.
func (mg *Dataset) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.EncryptionSpec != nil {
		{
			m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta2", "CryptoKey", "CryptoKeyList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EncryptionSpec.KMSKeyName),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.EncryptionSpec.KMSKeyNameRef,
				Selector:     mg.Spec.ForProvider.EncryptionSpec.KMSKeyNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EncryptionSpec.KMSKeyName")
		}
		mg.Spec.ForProvider.EncryptionSpec.KMSKeyName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EncryptionSpec.KMSKeyNameRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.EncryptionSpec != nil {
		{
			m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta2", "CryptoKey", "CryptoKeyList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EncryptionSpec.KMSKeyName),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.EncryptionSpec.KMSKeyNameRef,
				Selector:     mg.Spec.InitProvider.EncryptionSpec.KMSKeyNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EncryptionSpec.KMSKeyName")
		}
		mg.Spec.InitProvider.EncryptionSpec.KMSKeyName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.EncryptionSpec.KMSKeyNameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this DatasetIAMMember.
func (mg *DatasetIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("healthcare.gcp.upbound.io", "v1beta1", "Dataset", "DatasetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatasetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DatasetIDRef,
			Selector:     mg.Spec.ForProvider.DatasetIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatasetID")
	}
	mg.Spec.ForProvider.DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatasetIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("healthcare.gcp.upbound.io", "v1beta1", "Dataset", "DatasetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DatasetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.DatasetIDRef,
			Selector:     mg.Spec.InitProvider.DatasetIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DatasetID")
	}
	mg.Spec.InitProvider.DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DatasetIDRef = rsp.ResolvedReference

	return nil
}
