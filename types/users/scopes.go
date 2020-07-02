package users

import "strings"

type Scope string

const (
	FullAdmin   Scope = "admin"
	ReadOnly    Scope = "readonly"
	RServices   Scope = "read:services"
	RWServices  Scope = "write:services"
	RIncidents  Scope = "read:incidents"
	RWIncidents Scope = "write:incidents"

	EmptyUser Scope = "none"
)

func namedScope(name string) Scope {
	switch name {
	case "admin":
		return FullAdmin
	case "readonly":
		return ReadOnly
	case "read:services":
		return RServices
	case "write:services":
		return RWServices
	case "read:incidents":
		return RIncidents
	case "write:incidents":
		return RWIncidents
	default:
		return EmptyUser
	}
}

func (u *User) AllScopes() []Scope {
	var scopes []Scope
	for _, s := range strings.Split(u.Scopes, ",") {
		scopes = append(scopes, namedScope(s))
	}
	return scopes
}
