# file: features/previewer.feature

# http://http-server:8080
# http://image-server

Feature: Image Previewer
	As http client of image previewer service
	In order to understand that the user got resized and cropped image from source image
	I want to receive resized and cropped image from original source
	And it must be cached on service

	Scenario: Previewer service is available
		When I send "GET" request to "http://http-server:8080/hello/"
		Then The response code should be 200
		And The response should match text "Hello, my name is Previewer!"

	Scenario: Image server is available
    		When I send "GET" request to "http://image-server/"
    		Then The response code should be 200
    		And The response should match text "Image Server for Image-Previewer!"

	Scenario: Invalid request params
    	When I send "GET" request to "http://http-server:8080/preview/"
    	Then The response code should be 400
    	And I receive error with text "invalid request params"

	Scenario: Invalid preview size
    	When I send "GET" request to "http://http-server:8080/preview/" with data "3OO/20I/image-server/_gopher_original_1024x504.jpg"
    	Then The response code should be 400
    	And I receive error with text "invalid request preview size"

	Scenario: Image not found
    	When I send "GET" request to "http://http-server:8080/preview/" with data "300/200/image-server/_gopher_original_1024x504.tiff"
    	Then The response code should be 404
    	And I receive error with text "invalid image source"

    Scenario: 1st image preview (not from cache)
        	When I send "GET" request to "http://http-server:8080/preview/" with data "300/200/image-server/_gopher_original_1024x504.jpg"
        	Then The response code should be 200
        	# And I receive image with size "300x200"
        	And I received header "From-Cache" = "false"

    Scenario: 2nd image preview (from cache)
        	When I send "GET" request to "http://http-server:8080/preview/" with data "400/300/image-server/_gopher_original_1024x504.jpg"
        	Then The response code should be 200
        	# And I receive image with size "400x300"
        	And I received header "From-Cache" = "true"