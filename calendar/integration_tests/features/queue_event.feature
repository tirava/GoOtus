Feature: Put alerting events into message queue and get it for sending to users
	As API client of API service
	In order to understand that alerting events are sending to users
	I want to create event with alerting OccursAt
	And get correct user id from message queue for feature sending message

	Scenario: Adding alerting event into database via service API for test user
		When I CreateEvent with EventRequest to service API with OccursAt = Now
		Then added event will be scheduled into message queue
		And GetError has no error

	Scenario: Get correct event from message queue via queue API for test user
		When I consume message queue
        Then I get event with correct test user id
        And will be ready to send message rot this user
