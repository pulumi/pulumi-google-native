// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v4

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new company entity.
type Company struct {
	pulumi.CustomResourceState

	// The URI to employer's career site or careers page on the employer's web site, for example, "https://careers.google.com".
	CareerSiteUri pulumi.StringOutput `pulumi:"careerSiteUri"`
	// Derived details about the company.
	DerivedInfo CompanyDerivedInfoResponseOutput `pulumi:"derivedInfo"`
	// The display name of the company, for example, "Google LLC".
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Equal Employment Opportunity legal disclaimer text to be associated with all jobs, and typically to be displayed in all roles. The maximum number of allowed characters is 500.
	EeoText pulumi.StringOutput `pulumi:"eeoText"`
	// Client side company identifier, used to uniquely identify the company. The maximum number of allowed characters is 255.
	ExternalId pulumi.StringOutput `pulumi:"externalId"`
	// The street address of the company's main headquarters, which may be different from the job location. The service attempts to geolocate the provided address, and populates a more specific location wherever possible in DerivedInfo.headquarters_location.
	HeadquartersAddress pulumi.StringOutput `pulumi:"headquartersAddress"`
	// Set to true if it is the hiring agency that post jobs for other employers. Defaults to false if not provided.
	HiringAgency pulumi.BoolOutput `pulumi:"hiringAgency"`
	// A URI that hosts the employer's company logo.
	ImageUri pulumi.StringOutput `pulumi:"imageUri"`
	// A list of keys of filterable Job.custom_attributes, whose corresponding `string_values` are used in keyword searches. Jobs with `string_values` under these specified field keys are returned if any of the values match the search keyword. Custom field values with parenthesis, brackets and special symbols are not searchable as-is, and those keyword queries must be surrounded by quotes.
	KeywordSearchableJobCustomAttributes pulumi.StringArrayOutput `pulumi:"keywordSearchableJobCustomAttributes"`
	// Required during company update. The resource name for a company. This is generated by the service when a company is created. The format is "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for example, "projects/foo/tenants/bar/companies/baz".
	Name pulumi.StringOutput `pulumi:"name"`
	// The employer's company size.
	Size pulumi.StringOutput `pulumi:"size"`
	// Indicates whether a company is flagged to be suspended from public availability by the service when job content appears suspicious, abusive, or spammy.
	Suspended pulumi.BoolOutput `pulumi:"suspended"`
	// The URI representing the company's primary web site or home page, for example, "https://www.google.com". The maximum number of allowed characters is 255.
	WebsiteUri pulumi.StringOutput `pulumi:"websiteUri"`
}

