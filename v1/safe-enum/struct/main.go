package main

import (
	"fmt"

	"example.com/safe-enum/role"
	. "example.com/safe-enum/role"
)

type User struct {
	role role.Role
}

func CreateUser(role Role) (User, error) {
	fmt.Println("Creating user with role", role)
	return User{role}, nil
}

func main() {
	doCreateUser := func(msg string, role Role) {
		fmt.Println(msg)
		fmt.Printf("role = «%v», role == Unknown = %v\n", role, role == Unknown)
		user, err := CreateUser(role)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("user =", user)
		}
		fmt.Println()
	}
	doCreateUser("Oops, creating an empty/zero Role{}", Role{})
	doCreateUser("Creating a guest", Guest)
	doCreateUser("Creating a member", Member)
	doCreateUser("Creating a moderator", Moderator)
	doCreateUser("Creating an admin", Admin)

	doFromString := func(s string) {
		role, err := role.FromString(s)
		if err != nil {
			fmt.Printf("Illegal role '%s', err=«%v» role=«%v»\n", s, err, role)
		} else {
			doCreateUser(fmt.Sprintf("Creating a role from string '%s'", s), role)
		}
	}
	doFromString("guest")
	doFromString("user")
	doFromString("foo")
	doFromString("")
}
