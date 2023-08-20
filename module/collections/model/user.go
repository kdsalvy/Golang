package model

import (
	"errors"
	"fmt"
)

// User model
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers returns all the users
func GetUsers() []*User {
	return users
}

// AddUser adds a new user
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user id should be nil or 0")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserByID to get user by Id
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("user with id '%v' not found", id)
}