// NewCompany registers a new resource with the given unique name, arguments, and options.
func NewCompany(ctx *pulumi.Context,
	name string, args *CompanyArgs, opts ...pulumi.ResourceOption) (*Company, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ExternalId == nil {
		return nil, errors.New("invalid value for required argument 'ExternalId'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	var resource Company
	err := ctx.RegisterResource("google-native:jobs/v4:Company", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCompany gets an existing Company resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCompany(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CompanyState, opts ...pulumi.ResourceOption) (*Company, error) {
	var resource Company
	err := ctx.ReadResource("google-native:jobs/v4:Company", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Company resources.
type companyState struct {
}

type CompanyState struct {
}

func (CompanyState) ElementType() reflect.Type {
	return reflect.TypeOf((*companyState)(nil)).Elem()
}

type companyArgs struct {
	// The URI to employer's career site or careers page on the employer's web site, for example, "https://careers.google.com".
	CareerSiteUri *string `pulumi:"careerSiteUri"`
	// The display name of the company, for example, "Google LLC".
	DisplayName string `pulumi:"displayName"`
	// Equal Employment Opportunity legal disclaimer text to be associated with all jobs, and typically to be displayed in all roles. The maximum number of allowed characters is 500.
	EeoText *string `pulumi:"eeoText"`
	// Client side company identifier, used to uniquely identify the company. The maximum number of allowed characters is 255.
	ExternalId string `pulumi:"externalId"`
	// The street address of the company's main headquarters, which may be different from the job location. The service attempts to geolocate the provided address, and populates a more specific location wherever possible in DerivedInfo.headquarters_location.
	HeadquartersAddress *string `pulumi:"headquartersAddress"`
	// Set to true if it is the hiring agency that post jobs for other employers. Defaults to false if not provided.
	HiringAgency *bool `pulumi:"hiringAgency"`
	// A URI that hosts the employer's company logo.
	ImageUri *string `pulumi:"imageUri"`
	// A list of keys of filterable Job.custom_attributes, whose corresponding `string_values` are used in keyword searches. Jobs with `string_values` under these specified field keys are returned if any of the values match the search keyword. Custom field values with parenthesis, brackets and special symbols are not searchable as-is, and those keyword queries must be surrounded by quotes.
	KeywordSearchableJobCustomAttributes []string `pulumi:"keywordSearchableJobCustomAttributes"`
	// Required during company update. The resource name for a company. This is generated by the service when a company is created. The format is "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for example, "projects/foo/tenants/bar/companies/baz".
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The employer's company size.
	Size     *CompanySize `pulumi:"size"`
	TenantId string       `pulumi:"tenantId"`
	// The URI representing the company's primary web site or home page, for example, "https://www.google.com". The maximum number of allowed characters is 255.
	WebsiteUri *string `pulumi:"websiteUri"`
}

// The set of arguments for constructing a Company resource.
type CompanyArgs struct {
	// The URI to employer's career site or careers page on the employer's web site, for example, "https://careers.google.com".
	CareerSiteUri pulumi.StringPtrInput
	// The display name of the company, for example, "Google LLC".
	DisplayName pulumi.StringInput
	// Equal Employment Opportunity legal disclaimer text to be associated with all jobs, and typically to be displayed in all roles. The maximum number of allowed characters is 500.
	EeoText pulumi.StringPtrInput
	// Client side company identifier, used to uniquely identify the company. The maximum number of allowed characters is 255.
	ExternalId pulumi.StringInput
	// The street address of the company's main headquarters, which may be different from the job location. The service attempts to geolocate the provided address, and populates a more specific location wherever possible in DerivedInfo.headquarters_location.
	HeadquartersAddress pulumi.StringPtrInput
	// Set to true if it is the hiring agency that post jobs for other employers. Defaults to false if not provided.
	HiringAgency pulumi.BoolPtrInput
	// A URI that hosts the employer's company logo.
	ImageUri pulumi.StringPtrInput
	// A list of keys of filterable Job.custom_attributes, whose corresponding `string_values` are used in keyword searches. Jobs with `string_values` under these specified field keys are returned if any of the values match the search keyword. Custom field values with parenthesis, brackets and special symbols are not searchable as-is, and those keyword queries must be surrounded by quotes.
	KeywordSearchableJobCustomAttributes pulumi.StringArrayInput
	// Required during company update. The resource name for a company. This is generated by the service when a company is created. The format is "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for example, "projects/foo/tenants/bar/companies/baz".
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The employer's company size.
	Size     CompanySizePtrInput
	TenantId pulumi.StringInput
	// The URI representing the company's primary web site or home page, for example, "https://www.google.com". The maximum number of allowed characters is 255.
	WebsiteUri pulumi.StringPtrInput
}

func (CompanyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*companyArgs)(nil)).Elem()
}

type CompanyInput interface {
	pulumi.Input

	ToCompanyOutput() CompanyOutput
	ToCompanyOutputWithContext(ctx context.Context) CompanyOutput
}

func (*Company) ElementType() reflect.Type {
	return reflect.TypeOf((**Company)(nil)).Elem()
}

func (i *Company) ToCompanyOutput() CompanyOutput {
	return i.ToCompanyOutputWithContext(context.Background())
}

func (i *Company) ToCompanyOutputWithContext(ctx context.Context) CompanyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompanyOutput)
}

type CompanyOutput struct{ *pulumi.OutputState }

func (CompanyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Company)(nil)).Elem()
}

