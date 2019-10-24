/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 23.10.2019 19:36
 * Copyright (c) 2019 - Eugene Klimov
 */

package calendar

type db interface {
	newDB() interface{}
	addEvent(event Event) error
	editEvent(event Event) error
	delEvent(id uint32) error
	getEvent(id uint32) (Event, error)
	getAllEvents() []Event
}

func newDB(d db) interface{} {
	return d.newDB()
}
