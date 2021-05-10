# CloudProject

Welcome to the project for the PROG2005 Cloud Technologies course. 

This project is made by Group 2: *implement group name here*.

In this project the focus is on the use of a REST web application created with the use of third party APIs, and the
creation and use of webhooks. It is created with the programming language Golang and then deployed to a server using
OpenStack **SHOULD WE USE DOCKER?**, which provides the client with the possibility to retrieve information about a 
travel plan using car. Here the client can get information about *route*, *nearby filling- and EV stations*, *traffic 
incidents and flow*. **UPDATE IF MORE INFORMATION IS USED**
___
### The REST services in use are:

***Related route API:***
- https://openrouteservice.org/

***Traffic news- and filling station API:***
- https://developer.tomtom.com

***Position related API:**
- https://developer.mapquest.com

***Weather api***
- https://openweathermap.org/api


## Endpoints
In the application we are using the following endpoints

```
GET
/roadTripPlanner/v1/Route/{start}/{end}/
/roadTripPlanner/v1/fuel/{fuelType}/{location}/
/roadTripPlanner/v1/pitStops/{location}/
/roadTripPlanner/v1/traffic/{location1}/{location2}/
/roadTripPlanner/v1/weather/{location}/
/roadTripPlanner/v1/diag/

POST 
/roadTripPlanner/v1/plan/ 

DELETE
/roadTripPlanner/v1/planner/{id}
```

Running on the server, the endpoints will be as follows:
```
GET

http://OPENSTACKIP/roadTripPlanner/v1/planRoute/{start}/{end}/
http://OPENSTACKIP/roadTripPlanner/v1/fuel/{fuelType}/{location}/
http://OPENSTACKIP/roadTripPlanner/v1/pitStops/{location}/
http://OPENSTACKIP/roadTripPlanner/v1/traffic/{location1}/{location2}/ 
http://OPENSTACKIP/roadTripPlanner/v1/weather/{location}/
http://OPENSTACKIP/roadTripPlanner/v1/diag/

POST 
http://OPENSTACKIP/roadTripPlanner/v1/plan/ 

DELETE
http://OPENSTACKIP/roadTripPlanner/v1/planner/{id}
```

## Route
The ***Routes***-endpoint focuses on returning a travel route based on the start and end location. 
The user is able to enter their destination to the Position API, which then sends their longitude and latitude 
to the Route-API. 
From there the client is able to get a detailed description about which exits to take in the roundabouts,
where to turn left and right

### Request
Main request method:
```
Method: GET
Path: /roadTripPlanner/v1/planRoute/{start}/{end}/
```

We find the use of an alternative request method necessary due to the possibility of not being on an address accepted by
the Position API. Therefore, the user will be able to manually enter their destination-coordinates. 

`{start}` refers to the address, place name or attraction (Eks: Slottsplassen 1, Washington DC or Eiffel Tower)
provided by the ***Implement the API link here***.

`{end}` refers to the address, place name or attraction (Eks: Slottsplassen 1, Washington DC or Eiffel Tower)
provided by the ***Implement the API link here***.

Example request 1: `/roadTripPlanner/v1/route/Gjøvik/Lillehammer` 

### Response
A list of directions
```
{
    "EstimatedArrival": "2021-05-10 09:35:43",
    "LengthKM": 46,
    "Route":
        {
            "Street": "Niels Ødegaards Gate",
            "Maneuver": "Leave.",
            "RoadNumber": "",
            "JunctionType": ""
        },
        {
            "Street": "Niels Ødegaards Gate",
            "Maneuver": "Turn right.",
            "RoadNumber": "",
            "JunctionType": "REGULAR"
        },
        {
            "Street": "Strandgata",
            "Maneuver": "Turn right.",
            "RoadNumber": "",
            "JunctionType": "REGULAR"
        },
        {
            "Street": "Vestre Totenveg",
            "Maneuver": "At the roundabout take the exit on the left.",
            "RoadNumber": "33",
            "JunctionType": "ROUNDABOUT"
        }
}        
```


## Traffic news- and filling stations
### Request
Main request method: 

### Response
***Implement response***

## Notifications
### Request
***Implement request***

### Response
***Implement response***

## Diagnostics interface
### Request
The diagnostics interface indicates the availability of all individual services this service depends on.
The reporting occurs based on status codes returned by the dependent services. The diag interface further provides
information about the number of registered webhooks, and the uptime of the service.

```
Method: GET
Path: /diag/
```

Example request: `/diag/`

### Response
***Implement response***
Body (Example):

```
{
   "tomtom": "200",
   "openrouteservice": "200",
   "openweathermap": "500",
   "mapquest": "200",
   "registered": 4,
   "version": "v1",
   "uptime": 412 seconds
}
```

# Extra:
- If snowing - notify 15 minutes before planned take-off. If not, sleep and notify at take-off.
- 
