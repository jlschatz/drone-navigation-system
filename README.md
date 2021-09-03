# drone-navigation-system

    drone-navigation-system is a REST API for the location retrieval of a drone out in the field. 

## Installation

    docker build -t drone .

## Run the app

    docker run -p 8080:8080 drone

    Please note: If Docker is not used and the binary is executed directly, the following two environment variables will be required:

        export BASE_URL=/api/v1
        export SECTOR_ID=<any random number>

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

## Confirm that API is active

### Request

`GET /ping`

    curl http://localhost:8080/ping

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Content-Type: text/plain; charset=utf-8
    Content-Length: 4

    pong

## Retrive Swagger documentation

### Request

`GET /ping`

    curl http://localhost:8080/swagger

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Content-Type: text/plain; charset=utf-8
    Content-Length: 4

    swagger: '2.0'
    info:
    description: Drone navigation system interface
    version: 1.0.0
    title: DNS API
    basePath: /api/v1
    tags:
    - name: DNS
        description: Drone navigation system interface contacted via this swagger
        externalDocs:
        description: Find out more
        url: 
    paths:
    /:
        post:
        summary: Retrieve drone location
        description: Retrieve drone location via API call
        operationId: GetLocationRequest
        consumes:
            - application/json
        produces:
            - application/json
        parameters:
            - in: body
            name: GetLocationRequest
            required: true
            schema:
                $ref: '#/definitions/GetLocationRequest'
        responses:
            '200':
            description: 'Location retrieved'
            '400':
            description: 'Invalid input, object invalid'
            '500':
            description: 'Service failure'
    definitions:
    GetLocationResponse:
        type: object
        properties:
        loc:
            type: float64
            example: 434.00
    GetLocationRequest:
        type: object
        properties:
        x:
            type: float64
            example: 34.0
        y:
            type: float64
            example: 767.00
        z:
            type: float64
            example: 45.0
        vel:
            type: float64
            example: 123.0


## DevOps

    The Docker container is build via a CircleCI pipline integrated with Github. The container is then pushed to a private Docker registry called Treescale.
    Unit tests are also executed during this pipeline run and will fail the build if all tests do not pass.
    The configuration for this pipeline can be viewed in the .circleci/config.yml file



