/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 28.10.2019 21:40
 * Copyright (c) 2019 - Eugene Klimov
 */

package models

import (
	"github.com/google/uuid"
)

// User is the base user's struct.
type User struct {
	ID     uuid.UUID
	Name   string
	Email  string
	Mobile string
}

// NewUser returns new user struct.
//func NewUser() User {
//	return User{
//		ID:       uuid.NewV4(),
//		Name:     "qqq",
//		Email:    "www",
//		Mobile:   "+777",
//	}
//}
