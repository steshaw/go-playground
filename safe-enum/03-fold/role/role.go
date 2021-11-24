package role

import (
	"fmt"
)

// The Role algebra.
type roleAlg struct {
	member    func()
	moderator func()
	guest     func()
	admin     func()
}

type Role func(roleAlg)

var (
	Unknown   = Role(nil)
	Guest     = Role(func(alg roleAlg) { alg.guest() })
	Member    = Role(func(alg roleAlg) { alg.member() })
	Moderator = Role(func(alg roleAlg) { alg.moderator() })
	Admin     = Role(func(alg roleAlg) { alg.admin() })
)

func (r Role) String() string {
	if r == nil {
		return ""
	}
	var result string
	r(roleAlg{
		member:    func() { result = "member" },
		moderator: func() { result = "moderator" },
		guest:     func() { result = "guest" },
		admin:     func() { result = "admin" },
	})
	return result
}

func (r Role) Eq(r2 Role) bool {
	if r == nil || r2 == nil {
		return r == nil && r2 == nil
	}
	var result bool
	t := func() { result = true }
	f := func() { result = false }
	r(roleAlg{
		member: func() {
			r2(roleAlg{
				member:    t,
				moderator: f,
				guest:     f,
				admin:     f,
			})
		},
		moderator: func() {
			r2(roleAlg{
				member:    f,
				moderator: t,
				guest:     f,
				admin:     f,
			})
		},
		guest: func() {
			r2(roleAlg{
				member:    f,
				moderator: f,
				guest:     t,
				admin:     f,
			})
		},
		admin: func() {
			r2(roleAlg{
				member:    f,
				moderator: f,
				guest:     f,
				admin:     t,
			})
		},
	})
	return result
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

func Roles() []Role {
	return []Role{Guest, Member, Moderator, Admin}
}
