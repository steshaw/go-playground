package role

import (
	"errors"
	"fmt"
)

type Role struct {
	role int
}

const (
	unknown   = 0
	guest     = 1
	member    = 2
	moderator = 3
	admin     = 4
)

var (
	Unknown   = Role{unknown}
	Guest     = Role{guest}
	Member    = Role{member}
	Moderator = Role{moderator}
	Admin     = Role{admin}
)

func (r Role) String() string {
	switch r.role {
	case unknown:
		return ""
	case guest:
		return "guest"
	case member:
		return "member"
	case moderator:
		return "moderator"
	case admin:
		return "admin"
	}
	panic("invalid role. Cannot happen?")
}

func FromString(s string) (Role, error) {
	switch s {
	case Guest.String():
		return Guest, nil
	case Member.String():
		return Member, nil
	case Moderator.String():
		return Moderator, nil
	case Admin.String():
		return Admin, nil
	}

	return Unknown, errors.New(fmt.Sprintf("unknown role: «%s»", s))
}
