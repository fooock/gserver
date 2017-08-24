package user

import "time"

// User of the program
type User struct {
	UUID     string    `json:"uuid"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json_updated:"updated"`
}
