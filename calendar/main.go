/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 22.10.2019 22:44
 * Copyright (c) 2019 - Eugene Klimov
 */

//go:generate protoc --go_out=. calendar.proto

package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {

	events := newDBEvents()
	event1 := newEvent()

	out, err := proto.Marshal(event1)
	if err != nil {
		log.Fatalln("Failed to encode event:", err)
	}

	event1 = &Event{}
	if err := proto.Unmarshal(out, event1); err != nil {
		log.Fatalln("Failed to parse event:", err)
	}

	event1.Location = "qqqqqqqqqqqqqqqqqqqqqq"
	events.addEvent(event1)

	event2 := newEvent()
	event2.Subject = "222222222222222222222"
	event2.Body = "3333333333333333333"
	events.addEvent(event2)

	fmt.Println(events.events[1])
	fmt.Println(events.events[2])

	events.delEvent(event1)
	fmt.Println(events.events[1])
	fmt.Println(events.events[2])

	event2.User.Email = []string{"zzzzzzzzzzzzzzzz", "xxxxxxxxxxxxxxxxx"}
	events.editEvent(event2)
	fmt.Println(events.events[1])
	fmt.Println(events.events[2])
}
