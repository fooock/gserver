package user

import "time"

// User of the program
type User struct {
	UUID     string    `json:"uuid"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

// Users is a list of User
type Users []User

// Register is the struct used to register new users
type Register struct {
	UUID     string    `json:"uuid"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Created  time.Time `json:"created"`
}

// Login is the struct to login
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