func (o CompanyOutput) ToCompanyOutput() CompanyOutput {
	return o
}

func (o CompanyOutput) ToCompanyOutputWithContext(ctx context.Context) CompanyOutput {
	return o
}

// The URI to employer's career site or careers page on the employer's web site, for example, "https://careers.google.com".
func (o CompanyOutput) CareerSiteUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Company) pulumi.StringOutput { return v.CareerSiteUri }).(pulumi.StringOutput)
}

// Derived details about the company.
func (o CompanyOutput) DerivedInfo() CompanyDerivedInfoResponseOutput {
	return o.ApplyT(func(v *Company) CompanyDerivedInfoResponseOutput { return v.DerivedInfo }).(CompanyDerivedInfoResponseOutput)
}

// The display name of the company, for example, "Google LLC".
func (o CompanyOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Company) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Equal Employment Opportunity legal disclaimer text to be associated with all jobs, and typically to be displayed in all roles. The maximum number of allowed characters is 500.
func (o CompanyOutput) EeoText() pulumi.StringOutput {
	return o.ApplyT(func(v *Company) pulumi.StringOutput { return v.EeoText }).(pulumi.StringOutput)
}

// Client side company identifier, used to uniquely identify the company. The maximum number of allowed characters is 255.
func (o CompanyOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v *Company) pulumi.StringOutput { return v.ExternalId }).(pulumi.StringOutput)
}

// The street address of the company's main headquarters, which may be different from the job location. The service attempts to geolocate the provided address, and populates a more specific location wherever possible in DerivedInfo.headquarters_location.
func (o CompanyOutput) HeadquartersAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Company) pulumi.StringOutput { return v.HeadquartersAddress }).(pulumi.StringOutput)
}

// Set to true if it is the hiring agency that post jobs for other employers. Defaults to false if not provided.
func (o CompanyOutput) HiringAgency() pulumi.BoolOutput {
	return o.ApplyT(func(v *Company) pulumi.BoolOutput { return v.HiringAgency }).(pulumi.BoolOutput)
}

// A URI that hosts the employer's company logo.
func (o CompanyOutput) ImageUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Company) pulumi.StringOutput { return v.ImageUri }).(pulumi.StringOutput)
}

// A list of keys of filterable Job.custom_attributes, whose corresponding `string_values` are used in keyword searches. Jobs with `string_values` under these specified field keys are returned if any of the values match the search keyword. Custom field values with parenthesis, brackets and special symbols are not searchable as-is, and those keyword queries must be surrounded by quotes.
func (o CompanyOutput) KeywordSearchableJobCustomAttributes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Company) pulumi.StringArrayOutput { return v.KeywordSearchableJobCustomAttributes }).(pulumi.StringArrayOutput)
}

// Required during company update. The resource name for a company. This is generated by the service when a company is created. The format is "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for example, "projects/foo/tenants/bar/companies/baz".
func (o CompanyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Company) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The employer's company size.
func (o CompanyOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v *Company) pulumi.StringOutput { return v.Size }).(pulumi.StringOutput)
}

// Indicates whether a company is flagged to be suspended from public availability by the service when job content appears suspicious, abusive, or spammy.
func (o CompanyOutput) Suspended() pulumi.BoolOutput {
	return o.ApplyT(func(v *Company) pulumi.BoolOutput { return v.Suspended }).(pulumi.BoolOutput)
}

// The URI representing the company's primary web site or home page, for example, "https://www.google.com". The maximum number of allowed characters is 255.
func (o CompanyOutput) WebsiteUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Company) pulumi.StringOutput { return v.WebsiteUri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CompanyInput)(nil)).Elem(), &Company{})
	pulumi.RegisterOutputType(CompanyOutput{})
}