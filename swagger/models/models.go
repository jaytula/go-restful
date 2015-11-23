package models

type Authorizations map[string][]ScopeObject

type ScopeObject struct {
	Scope       string `json:"scope"`
	Description string `json:"description"`
}
