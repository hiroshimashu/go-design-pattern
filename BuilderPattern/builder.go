package main

type UserBuilder interface {
	SetName(n string) UserBuilder
	SetGender(g string) UserBuilder
	SetAge(n int) UserBuilder
	Build() *User
}

type userBuilder struct {
	kind
}
