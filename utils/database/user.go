package database

import (
	"fmt"
	"regexp"
	"time"
)

type RestoreUser struct {
	id           int
	username     string
	firstName    string
	lastName     string
	role         string
	email        string
	password     string
	createdAt    time.Time
	lastModified time.Time
}

//Validate an email address from string
func isEmail(email string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	return emailRegex.MatchString(email)
}

//Create a new restore user
func (db *DB) AddUser(user RestoreUser) error {
	//Validate user
	if len(user.username) < 1 || len(user.username) > 255 {
		return fmt.Errorf("invalid username")
	}
	if len(user.lastName) > 255 {
		return fmt.Errorf("invalid lastname")
	}
	if len(user.firstName) > 255 {
		return fmt.Errorf("invalid firstname")
	}
	if len(user.role) > 255 {
		return fmt.Errorf("invalid role")
	}

	//Validate email
	if !isEmail(user.email) {
		return fmt.Errorf("invalid email")
	}

	//Passwords have to contain a number, letter and symbol at least, with a length more than 7 but less than 255
	if len(user.password) > 7 && len(user.password) < 255 {
		return fmt.Errorf("invalid password")
	}

	//Create a hashed version of the password to store in the database

	return nil
}

func (db *DB) FetchUser() {

}

func (db *DB) UpdateUser() {

}
