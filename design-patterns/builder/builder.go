package main

import (
	"github.com/pkg/errors"
)

type User struct {
	Name string
	Age  int
}
type UserBuilder struct {
	User
	err error
}

func defaultUser() User {
	return User{
		Name: "",
		Age:  0,
	}
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		User: defaultUser(),
	}
}

func (b *UserBuilder) WithName(n string) *UserBuilder {
	if b.err != nil {
		return b
	}
	if n == "" {
		b.err = errors.New("name is empty")
	}
	b.Name = n
	return b
}

func (b *UserBuilder) WithAge(a int) *UserBuilder {
	if b.err != nil {
		return b
	}
	if a == 0 {
		b.err = errors.New("age is empty")
	}
	b.Age = a
	return b
}

func (b *UserBuilder) Builder() (*User, error) {
	if b.err != nil {
		return nil, b.err
	}
	return &b.User, nil
}
