package entity

import uuid "github.com/satori/go.uuid"

// CurrentUserID - Describes Context Value key
var CurrentUserID uuid.UUID

// User Describes - User entity in Database
type User struct {
	ID       uuid.UUID `json:"uuid"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
}

// UserDTO Describes - User Data Transfer Objects
type UserDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
