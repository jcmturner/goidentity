package goidentity

import "time"

type User struct {
	authenticated   bool
	domain          string
	userName        string
	displayName     string
	email           string
	human           bool
	groupMembership map[string]bool
	authTime        time.Time
}

func NewUser(username string) User {
	return User{
		userName:        username,
		groupMembership: make(map[string]bool),
	}
}

func (u *User) UserName() string {
	return u.userName
}

func (u *User) Domain() string {
	return u.domain
}

func (u *User) DisplayName() string {
	if u.displayName == "" {
		return u.userName
	}
	return u.displayName
}

func (u *User) Human() bool {
	return u.human
}

func (u *User) AuthTime() time.Time {
	return u.authTime
}

func (u *User) AuthzAttributes() []string {
	s := make([]string, len(u.groupMembership))
	i := 0
	for a := range u.groupMembership {
		s[i] = a
		i++
	}
	return s
}

func (u *User) Authenticated() bool {
	return u.authenticated
}

func (u *User) AddAuthzAttribute(a string) {
	if enabled, ok := u.groupMembership[a]; ok && !enabled {
		u.groupMembership[a] = true
	}
	u.groupMembership[a] = true
}

func (u *User) RemoveAuthzAttribute(a string) {
	if _, ok := u.groupMembership[a]; !ok {
		return
	}
	delete(u.groupMembership, a)
}

func (u *User) EnableAuthzAttribute(a string) {
	if enabled, ok := u.groupMembership[a]; ok && !enabled {
		u.groupMembership[a] = true
	}
}

func (u *User) DisableAuthzAttribute(a string) {
	if enabled, ok := u.groupMembership[a]; ok && enabled {
		u.groupMembership[a] = false
	}
}

func (u *User) Authorized(a string) bool {
	if enabled, ok := u.groupMembership[a]; ok && enabled {
		return true
	}
	return false
}
