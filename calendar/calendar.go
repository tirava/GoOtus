/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 22.10.2019 22:44
 * Copyright (c) 2019 - Eugene Klimov
 */

//go:generate protoc --go_out=. calendar.proto

// Package calendar implements simple event calendar via protobuf.
package calendar

//func main() {
//
//	events := newDB(&dbMapEvents{}).(*dbMapEvents)
//
//	event1 := newEvent()
//
//	out, err := proto.Marshal(event1)
//	if err != nil {
//		log.Fatalln("Failed to encode event:", err)
//	}
//
//	event1 = &Event{}
//	if err := proto.Unmarshal(out, event1); err != nil {
//		log.Fatalln("Failed to parse event:", err)
//	}
//
//	event1.Location = "qqqqqqqqqqqqqqqqqqqqqq"
//	_ = events.addEvent(event1)
//
//	event2 := newEvent()
//	event2.Subject = "222222222222222222222"
//	event2.Body = "3333333333333333333"
//	_ = events.addEvent(event2)
//
//	fmt.Println(events.getAllEvents())
//	fmt.Println("--------------------------")
//
//	_ = events.delEvent(event1.Id)
//	fmt.Println(events.getAllEvents())
//	fmt.Println("--------------------------")
//
//	event2.User.Email = []string{"zzzzzzzzzzzzzzzz", "xxxxxxxxxxxxxxxxx"}
//	_ = events.editEvent(event2)
//	fmt.Println(events.getAllEvents())
//	fmt.Println("--------------------------")
//
//	fmt.Println(events.getEvent(2))
//	fmt.Println("--------------------------")
//}
