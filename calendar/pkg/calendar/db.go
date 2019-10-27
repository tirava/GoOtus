/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 23.10.2019 19:36
 * Copyright (c) 2019 - Eugene Klimov
 */

package calendar

// Constants
const (
	MapDBType = "map"
	//PostgresDBType = "postgres"
)

// DB is thw main interface for any DBs
type DB interface {
	AddEvent(event Event) error
	EditEvent(event Event) error
	DelEvent(id uint32) error
	GetEvent(id uint32) (Event, error)
	GetAllEvents() []Event
}

// DBTypes struct helps for getting new db by type
type DBTypes struct {
	MapDB *DBMapEvents
	//PostgresDB *DBPostgresEvents
}

// NewDB returns DB by db type
func NewDB(dbType string) *DBTypes {
	switch dbType {
	case MapDBType:
		return &DBTypes{MapDB: newMapDB()}
		//case PostgresDBType:
		//return &DBType{postgresDB:newPostgresDB()}
	}
	return nil
}
