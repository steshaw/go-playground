package role

import (
	"fmt"
)

type Role struct {
	slug string
}

func (r Role) String() string {
	return r.slug
}

var (
	Unknown   = Role{""}
	Guest     = Role{"guest"}
	Member    = Role{"member"}
	Moderator = Role{"moderator"}
	Admin     = Role{"admin"}
)

func FromString(s string) (Role, error) {
	switch s {
	case Guest.slug:
		return Guest, nil
	case Member.slug:
		return Member, nil
	case Moderator.slug:
		return Moderator, nil
	case Admin.slug:
		return Admin, nil
	}

	return Unknown, fmt.Errorf("unknown role: «%s»", s)
}

func Roles() []Role {
	return []Role{Guest, Member, Moderator, Admin}
}
