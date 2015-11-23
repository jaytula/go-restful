package models

// 5.1.4 Authorizations Object
type AuthorizationsObject map[string]AuthorizationObject

// 5.1.5
type AuthorizationObject struct {
	Type       string      `json:"type"`
	PassAs     string      `json:"passAs"`
	Keyname    string      `json:"keyname"`
	Scopes     []Scope     `json:"scopes,omitempty"`
	GrantTypes []GrantType `json:"grantTypes,omitempty"`
}

// 5.1.6, 5.2.11
type Scope struct {
	// Required. The name of the scope.
	Scope string `json:"scope"`
	// Recommended. A short description of the scope.
	Description string `json:"description"`
}

// 5.1.7
type GrantType struct {
	Implicit          Implicit          `json:"implicit"`
	AuthorizationCode AuthorizationCode `json:"authorization_code"`
}

type Authorizations map[string][]ScopeObject

type ScopeObject struct {
	Scope       string `json:"scope"`
	Description string `json:"description"`
}

// 5.1.8 Implicit Object
type Implicit struct {
	// Required. The login endpoint definition.
	loginEndpoint LoginEndpoint `json:"loginEndpoint"`
	// An optional alternative name to standard "access_token" OAuth2 parameter.
	TokenName string `json:"tokenName"`
}

// 5.1.9 Authorization Code Object
type AuthorizationCode struct {
	TokenRequestEndpoint TokenRequestEndpoint `json:"tokenRequestEndpoint"`
	TokenEndpoint        TokenEndpoint        `json:"tokenEndpoint"`
}

// 5.1.10 Login Endpoint Object
type LoginEndpoint struct {
	// Required. The URL of the authorization endpoint for the implicit grant flow. The value SHOULD be in a URL format.
	Url string `json:"url"`
}

// 5.1.11 Token Request Endpoint Object
type TokenRequestEndpoint struct {
	// Required. The URL of the authorization endpoint for the authentication code grant flow. The value SHOULD be in a URL format.
	Url string `json:"url"`
	// An optional alternative name to standard "client_id" OAuth2 parameter.
	ClientIdName string `json:"clientIdName"`
	// An optional alternative name to the standard "client_secret" OAuth2 parameter.
	ClientSecretName string `json:"clientSecretName"`
}

// 5.1.12 Token Endpoint Object
type TokenEndpoint struct {
	// Required. The URL of the token endpoint for the authentication code grant flow. The value SHOULD be in a URL format.
	Url string `json:"url"`
	// An optional alternative name to standard "access_token" OAuth2 parameter.
	TokenName string `json:"tokenName"`
}
