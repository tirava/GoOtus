Feature: Adding event into database via service API
	As API client of API service
	In order to understand that an event was added into database
	I want to get the event via the service API

	Scenario: Adding event into database via service API
		When I send CreateEvent with EventRequest to service API
		Then added event will be returned by GetEvent with id of the event
		And GetError has no errors in both cases

	Scenario: Getting the event from database by id of the event via service API
		When I send GetEvent request with event id to service API
		Then I get EventResponse with id of the event
		And GetError has no errors

	Scenario: Getting non existing event from database by id via service API
    		When I send GetEvent request with non existing event id to service API
    		Then I get EventResponse with error code 'Event not found'
