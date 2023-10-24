package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	crypto_rand "crypto/rand"
	b64 "encoding/base64"
	"golang.org/x/crypto/nacl/box"
)

type keyPairResource struct{}

var (
	_ resource.Resource                 = (*keyPairResource)(nil)
	_ resource.ResourceWithUpgradeState = (*keyPairResource)(nil)
)

func NewKeyPairResource() resource.Resource {
	return &keyPairResource{}
}

func (r *keyPairResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_key_pair"
}

func (r *keyPairResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 0,
		Attributes: map[string]schema.Attribute{
			// Computed attributes
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "ID of the resource (public key).",
			},

			"secret_key": schema.StringAttribute{
				Computed:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Secret key data in base64 format.",
			},

			"public_key": schema.StringAttribute{
				Computed:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Public key data in base64 format.",
			},
		},
		MarkdownDescription: "Creates a Sodium formatted key pair.\n\n" +
			"Generates a secure key pair and encodes it in " +
			"base64",
	}
}

func (r *keyPairResource) Create(ctx context.Context, req resource.CreateRequest, res *resource.CreateResponse) {
	tflog.Debug(ctx, "Creating key pair resource")

	// Load entire configuration into the model
	var newState keyPairResourceModel
	res.Diagnostics.Append(req.Plan.Get(ctx, &newState)...)
	if res.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Loaded key pair configuration", map[string]interface{}{
		"KeyConfig": fmt.Sprintf("%+v", newState),
	})

	pk, sk, err := box.GenerateKey(crypto_rand.Reader)
	if err != nil {
		panic(err)
	}

	newState.Id = types.StringValue(b64.StdEncoding.EncodeToString([]byte(string(pk[:]))))
	newState.PublicKey = types.StringValue(b64.StdEncoding.EncodeToString([]byte(string(pk[:]))))
	newState.SecretKey = types.StringValue(b64.StdEncoding.EncodeToString([]byte(string(sk[:]))))

	// Store the model populated so far, onto the State
	tflog.Debug(ctx, "Storing into the state")
	res.Diagnostics.Append(res.State.Set(ctx, newState)...)
	if res.Diagnostics.HasError() {
		return
	}
}

func (r *keyPairResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}

func (r *keyPairResource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	// NO-OP: all there is to read is in the State, and response is already populated with that.
	tflog.Debug(ctx, "Reading key pair from state")
}

func (r *keyPairResource) Update(_ context.Context, _ resource.UpdateRequest, _ *resource.UpdateResponse) {
	// NO-OP: changes to this resource will force a "re-create".
}

func (r *keyPairResource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// NO-OP: Returning no error is enough for the framework to remove the resource from state.
	tflog.Debug(ctx, "Removing key pair from state")
}

func (r *keyPairResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

type keyPairResourceModel struct {
	SecretKey types.String `tfsdk:"secret_key"`
	PublicKey types.String `tfsdk:"public_key"`
	Id        types.String `tfsdk:"id"`
}
