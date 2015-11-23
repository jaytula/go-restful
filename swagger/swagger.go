// Package swagger implements the structures of the Swagger
// https://github.com/wordnik/swagger-spec/blob/master/versions/1.2.md
package swagger

import "github.com/jaytula/go-restful/swagger/models"

const swaggerVersion = "1.2"

// 4.3.3 Data Type Fields
type DataTypeFields struct {
	Type         *string  `json:"type,omitempty"` // if Ref not used
	Ref          *string  `json:"$ref,omitempty"` // if Type not used
	Format       string   `json:"format,omitempty"`
	DefaultValue Special  `json:"defaultValue,omitempty"`
	Enum         []string `json:"enum,omitempty"`
	Minimum      string   `json:"minimum,omitempty"`
	Maximum      string   `json:"maximum,omitempty"`
	Items        *Item    `json:"items,omitempty"`
	UniqueItems  *bool    `json:"uniqueItems,omitempty"`
}

type Special string

// 4.3.4 Items Object
type Item struct {
	Type   *string `json:"type,omitempty"`
	Ref    *string `json:"$ref,omitempty"`
	Format string  `json:"format,omitempty"`
}

// 5.1 Resource Listing
type ResourceListing struct {
	SwaggerVersion string                      `json:"swaggerVersion"` // e.g 1.2
	Apis           []Resource                  `json:"apis"`
	ApiVersion     string                      `json:"apiVersion"`
	Info           Info                        `json:"info"`
	Authorizations models.AuthorizationsObject `json:"authorizations,omitempty"`
}

// 5.1.2 Resource Object
type Resource struct {
	Path        string `json:"path"` // relative or absolute, must start with /
	Description string `json:"description"`
}

// 5.1.3 Info Object
type Info struct {
	Title             string `json:"title"`
	Description       string `json:"description"`
	TermsOfServiceUrl string `json:"termsOfServiceUrl,omitempty"`
	Contact           string `json:"contact,omitempty"`
	License           string `json:"license,omitempty"`
	LicenseUrl        string `json:"licenseUrl,omitempty"`
}

// 5.2 API Declaration
type ApiDeclaration struct {
	SwaggerVersion string                `json:"swaggerVersion"`
	ApiVersion     string                `json:"apiVersion"`
	BasePath       string                `json:"basePath"`
	ResourcePath   string                `json:"resourcePath"` // must start with /
	Apis           []Api                 `json:"apis,omitempty"`
	Models         ModelList             `json:"models,omitempty"`
	Produces       []string              `json:"produces,omitempty"`
	Consumes       []string              `json:"consumes,omitempty"`
	Authorizations models.Authorizations `json:"authorizations,omitempty"`
}

// 5.2.2 API Object
type Api struct {
	Path        string      `json:"path"` // relative or absolute, must start with /
	Description string      `json:"description"`
	Operations  []Operation `json:"operations,omitempty"`
}

// 5.2.3 Operation Object
type Operation struct {
	DataTypeFields
	Method           string                `json:"method"`
	Summary          string                `json:"summary,omitempty"`
	Notes            string                `json:"notes,omitempty"`
	Nickname         string                `json:"nickname"`
	Authorizations   models.Authorizations `json:"authorizations,omitempty"`
	Parameters       []Parameter           `json:"parameters"`
	ResponseMessages []ResponseMessage     `json:"responseMessages,omitempty"` // optional
	Produces         []string              `json:"produces,omitempty"`
	Consumes         []string              `json:"consumes,omitempty"`
	Deprecated       string                `json:"deprecated,omitempty"`
}

// 5.2.4 Parameter Object
type Parameter struct {
	DataTypeFields
	ParamType     string `json:"paramType"` // path,query,body,header,form
	Name          string `json:"name"`
	Description   string `json:"description"`
	Required      bool   `json:"required"`
	AllowMultiple bool   `json:"allowMultiple"`
}

// 5.2.5 Response Message Object
type ResponseMessage struct {
	Code          int    `json:"code"`
	Message       string `json:"message"`
	ResponseModel string `json:"responseModel,omitempty"`
}

// 5.2.6, 5.2.7 Models Object
type Model struct {
	Id            string            `json:"id"`
	Description   string            `json:"description,omitempty"`
	Required      []string          `json:"required,omitempty"`
	Properties    ModelPropertyList `json:"properties"`
	SubTypes      []string          `json:"subTypes,omitempty"`
	Discriminator string            `json:"discriminator,omitempty"`
}

// 5.2.8 Properties Object
type ModelProperty struct {
	DataTypeFields
	Description string `json:"description,omitempty"`
}
