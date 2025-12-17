package models

import "time"

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	DOB  string `json:"dob"  validate:"required,datetime=2006-01-02"`
}

type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age,omitempty"`
}

type User struct {
	ID   int
	Name string
	DOB  time.Time
	Age  int
}
