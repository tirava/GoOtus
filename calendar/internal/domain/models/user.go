/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 28.10.2019 21:40
 * Copyright (c) 2019 - Eugene Klimov
 */

package models

import (
	"time"
)

// User is the base user's struct.
type User struct {
	ID       int
	Name     string
	Email    []string
	Mobile   []string
	Birthday time.Time
}
