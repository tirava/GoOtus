/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 22.10.2019 22:44
 * Copyright (c) 2019 - Eugene Klimov
 */

//go:generate protoc --go_out=. calendar.proto

// Package calendar implements simple event calendar via protobuf.
package calendar
