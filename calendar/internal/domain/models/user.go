/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 28.10.2019 21:40
 * Copyright (c) 2019 - Eugene Klimov
 */

package models

import (
	"github.com/satori/go.uuid"
	"time"
)

// User is the base user's struct.
type User struct {
	ID       uuid.UUID
	Name     string
	Email    []string
	Mobile   []string
	Birthday time.Time
}

// NewUser returns new user struct.
// TODO get events for specific user
//func NewUser() User {
//	return User{
//		ID:       uuid.NewV4(),
//		Name:     "qqq",
//		Email:    []string{"www"},
//		Mobile:   []string{"+777"},
//		Birthday: time.Now(),
//	}
//}
