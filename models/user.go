package models

import "strconv"

// User defines an user
type User struct {
	ID   int
	Name string
	Age  int
	Sex  string
}

// NewUser creates user instance
func NewUser(name string, age int, sex string) *User {
	return &User{
		Name: name,
		Age:  age,
		Sex:  sex,
	}
}

func (u *User) String() string {
	return u.Name + "(" + strconv.Itoa(u.Age) + ")"
}
