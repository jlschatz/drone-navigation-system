# drone-navigation-system

Foobar is a Python library for dealing with word pluralization.

## Installation

    docker build -t drone .

## Run the app

    docker run -p 8080:8080 drone

## Run the tests

    go test

# REST API

    The REST API to the Drone-Navigation-System app is described below.
    Swagger doumentation has also been included for ease of use.

## Retrieve drones current location

### Request

`POST /api/v1/loc`

    curl -d '{"x":"23.00","y":"326.0","z":"78.00","vel":"656.0"}' -H "Content-Type: application/json" -X POST http://localhost:8080/api/v1/loc

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Content-Type: application/json
    Content-Length: 20

    {
        "Location": 7282.05
    }
