package main

type UserBuilder interface {
	SetName(n string) UserBuilder
	SetGender(g string) UserBuilder
	SetAge(n int) UserBuilder
	Build() *User
}

type userBuilder struct {
	kind   Kind
	Name   string
	Gender string
	Age    int
}

func (u *userBuilder) SetName(n string) UserBuilder {
	u.Name = n
	return u
}

func (u *userBuilder) SetGender(g string) UserBuilder {
	u.Gender = g
	return u
}

func (u *userBuilder) SetAge(n int) UserBuilder {
	u.Age = n
	return u
}

func (u *userBuilder) Build() *User {
	return &User{
		Kind:   u.kind,
		Name:   u.Name,
		Age:    u.Age,
		Gender: u.Gender,
	}
}

func NewUserBuilder(kind Kind) UserBuilder {
	return &userBuilder{
		kind: kind,
	}
}
