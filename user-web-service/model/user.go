package model

import (
	"errors"
	"fmt"
)

// User : model struct
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers : to get list of all users
func GetUsers() []*User {
	return users
}

// AddUser : Add new User to list
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New USer must not include id")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserByID : get User by ID
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

// UpdateUser : update user based on input
func UpdateUser(user User) (User, error) {
	for i, userToUpdate := range users {
		if userToUpdate.ID == user.ID {
			users[i] = &user
			return user, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", user.ID)
}

// RemoveUserByID : remove user based on input
func RemoveUserByID(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
