package repository

import (
	"fmt"
)

type UniqueFieldErr struct {
	Value string
	Err   error
}

func (e *UniqueFieldErr) Error() string {
	return fmt.Sprintf("value: %s, already exist.", e.Value)
}

func NewUniqueFieldErr(v string, err error) error {
	return &UniqueFieldErr{
		Value: v,
		Err:   err,
	}
}

type NotFountErr struct {
	Entity string
	Field  string
	Value  string
	Err    error
}

func (e *NotFountErr) Error() string {
	return fmt.Sprintf("%s with %s: %s not found", e.Entity, e.Field, e.Value)
}

func NewNotFountErr(entity, field, value string, err error) error {
	return &NotFountErr{
		Entity: entity,
		Field:  field,
		Value:  value,
		Err:    err,
	}
}

type ConflictErr struct {
	Value string
	Err   error
}

func (e *ConflictErr) Error() string {
	return fmt.Sprintf("value: %s, already uploaded by another user", e.Value)
}

func NewConflictErr(v string, err error) error {
	return &ConflictErr{
		Value: v,
		Err:   err,
	}
}
