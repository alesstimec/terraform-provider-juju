// Copyright 2025 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/juju/juju/core/model"
	"github.com/juju/terraform-provider-juju/internal/juju"
)

type caasModelValidator struct {
	client *juju.Client
}

// Description returns a plain text description of the validator's behavior, suitable for a practitioner to understand its impact.
func (v caasModelValidator) Description(context.Context) string {
	return "string must conform to name@channel, e.g. ubuntu@22.04"
}

// MarkdownDescription returns a markdown formatted description of the validator's behavior, suitable for a practitioner to understand its impact.
func (v caasModelValidator) MarkdownDescription(context.Context) string {
	return "string must conform to name@channel, e.g. ubuntu@22.04"
}

// Validate runs the main validation logic of the validator, reading configuration data out of `req` and updating `resp` with diagnostics.
func (v caasModelValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	// If the value is unknown or null, there is nothing to validate.
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	var modelNameField basetypes.StringValuable
	diags := req.Config.GetAttribute(ctx, path.Root("model"), &modelNameField)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	modelName := modelNameField.String()

	if v.client == nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Cannot determine model type",
			"Juju client not specified.",
		)
		return
	}
	modelType, err := v.client.Applications.SharedClient.ModelType(modelName)
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Cannot determine model type",
			fmt.Sprintf("Cannot determine type of model %q", modelName),
		)
		return
	}

	if modelType != model.IAAS {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid model type",
			"Machines may only be specified for IAAS models.",
		)
		return
	}
}
