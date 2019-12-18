Feature: List events from database via service API for day, week and month
	As API client of API service
	In order to understand that business logic for list events works expected
	I want to get events list via the service API for day, week and month

	Scenario: Adding events into database via service API
		When I send CreateEvent to service API for cycle with 5 events for same user and step 3 days for OccursAt
		Then all 5 added events will be returned by GetUserEvents for given user
		And GetError has no errors in these cases

	Scenario: List events for day via service API
		When I send GetEventsForDay request with current day to service API
        Then I get EventsResponse with 1 event in it with OccursAt in current day
        And GetError returns no errors

    Scenario: List events for week via service API
    	When I send GetEventsForWeek request with current day to service API
        Then I get EventsResponse with 3 events in it with OccursAt in near week
        And GetError returns no errors

    Scenario: List events for month via service API
        When I send GetEventsForMonth request with current day to service API
        Then I get EventsResponse with 5 events in it with OccursAt in near month
        And GetError returns no errors