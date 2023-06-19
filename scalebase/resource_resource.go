package scalebase

import (
	"context"
	"net/http"
	"strings"
	"time"

	"golang.org/x/exp/slices"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	apiclient "github.com/kzmake/terraform-provider-scalebase/api/client"
	"github.com/kzmake/terraform-provider-scalebase/api/client/custom_field_service"
	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

var (
	_ resource.Resource                = &resourceResource{}
	_ resource.ResourceWithConfigure   = &resourceResource{}
	_ resource.ResourceWithImportState = &resourceResource{}
)

func NewResourceResource() resource.Resource {
	return &resourceResource{}
}

type resourceResource struct {
	client *apiclient.API
}

func (r *resourceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*apiclient.API)
}

func (r *resourceResource) Metadata(_ context.Context, req resource.MetadataRequest, res *resource.MetadataResponse) {
	res.TypeName = req.ProviderTypeName + "_resource"
}

func (r *resourceResource) Schema(_ context.Context, _ resource.SchemaRequest, res *resource.SchemaResponse) {
	res.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"srn": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Required: true,
					},
					"id": schema.StringAttribute{
						Required: true,
					},
				},
			},
			"custom_fields": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required: true,
						},
						"string": schema.StringAttribute{
							Optional: true,
						},
						"select": schema.StringAttribute{
							Optional: true,
						},
						"date": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func (r *resourceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, res *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("optional_id"), req, res)
}

func (r *resourceResource) Create(ctx context.Context, req resource.CreateRequest, res *resource.CreateResponse) {
	tflog.Debug(ctx, "Call api: CreateResource")

	var plan ResourceModel
	diags := req.Plan.Get(ctx, &plan)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}

	newCustomFields := []*models.Publicv1CustomField{}
	for _, c := range plan.CustomFields {
		var date *strfmt.DateTime
		if s := c.Date.ValueStringPointer(); s != nil {
			dt, _ := time.Parse(time.DateOnly, *s)
			sdt := strfmt.DateTime(dt.UTC().Add(9 * time.Hour))
			date = &sdt
		}
		x := &models.Publicv1CustomField{
			MasterID:   c.ID.ValueString(),
			String:     c.String.ValueString(),
			SelectList: c.Select.ValueString(),
			Date:       date,
		}
		newCustomFields = append(newCustomFields, x)
	}
	in := &models.V1Resource{
		Srn: &models.V1SRN{
			ResourceID:   plan.SRN.ID.ValueStringPointer(),
			ResourceType: models.V1SRNType(strings.ToUpper("TYPE_" + plan.SRN.Type.ValueString())),
		},
		CustomFields: newCustomFields,
	}

	tflog.Debug(ctx, "  request: "+spew.Sdump(in))

	out, err := r.client.CustomFieldService.CustomFieldServiceUpdateResource(custom_field_service.NewCustomFieldServiceUpdateResourceParams().WithResource(in))
	if err != nil {
		res.Diagnostics.AddError(
			"Error creating resource",
			"Could not create resource, unexpected error: "+err.Error(),
		)
		return
	}

	tflog.Info(ctx, "  resonse: "+spew.Sdump(out.Payload))

	v := out.Payload.Resource

	plan.SRN.ID = types.StringPointerValue(v.Srn.ResourceID)
	plan.SRN.Type = types.StringValue(strings.ToLower(strings.TrimPrefix(string(v.Srn.ResourceType), "TYPE_")))
	customFields := []*CustomField{}
	var masterIds []string
	for _, c := range plan.CustomFields {
		masterIds = append(masterIds, c.ID.ValueString())
	}
	for _, x := range v.CustomFields {
		if slices.Contains(masterIds, x.MasterID) {
			c := &CustomField{
				ID:     types.StringValue(x.MasterID),
				String: types.StringNull(),
				Select: types.StringNull(),
				Date:   types.StringNull(),
			}
			if x.String != "" {
				c.String = types.StringValue(x.String)
			}
			if x.SelectList != "" {
				c.Select = types.StringValue(x.SelectList)
			}
			if x.Date != nil {
				dt, _ := time.Parse(strfmt.MarshalFormat, x.Date.String())
				c.Date = types.StringValue(dt.Add(9 * time.Hour).Format(time.DateOnly))
			}

			customFields = append(customFields, c)
		}
	}
	plan.CustomFields = customFields

	diags = res.State.Set(ctx, plan)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
}

