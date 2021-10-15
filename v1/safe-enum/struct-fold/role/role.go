package role

import (
	"fmt"
)

type Role struct {
	fold func(func(), func(), func(), func())
}

var (
	Unknown = Role{
		fold: nil,
	}
	Guest = Role{
		fold: func(
			member func(),
			moderator func(),
			guest func(),
			admin func(),
		) {
			guest()
		},
	}
	Member = Role{
		fold: func(
			member func(),
			moderator func(),
			guest func(),
			admin func(),
		) {
			member()
		},
	}
	Moderator = Role{
		fold: func(
			member func(),
			moderator func(),
			guest func(),
			admin func(),
		) {
			moderator()
		},
	}
	Admin = Role{
		fold: func(
			member func(),
			moderator func(),
			guest func(),
			admin func(),
		) {
			admin()
		},
	}
)

func (r Role) String() string {
	if r.fold == nil {
		return ""
	}
	var result string
	r.fold(func() {
		result = "member"
	}, func() {
		result = "moderator"
	}, func() {
		result = "guest"
	}, func() {
		result = "admin"
	})
	return result
}

func (r Role) Eq(r2 Role) bool {
	return r.String() == r2.String()
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

	return Unknown, fmt.Errorf("unknown role: «%s»", s)
}
