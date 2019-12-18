/*
 * HomeWork-9: Integration tests
 * Created on 16.12.2019 12:10
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

func parseDateTime(s, layout string) *timestamp.Timestamp {
	t, err := time.Parse(layout, s)
	if err != nil {
		return nil
	}
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil
	}
	return ts
}

func parseDuration(s string) *duration.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		return nil
	}
	dr := ptypes.DurationProto(d)
	return dr
}