func (r *resourceResource) Read(ctx context.Context, req resource.ReadRequest, res *resource.ReadResponse) {
	tflog.Debug(ctx, "Call api: GetResource")

	var state ResourceModel
	diags := req.State.Get(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}

	in := &models.V1GetResourceRequest{Srn: &models.V1SRN{
		ResourceID:   state.SRN.ID.ValueStringPointer(),
		ResourceType: models.V1SRNType(strings.ToUpper("TYPE_" + state.SRN.Type.ValueString())),
	}}

	tflog.Debug(ctx, "  request: "+spew.Sdump(in))

	out, err := r.client.CustomFieldService.CustomFieldServiceGetResource(custom_field_service.NewCustomFieldServiceGetResourceParams().WithBody(in))
	if err != nil {
		res.Diagnostics.AddError(
			"Error reading scalebase_resource",
			"Could not read scalebase_resource srn "+state.SRN.String()+": "+err.Error(),
		)
		return
	}

	tflog.Info(ctx, "  resonse: "+spew.Sdump(out.Payload))

	v := out.Payload.Resource

	state.SRN.ID = types.StringPointerValue(v.Srn.ResourceID)
	state.SRN.Type = types.StringValue(strings.ToLower(strings.TrimPrefix(string(v.Srn.ResourceType), "TYPE_")))
	customFields := []*CustomField{}
	var masterIds []string
	for _, c := range state.CustomFields {
		masterIds = append(masterIds, c.ID.ValueString())
	}
	for _, x := range v.CustomFields {
		if slices.Contains(masterIds, x.MasterID) {
			c := &CustomField{
				ID:     types.StringValue(x.MasterID),
				String: types.StringNull(),
				Select: types.StringNull(),
				Date:   types.StringNull(),
			}
			if x.String != "" {
				c.String = types.StringValue(x.String)
			}
			if x.SelectList != "" {
				c.Select = types.StringValue(x.SelectList)
			}
			if x.Date != nil {
				dt, _ := time.Parse(strfmt.MarshalFormat, x.Date.String())
				c.Date = types.StringValue(dt.Add(9 * time.Hour).Format(time.DateOnly))
			}

			customFields = append(customFields, c)
		}
	}
	state.CustomFields = customFields
	diags = res.State.Set(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
}

func (r *resourceResource) Update(ctx context.Context, req resource.UpdateRequest, res *resource.UpdateResponse) {
	tflog.Debug(ctx, "Call api: UpdateResource")

	var plan ResourceModel
	diags := req.Plan.Get(ctx, &plan)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}

	newCustomFields := []*models.Publicv1CustomField{}
	for _, c := range plan.CustomFields {
		var date *strfmt.DateTime
		if s := c.Date.ValueStringPointer(); s != nil {
			dt, _ := time.Parse(time.DateOnly, *s)
			sdt := strfmt.DateTime(dt.UTC().Add(9 * time.Hour))
			date = &sdt
		}
		x := &models.Publicv1CustomField{
			MasterID:   c.ID.ValueString(),
			String:     c.String.ValueString(),
			SelectList: c.Select.ValueString(),
			Date:       date,
		}
		newCustomFields = append(newCustomFields, x)
	}
	in := &models.V1Resource{
		Srn: &models.V1SRN{
			ResourceID:   plan.SRN.ID.ValueStringPointer(),
			ResourceType: models.V1SRNType(strings.ToUpper("TYPE_" + plan.SRN.Type.ValueString())),
		},
		CustomFields: newCustomFields,
	}

	tflog.Debug(ctx, "  request: "+spew.Sdump(in))

	out, err := r.client.CustomFieldService.CustomFieldServiceUpdateResource(custom_field_service.NewCustomFieldServiceUpdateResourceParams().WithResource(in))
	if err != nil {
		res.Diagnostics.AddError(
			"Error updating resource",
			"Could not update resource, unexpected error: "+err.Error(),
		)
		return
	}

	tflog.Info(ctx, "  resonse: "+spew.Sdump(out.Payload))

	v := out.Payload.Resource

	plan.SRN.ID = types.StringPointerValue(v.Srn.ResourceID)
	plan.SRN.Type = types.StringValue(strings.ToLower(strings.TrimPrefix(string(v.Srn.ResourceType), "TYPE_")))
	customFields := []*CustomField{}
	var masterIds []string
	for _, c := range plan.CustomFields {
		masterIds = append(masterIds, c.ID.ValueString())
	}
	for _, x := range v.CustomFields {
		if slices.Contains(masterIds, x.MasterID) {
			c := &CustomField{
				ID:     types.StringValue(x.MasterID),
				String: types.StringNull(),
				Select: types.StringNull(),
				Date:   types.StringNull(),
			}
			if x.String != "" {
				c.String = types.StringValue(x.String)
			}
			if x.SelectList != "" {
				c.Select = types.StringValue(x.SelectList)
			}
			if x.Date != nil {
				dt, _ := time.Parse(strfmt.MarshalFormat, x.Date.String())
				c.Date = types.StringValue(dt.Add(9 * time.Hour).Format(time.DateOnly))
			}

			customFields = append(customFields, c)
		}
	}
	plan.CustomFields = customFields

	diags = res.State.Set(ctx, plan)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
}

func (r *resourceResource) Delete(ctx context.Context, req resource.DeleteRequest, res *resource.DeleteResponse) {
	tflog.Debug(ctx, "Call api: DeleteResource")

	var plan ResourceModel

	// Read Terraform prior state data into the model
	res.Diagnostics.Append(req.State.Get(ctx, &plan)...)

	in := &models.V1Resource{
		Srn: &models.V1SRN{
			ResourceID:   plan.SRN.ID.ValueStringPointer(),
			ResourceType: models.V1SRNType(strings.ToUpper("TYPE_" + plan.SRN.Type.ValueString())),
		},
		CustomFields: []*models.Publicv1CustomField{}, // NOTE: update empty
	}

	tflog.Debug(ctx, "  request: "+spew.Sdump(in))

	out, err := r.client.CustomFieldService.CustomFieldServiceUpdateResource(custom_field_service.NewCustomFieldServiceUpdateResourceParams().WithResource(in))
	if err != nil {
		res.Diagnostics.AddError(
			"Error creating resource",
			"Could not create resource, unexpected error: "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "  resonse: "+spew.Sdump(out.Payload))

	// Return error if the HTTP status code is not 200 OK or 404 Not Found
	if !out.IsCode(http.StatusNotFound) && !out.IsSuccess() {
		res.Diagnostics.AddError(
			"Unable to delete resource",
			"An unexpected error occurred while attempting to delete the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Status: "+string(rune(out.Code())),
		)

		return
	}
}
