package goidentity

import "time"

type User struct {
	Authenticated bool
	Domain          string
	UserName        string
	DisplayName     string
	Email           string
	Human           bool
	GroupMembership map[string]bool
	AuthTime        time.Time
}

func (u *User) UserName() string {
	return u.UserName
}

func (u *User) Domain() string {
	return u.Domain
}

func (u *User) DisplayName() string {
	if u.DisplayName == "" {
		return u.UserName
	}
	return u.DisplayName
}

func (u *User) Human() bool {
	return u.Human
}

func (u *User) AuthTime() time.Time {
	return u.AuthTime
}

func (u *User) AuthzAttributes() []string {
	s := make([]string, len(u.GroupMembership))
	i := 0
	for a := range u.GroupMembership {
		s[i] = a
		i++
	}
	return s
}

func NewUser(username string) User {
	return User {
		UserName: username,
	}
}

func (u *User) Authenticated() bool {
	return u.Authenticated
}

func (u *User) AddAuthzAttribute(a string) {
	if enabled, ok := u.GroupMembership[a]; ok && !enabled {
			u.GroupMembership[a] = true
	}
	u.GroupMembership[a] = true
}

func (u *User) RemoveAuthzAttribute(a string) {
	if _, ok := u.GroupMembership[a]; !ok {
		return
	}
	delete(u.GroupMembership, a)
}

func (u *User) EnableAuthzAttribute(a string) {
	if enabled, ok := u.GroupMembership[a]; ok && !enabled {
			u.GroupMembership[a] = true
	}
}

func (u *User) DisableAuthzAttribute(a string) {
	if enabled, ok := u.GroupMembership[a]; ok && enabled {
			u.GroupMembership[a] = false
	}
}

func (u *User) Authorized(a string) bool {
	if enabled, ok := u.GroupMembership[a]; ok && enabled {
			return true
	}
	return false
}