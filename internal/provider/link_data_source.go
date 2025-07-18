// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/ryan-blunden/terraform-provider-dub/internal/provider/types"
	"github.com/ryan-blunden/terraform-provider-dub/internal/sdk"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &LinkDataSource{}
var _ datasource.DataSourceWithConfigure = &LinkDataSource{}

func NewLinkDataSource() datasource.DataSource {
	return &LinkDataSource{}
}

// LinkDataSource is the data source implementation.
type LinkDataSource struct {
	// Provider configured SDK client.
	client *sdk.Dub
}

// LinkDataSourceModel describes the data model.
type LinkDataSourceModel struct {
	Android         types.String          `tfsdk:"android"`
	Archived        types.Bool            `tfsdk:"archived"`
	Clicks          types.Float64         `tfsdk:"clicks"`
	Comments        types.String          `tfsdk:"comments"`
	CreatedAt       types.String          `tfsdk:"created_at"`
	Description     types.String          `tfsdk:"description"`
	DoIndex         types.Bool            `tfsdk:"do_index"`
	Domain          types.String          `queryParam:"style=form,explode=true,name=domain" tfsdk:"domain"`
	ExpiredURL      types.String          `tfsdk:"expired_url"`
	ExpiresAt       types.String          `tfsdk:"expires_at"`
	ExternalID      types.String          `queryParam:"style=form,explode=true,name=externalId" tfsdk:"external_id"`
	FolderID        types.String          `tfsdk:"folder_id"`
	Geo             *tfTypes.Geo          `tfsdk:"geo"`
	ID              types.String          `queryParam:"style=form,explode=true,name=linkId" tfsdk:"id"`
	Image           types.String          `tfsdk:"image"`
	Ios             types.String          `tfsdk:"ios"`
	Key             types.String          `queryParam:"style=form,explode=true,name=key" tfsdk:"key"`
	LastClicked     types.String          `tfsdk:"last_clicked"`
	Leads           types.Float64         `tfsdk:"leads"`
	PartnerID       types.String          `tfsdk:"partner_id"`
	Password        types.String          `tfsdk:"password"`
	ProgramID       types.String          `tfsdk:"program_id"`
	ProjectID       types.String          `tfsdk:"project_id"`
	Proxy           types.Bool            `tfsdk:"proxy"`
	PublicStats     types.Bool            `tfsdk:"public_stats"`
	QrCode          types.String          `tfsdk:"qr_code"`
	Rewrite         types.Bool            `tfsdk:"rewrite"`
	SaleAmount      types.Float64         `tfsdk:"sale_amount"`
	Sales           types.Float64         `tfsdk:"sales"`
	ShortLink       types.String          `tfsdk:"short_link"`
	TagID           types.String          `tfsdk:"tag_id"`
	Tags            []tfTypes.TagSchema   `tfsdk:"tags"`
	TenantID        types.String          `tfsdk:"tenant_id"`
	TestCompletedAt types.String          `tfsdk:"test_completed_at"`
	TestStartedAt   types.String          `tfsdk:"test_started_at"`
	TestVariants    []tfTypes.TestVariant `tfsdk:"test_variants"`
	Title           types.String          `tfsdk:"title"`
	TrackConversion types.Bool            `tfsdk:"track_conversion"`
	UpdatedAt       types.String          `tfsdk:"updated_at"`
	URL             types.String          `tfsdk:"url"`
	UserID          types.String          `tfsdk:"user_id"`
	UtmCampaign     types.String          `tfsdk:"utm_campaign"`
	UtmContent      types.String          `tfsdk:"utm_content"`
	UtmMedium       types.String          `tfsdk:"utm_medium"`
	UtmSource       types.String          `tfsdk:"utm_source"`
	UtmTerm         types.String          `tfsdk:"utm_term"`
	Video           types.String          `tfsdk:"video"`
	WebhookIds      []types.String        `tfsdk:"webhook_ids"`
	WorkspaceID     types.String          `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *LinkDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_link"
}

// Schema defines the schema for the data source.
func (r *LinkDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Link DataSource",

		Attributes: map[string]schema.Attribute{
			"android": schema.StringAttribute{
				Computed:    true,
				Description: `The Android destination URL for the short link for Android device targeting.`,
			},
			"archived": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the short link is archived.`,
			},
			"clicks": schema.Float64Attribute{
				Computed:    true,
				Description: `The number of clicks on the short link.`,
			},
			"comments": schema.StringAttribute{
				Computed:    true,
				Description: `The comments for the short link.`,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `The date and time when the short link was created.`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: `The description of the short link. Will be used for Custom Link Previews if ` + "`" + `proxy` + "`" + ` is true.`,
			},
			"do_index": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether to allow search engines to index the short link.`,
			},
			"domain": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"expired_url": schema.StringAttribute{
				Computed:    true,
				Description: `The URL to redirect to when the short link has expired.`,
			},
			"expires_at": schema.StringAttribute{
				Computed:    true,
				Description: `The date and time when the short link will expire in ISO-8601 format.`,
			},
			"external_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `This is the ID of the link in the your database.`,
			},
			"folder_id": schema.StringAttribute{
				Computed:    true,
				Description: `The unique ID of the folder assigned to the short link.`,
			},
			"geo": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"ad": schema.StringAttribute{
						Computed: true,
					},
					"ae": schema.StringAttribute{
						Computed: true,
					},
					"af": schema.StringAttribute{
						Computed: true,
					},
					"ag": schema.StringAttribute{
						Computed: true,
					},
					"ai": schema.StringAttribute{
						Computed: true,
					},
					"al": schema.StringAttribute{
						Computed: true,
					},
					"am": schema.StringAttribute{
						Computed: true,
					},
					"ao": schema.StringAttribute{
						Computed: true,
					},
					"aq": schema.StringAttribute{
						Computed: true,
					},
					"ar": schema.StringAttribute{
						Computed: true,
					},
					"as": schema.StringAttribute{
						Computed: true,
					},
					"at": schema.StringAttribute{
						Computed: true,
					},
					"au": schema.StringAttribute{
						Computed: true,
					},
					"aw": schema.StringAttribute{
						Computed: true,
					},
					"ax": schema.StringAttribute{
						Computed: true,
					},
					"az": schema.StringAttribute{
						Computed: true,
					},
					"ba": schema.StringAttribute{
						Computed: true,
					},
					"bb": schema.StringAttribute{
						Computed: true,
					},
					"bd": schema.StringAttribute{
						Computed: true,
					},
					"be": schema.StringAttribute{
						Computed: true,
					},
					"bf": schema.StringAttribute{
						Computed: true,
					},
					"bg": schema.StringAttribute{
						Computed: true,
					},
					"bh": schema.StringAttribute{
						Computed: true,
					},
					"bi": schema.StringAttribute{
						Computed: true,
					},
					"bj": schema.StringAttribute{
						Computed: true,
					},
					"bl": schema.StringAttribute{
						Computed: true,
					},
					"bm": schema.StringAttribute{
						Computed: true,
					},
					"bn": schema.StringAttribute{
						Computed: true,
					},
					"bo": schema.StringAttribute{
						Computed: true,
					},
					"bq": schema.StringAttribute{
						Computed: true,
					},
					"br": schema.StringAttribute{
						Computed: true,
					},
					"bs": schema.StringAttribute{
						Computed: true,
					},
					"bt": schema.StringAttribute{
						Computed: true,
					},
					"bv": schema.StringAttribute{
						Computed: true,
					},
					"bw": schema.StringAttribute{
						Computed: true,
					},
					"by": schema.StringAttribute{
						Computed: true,
					},
					"bz": schema.StringAttribute{
						Computed: true,
					},
					"ca": schema.StringAttribute{
						Computed: true,
					},
					"cc": schema.StringAttribute{
						Computed: true,
					},
					"cd": schema.StringAttribute{
						Computed: true,
					},
					"cf": schema.StringAttribute{
						Computed: true,
					},
					"cg": schema.StringAttribute{
						Computed: true,
					},
					"ch": schema.StringAttribute{
						Computed: true,
					},
					"ci": schema.StringAttribute{
						Computed: true,
					},
					"ck": schema.StringAttribute{
						Computed: true,
					},
					"cl": schema.StringAttribute{
						Computed: true,
					},
					"cm": schema.StringAttribute{
						Computed: true,
					},
					"cn": schema.StringAttribute{
						Computed: true,
					},
					"co": schema.StringAttribute{
						Computed: true,
					},
					"cr": schema.StringAttribute{
						Computed: true,
					},
					"cu": schema.StringAttribute{
						Computed: true,
					},
					"cv": schema.StringAttribute{
						Computed: true,
					},
					"cw": schema.StringAttribute{
						Computed: true,
					},
					"cx": schema.StringAttribute{
						Computed: true,
					},
					"cy": schema.StringAttribute{
						Computed: true,
					},
					"cz": schema.StringAttribute{
						Computed: true,
					},
					"de": schema.StringAttribute{
						Computed: true,
					},
					"dj": schema.StringAttribute{
						Computed: true,
					},
					"dk": schema.StringAttribute{
						Computed: true,
					},
					"dm": schema.StringAttribute{
						Computed: true,
					},
					"do": schema.StringAttribute{
						Computed: true,
					},
					"dz": schema.StringAttribute{
						Computed: true,
					},
					"ec": schema.StringAttribute{
						Computed: true,
					},
					"ee": schema.StringAttribute{
						Computed: true,
					},
					"eg": schema.StringAttribute{
						Computed: true,
					},
					"eh": schema.StringAttribute{
						Computed: true,
					},
					"er": schema.StringAttribute{
						Computed: true,
					},
					"es": schema.StringAttribute{
						Computed: true,
					},
					"et": schema.StringAttribute{
						Computed: true,
					},
					"fi": schema.StringAttribute{
						Computed: true,
					},
					"fj": schema.StringAttribute{
						Computed: true,
					},
					"fk": schema.StringAttribute{
						Computed: true,
					},
					"fm": schema.StringAttribute{
						Computed: true,
					},
					"fo": schema.StringAttribute{
						Computed: true,
					},
					"fr": schema.StringAttribute{
						Computed: true,
					},
					"ga": schema.StringAttribute{
						Computed: true,
					},
					"gb": schema.StringAttribute{
						Computed: true,
					},
					"gd": schema.StringAttribute{
						Computed: true,
					},
					"ge": schema.StringAttribute{
						Computed: true,
					},
					"gf": schema.StringAttribute{
						Computed: true,
					},
					"gg": schema.StringAttribute{
						Computed: true,
					},
					"gh": schema.StringAttribute{
						Computed: true,
					},
					"gi": schema.StringAttribute{
						Computed: true,
					},
					"gl": schema.StringAttribute{
						Computed: true,
					},
					"gm": schema.StringAttribute{
						Computed: true,
					},
					"gn": schema.StringAttribute{
						Computed: true,
					},
					"gp": schema.StringAttribute{
						Computed: true,
					},
					"gq": schema.StringAttribute{
						Computed: true,
					},
					"gr": schema.StringAttribute{
						Computed: true,
					},
					"gs": schema.StringAttribute{
						Computed: true,
					},
					"gt": schema.StringAttribute{
						Computed: true,
					},
					"gu": schema.StringAttribute{
						Computed: true,
					},
					"gw": schema.StringAttribute{
						Computed: true,
					},
					"gy": schema.StringAttribute{
						Computed: true,
					},
					"hk": schema.StringAttribute{
						Computed: true,
					},
					"hm": schema.StringAttribute{
						Computed: true,
					},
					"hn": schema.StringAttribute{
						Computed: true,
					},
					"hr": schema.StringAttribute{
						Computed: true,
					},
					"ht": schema.StringAttribute{
						Computed: true,
					},
					"hu": schema.StringAttribute{
						Computed: true,
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"ie": schema.StringAttribute{
						Computed: true,
					},
					"il": schema.StringAttribute{
						Computed: true,
					},
					"im": schema.StringAttribute{
						Computed: true,
					},
					"in": schema.StringAttribute{
						Computed: true,
					},
					"io": schema.StringAttribute{
						Computed: true,
					},
					"iq": schema.StringAttribute{
						Computed: true,
					},
					"ir": schema.StringAttribute{
						Computed: true,
					},
					"is": schema.StringAttribute{
						Computed: true,
					},
					"it": schema.StringAttribute{
						Computed: true,
					},
					"je": schema.StringAttribute{
						Computed: true,
					},
					"jm": schema.StringAttribute{
						Computed: true,
					},
					"jo": schema.StringAttribute{
						Computed: true,
					},
					"jp": schema.StringAttribute{
						Computed: true,
					},
					"ke": schema.StringAttribute{
						Computed: true,
					},
					"kg": schema.StringAttribute{
						Computed: true,
					},
					"kh": schema.StringAttribute{
						Computed: true,
					},
					"ki": schema.StringAttribute{
						Computed: true,
					},
					"km": schema.StringAttribute{
						Computed: true,
					},
					"kn": schema.StringAttribute{
						Computed: true,
					},
					"kp": schema.StringAttribute{
						Computed: true,
					},
					"kr": schema.StringAttribute{
						Computed: true,
					},
					"kw": schema.StringAttribute{
						Computed: true,
					},
					"ky": schema.StringAttribute{
						Computed: true,
					},
					"kz": schema.StringAttribute{
						Computed: true,
					},
					"la": schema.StringAttribute{
						Computed: true,
					},
					"lb": schema.StringAttribute{
						Computed: true,
					},
					"lc": schema.StringAttribute{
						Computed: true,
					},
					"li": schema.StringAttribute{
						Computed: true,
					},
					"lk": schema.StringAttribute{
						Computed: true,
					},
					"lr": schema.StringAttribute{
						Computed: true,
					},
					"ls": schema.StringAttribute{
						Computed: true,
					},
					"lt": schema.StringAttribute{
						Computed: true,
					},
					"lu": schema.StringAttribute{
						Computed: true,
					},
					"lv": schema.StringAttribute{
						Computed: true,
					},
					"ly": schema.StringAttribute{
						Computed: true,
					},
					"ma": schema.StringAttribute{
						Computed: true,
					},
					"mc": schema.StringAttribute{
						Computed: true,
					},
					"md": schema.StringAttribute{
						Computed: true,
					},
					"me": schema.StringAttribute{
						Computed: true,
					},
					"mf": schema.StringAttribute{
						Computed: true,
					},
					"mg": schema.StringAttribute{
						Computed: true,
					},
					"mh": schema.StringAttribute{
						Computed: true,
					},
					"mk": schema.StringAttribute{
						Computed: true,
					},
					"ml": schema.StringAttribute{
						Computed: true,
					},
					"mm": schema.StringAttribute{
						Computed: true,
					},
					"mn": schema.StringAttribute{
						Computed: true,
					},
					"mo": schema.StringAttribute{
						Computed: true,
					},
					"mp": schema.StringAttribute{
						Computed: true,
					},
					"mq": schema.StringAttribute{
						Computed: true,
					},
					"mr": schema.StringAttribute{
						Computed: true,
					},
					"ms": schema.StringAttribute{
						Computed: true,
					},
					"mt": schema.StringAttribute{
						Computed: true,
					},
					"mu": schema.StringAttribute{
						Computed: true,
					},
					"mv": schema.StringAttribute{
						Computed: true,
					},
					"mw": schema.StringAttribute{
						Computed: true,
					},
					"mx": schema.StringAttribute{
						Computed: true,
					},
					"my": schema.StringAttribute{
						Computed: true,
					},
					"mz": schema.StringAttribute{
						Computed: true,
					},
					"na": schema.StringAttribute{
						Computed: true,
					},
					"nc": schema.StringAttribute{
						Computed: true,
					},
					"ne": schema.StringAttribute{
						Computed: true,
					},
					"nf": schema.StringAttribute{
						Computed: true,
					},
					"ng": schema.StringAttribute{
						Computed: true,
					},
					"ni": schema.StringAttribute{
						Computed: true,
					},
					"nl": schema.StringAttribute{
						Computed: true,
					},
					"no": schema.StringAttribute{
						Computed: true,
					},
					"np": schema.StringAttribute{
						Computed: true,
					},
					"nr": schema.StringAttribute{
						Computed: true,
					},
					"nu": schema.StringAttribute{
						Computed: true,
					},
					"nz": schema.StringAttribute{
						Computed: true,
					},
					"om": schema.StringAttribute{
						Computed: true,
					},
					"pa": schema.StringAttribute{
						Computed: true,
					},
					"pe": schema.StringAttribute{
						Computed: true,
					},
					"pf": schema.StringAttribute{
						Computed: true,
					},
					"pg": schema.StringAttribute{
						Computed: true,
					},
					"ph": schema.StringAttribute{
						Computed: true,
					},
					"pk": schema.StringAttribute{
						Computed: true,
					},
					"pl": schema.StringAttribute{
						Computed: true,
					},
					"pm": schema.StringAttribute{
						Computed: true,
					},
					"pn": schema.StringAttribute{
						Computed: true,
					},
					"pr": schema.StringAttribute{
						Computed: true,
					},
					"ps": schema.StringAttribute{
						Computed: true,
					},
					"pt": schema.StringAttribute{
						Computed: true,
					},
					"pw": schema.StringAttribute{
						Computed: true,
					},
					"py": schema.StringAttribute{
						Computed: true,
					},
					"qa": schema.StringAttribute{
						Computed: true,
					},
					"re": schema.StringAttribute{
						Computed: true,
					},
					"ro": schema.StringAttribute{
						Computed: true,
					},
					"rs": schema.StringAttribute{
						Computed: true,
					},
					"ru": schema.StringAttribute{
						Computed: true,
					},
					"rw": schema.StringAttribute{
						Computed: true,
					},
					"sa": schema.StringAttribute{
						Computed: true,
					},
					"sb": schema.StringAttribute{
						Computed: true,
					},
					"sc": schema.StringAttribute{
						Computed: true,
					},
					"sd": schema.StringAttribute{
						Computed: true,
					},
					"se": schema.StringAttribute{
						Computed: true,
					},
					"sg": schema.StringAttribute{
						Computed: true,
					},
					"sh": schema.StringAttribute{
						Computed: true,
					},
					"si": schema.StringAttribute{
						Computed: true,
					},
					"sj": schema.StringAttribute{
						Computed: true,
					},
					"sk": schema.StringAttribute{
						Computed: true,
					},
					"sl": schema.StringAttribute{
						Computed: true,
					},
					"sm": schema.StringAttribute{
						Computed: true,
					},
					"sn": schema.StringAttribute{
						Computed: true,
					},
					"so": schema.StringAttribute{
						Computed: true,
					},
					"sr": schema.StringAttribute{
						Computed: true,
					},
					"ss": schema.StringAttribute{
						Computed: true,
					},
					"st": schema.StringAttribute{
						Computed: true,
					},
					"sv": schema.StringAttribute{
						Computed: true,
					},
					"sx": schema.StringAttribute{
						Computed: true,
					},
					"sy": schema.StringAttribute{
						Computed: true,
					},
					"sz": schema.StringAttribute{
						Computed: true,
					},
					"tc": schema.StringAttribute{
						Computed: true,
					},
					"td": schema.StringAttribute{
						Computed: true,
					},
					"tf": schema.StringAttribute{
						Computed: true,
					},
					"tg": schema.StringAttribute{
						Computed: true,
					},
					"th": schema.StringAttribute{
						Computed: true,
					},
					"tj": schema.StringAttribute{
						Computed: true,
					},
					"tk": schema.StringAttribute{
						Computed: true,
					},
					"tl": schema.StringAttribute{
						Computed: true,
					},
					"tm": schema.StringAttribute{
						Computed: true,
					},
					"tn": schema.StringAttribute{
						Computed: true,
					},
					"to": schema.StringAttribute{
						Computed: true,
					},
					"tr": schema.StringAttribute{
						Computed: true,
					},
					"tt": schema.StringAttribute{
						Computed: true,
					},
					"tv": schema.StringAttribute{
						Computed: true,
					},
					"tw": schema.StringAttribute{
						Computed: true,
					},
					"tz": schema.StringAttribute{
						Computed: true,
					},
					"ua": schema.StringAttribute{
						Computed: true,
					},
					"ug": schema.StringAttribute{
						Computed: true,
					},
					"um": schema.StringAttribute{
						Computed: true,
					},
					"us": schema.StringAttribute{
						Computed: true,
					},
					"uy": schema.StringAttribute{
						Computed: true,
					},
					"uz": schema.StringAttribute{
						Computed: true,
					},
					"va": schema.StringAttribute{
						Computed: true,
					},
					"vc": schema.StringAttribute{
						Computed: true,
					},
					"ve": schema.StringAttribute{
						Computed: true,
					},
					"vg": schema.StringAttribute{
						Computed: true,
					},
					"vi": schema.StringAttribute{
						Computed: true,
					},
					"vn": schema.StringAttribute{
						Computed: true,
					},
					"vu": schema.StringAttribute{
						Computed: true,
					},
					"wf": schema.StringAttribute{
						Computed: true,
					},
					"ws": schema.StringAttribute{
						Computed: true,
					},
					"xk": schema.StringAttribute{
						Computed: true,
					},
					"ye": schema.StringAttribute{
						Computed: true,
					},
					"yt": schema.StringAttribute{
						Computed: true,
					},
					"za": schema.StringAttribute{
						Computed: true,
					},
					"zm": schema.StringAttribute{
						Computed: true,
					},
					"zw": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `Geo targeting information for the short link in JSON format ` + "`" + `{[COUNTRY]: https://example.com }` + "`" + `. Learn more: https://d.to/geo`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `The unique ID of the short link.`,
			},
			"image": schema.StringAttribute{
				Computed:    true,
				Description: `The image of the short link. Will be used for Custom Link Previews if ` + "`" + `proxy` + "`" + ` is true.`,
			},
			"ios": schema.StringAttribute{
				Computed:    true,
				Description: `The iOS destination URL for the short link for iOS device targeting.`,
			},
			"key": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The key of the link to retrieve. E.g. for ` + "`" + `d.to/github` + "`" + `, the key is ` + "`" + `github` + "`" + `.`,
			},
			"last_clicked": schema.StringAttribute{
				Computed:    true,
				Description: `The date and time when the short link was last clicked.`,
			},
			"leads": schema.Float64Attribute{
				Computed:    true,
				Description: `The number of leads the short links has generated.`,
			},
			"partner_id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the partner the short link is associated with.`,
			},
			"password": schema.StringAttribute{
				Computed:    true,
				Description: `The password required to access the destination URL of the short link.`,
			},
			"program_id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the program the short link is associated with.`,
			},
			"project_id": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: `This will be removed in a future release, please migrate away from it as soon as possible`,
				Description:        `The project ID of the short link. This field is deprecated – use ` + "`" + `workspaceId` + "`" + ` instead.`,
			},
			"proxy": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the short link uses Custom Link Previews feature.`,
			},
			"public_stats": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the short link's stats are publicly accessible.`,
			},
			"qr_code": schema.StringAttribute{
				Computed:    true,
				Description: `The full URL of the QR code for the short link (e.g. ` + "`" + `https://api.dub.co/qr?url=https://dub.sh/try` + "`" + `).`,
			},
			"rewrite": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the short link uses link cloaking.`,
			},
			"sale_amount": schema.Float64Attribute{
				Computed:    true,
				Description: `The total dollar amount of sales the short links has generated (in cents).`,
			},
			"sales": schema.Float64Attribute{
				Computed:    true,
				Description: `The number of sales the short links has generated.`,
			},
			"short_link": schema.StringAttribute{
				Computed:    true,
				Description: `The full URL of the short link, including the https protocol (e.g. ` + "`" + `https://dub.sh/try` + "`" + `).`,
			},
			"tag_id": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: `This will be removed in a future release, please migrate away from it as soon as possible`,
				Description:        `The unique ID of the tag assigned to the short link. This field is deprecated – use ` + "`" + `tags` + "`" + ` instead.`,
			},
			"tags": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"color": schema.StringAttribute{
							Computed:    true,
							Description: `The color of the tag.`,
						},
						"id": schema.StringAttribute{
							Computed:    true,
							Description: `The unique ID of the tag.`,
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: `The name of the tag.`,
						},
					},
				},
				Description: `The tags assigned to the short link.`,
			},
			"tenant_id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the tenant that created the link inside your system. If set, it can be used to fetch all links for a tenant.`,
			},
			"test_completed_at": schema.StringAttribute{
				Computed:    true,
				Description: `The date and time when the tests were or will be completed.`,
			},
			"test_started_at": schema.StringAttribute{
				Computed:    true,
				Description: `The date and time when the tests started.`,
			},
			"test_variants": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"percentage": schema.Float64Attribute{
							Computed: true,
						},
						"url": schema.StringAttribute{
							Computed: true,
						},
					},
				},
				Description: `An array of A/B test URLs and the percentage of traffic to send to each URL.`,
			},
			"title": schema.StringAttribute{
				Computed:    true,
				Description: `The title of the short link. Will be used for Custom Link Previews if ` + "`" + `proxy` + "`" + ` is true.`,
			},
			"track_conversion": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether to track conversions for the short link.`,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `The date and time when the short link was last updated.`,
			},
			"url": schema.StringAttribute{
				Computed:    true,
				Description: `The destination URL of the short link.`,
			},
			"user_id": schema.StringAttribute{
				Computed:    true,
				Description: `The user ID of the creator of the short link.`,
			},
			"utm_campaign": schema.StringAttribute{
				Computed:    true,
				Description: `The UTM campaign of the short link.`,
			},
			"utm_content": schema.StringAttribute{
				Computed:    true,
				Description: `The UTM content of the short link.`,
			},
			"utm_medium": schema.StringAttribute{
				Computed:    true,
				Description: `The UTM medium of the short link.`,
			},
			"utm_source": schema.StringAttribute{
				Computed:    true,
				Description: `The UTM source of the short link.`,
			},
			"utm_term": schema.StringAttribute{
				Computed:    true,
				Description: `The UTM term of the short link.`,
			},
			"video": schema.StringAttribute{
				Computed:    true,
				Description: `The custom link preview video (og:video). Will be used for Custom Link Previews if ` + "`" + `proxy` + "`" + ` is true. Learn more: https://d.to/og`,
			},
			"webhook_ids": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `The IDs of the webhooks that the short link is associated with.`,
			},
			"workspace_id": schema.StringAttribute{
				Computed:    true,
				Description: `The workspace ID of the short link.`,
			},
		},
	}
}

func (r *LinkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Dub)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.Dub, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *LinkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *LinkDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request, requestDiags := data.ToOperationsGetLinkInfoRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Links.GetLinkInfo(ctx, *request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.LinkSchema != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedLinkSchema(ctx, res.LinkSchema)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
