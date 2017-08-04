package goidentity

import "time"

type Identity interface {
	UserName() string
	Domain() string
	DisplayName() string
	Human() bool
	AuthTime() time.Time
	AuthzAttributes() []string
	AddAuthzAttribute(a string)
	RemoveAuthzAttribute(a string)
	Authenticated() bool
	Authorized(a string) bool
}
