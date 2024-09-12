package entity

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// Order Describes - Order entity in Database
type Order struct {
	ID         int
	UserID     uuid.UUID `json:"user_id"`
	Number     string    `json:"number"`
	Status     string    `json:"status"`
	Accrual    int       `json:"accrual"`
	UploadedAt time.Time `json:"uploaded_at"`
}

// OrderDTO Describes - Order Data Transfer Objects
type OrderDTO struct {
	UserID     uuid.UUID `json:"user_id"`
	Number     string    `json:"number"`
	Status     string    `json:"status"`
	Accrual    int       `json:"accrual"`
	UploadedAt time.Time `json:"uploaded_at"`
}
