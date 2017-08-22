package goidentity

type Authenticator interface {
	Authenticate() (Identity, bool, error)
}
