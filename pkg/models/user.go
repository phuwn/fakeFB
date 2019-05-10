package models

import "github.com/phuwn/fakeFB/pkg/util/errs"

// User object
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate func validate user info
func (u *User) Validate() error {
	if u.Email == "" || u.Password == "" {
		return errs.Invalid
	}
	return nil
}

// Response struct
type Response struct {
	Success string `json:"success"`
}
