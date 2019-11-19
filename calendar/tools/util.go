/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 04.11.2019 12:18
 * Copyright (c) 2019 - Eugene Klimov
 */

package tools

import (
	"github.com/google/uuid"
	"time"
)

// IDString2UUIDorNil returns UUID from string or UUID Nil if error.
func IDString2UUIDorNil(id string) uuid.UUID {
	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil
	}
	return uid
}

// DayString2TimeOrNil return date parse of string date
func DayString2TimeOrNil(date string) time.Time {
	d, _ := time.Parse("2006-01-02", date)
	return d
}
