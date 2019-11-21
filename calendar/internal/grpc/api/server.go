/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 28.10.2019 21:50
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package api implements grpc api.
package api

//go:generate protoc --go_out=plugins=grpc:. --proto_path=../../../api api.proto
