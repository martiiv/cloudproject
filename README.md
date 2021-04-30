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

***Position related API:***
- https://positionstack.com
___

## Endpoints
In the application we are using ***Implement number of main endpoints here***

```
/{root_endpoint_path}/v1/{endpoint_1}/
/{root_endpoint_path}/v1/{endpoint_2}/
/{root_endpoint_path}/v1/{endpoint_3}/
/{root_endpoint_path}/v1/{endpoint_4}/
```

Running on the server, the endpoints will be as follows:
```
http://{IP_Address}/{root_endpoint_path}/v1/{endpoint_1}/
http://{IP_Address}/{root_endpoint_path}/v1/{endpoint_2}/
http://{IP_Address}/{root_endpoint_path}/v1/{endpoint_3}/
http://{IP_Address}/{root_endpoint_path}/v1/{endpoint_4}/
```

## Route
The ***Routes***-endpoint focuses on returning a travel route based on the longitude and latitude of the start- and end 
location. The user is able to enter their destination to the Position API, which then sends their longitude and latitude 
to the Route-API. From there the client is able to get a detailed description about which exits to take in the roundabouts,
where to turn left and right etc...

### Request
Main request method:
```
Method: GET
Path: /{root_endpoint_path}/v1/{endpoint_1}/{start_address}/{end_address}/
```

Alternative request method:
```
Method: GET
Path: /{root_endpoint_path}/v1/{endpoint_1}/{latitude_start}-{longitude_start}/{latitude_end}-{longitude_end}/
```
We find the use of an alternative request method necessary due to the possibility of not being on an address accepted by
the Position API. Therefore, the user will be able to manually enter their destination-coordinates. 

`{start_address}` refers to the address, place name or attraction (Eks: Slottsplassen 1, Washington DC or Eiffel Tower)
provided by the ***Implement the API link here***.

`{end_address}` refers to the address, place name or attraction (Eks: Slottsplassen 1, Washington DC or Eiffel Tower)
provided by the ***Implement the API link here***.

`{latitude_start}` refers to the latitude of the **start**-destination written in the ISO 6709-format provided by the ***Implement the API link here***.

`{logitude_start}` refers to the longitude of the **start**-destination written in the ISO 6709-format provided by the ***Implement the API link here***.

`{latitude_end}` refers to the latitude of the **end**-destination written in the ISO 6709-format provided by the ***Implement the API link here***.

`{logitude_end}` refers to the longitude of the **end**-destination written in the ISO 6709-format provided by the ***Implement the API link here***.

Example request 1: `/{root_endpoint_path}/v1/route/Gjøvik/Lillehammer` \
Example request 1: `/{root_endpoint_path}/v1/route/60.786489-10.685456/61.122137-10.464437` 

### Response
***Implement response***

## Traffic news- and filling stations
### Request
***Implement request***

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
   "positionstack": "500"
   "tomtom": "200",
   "openrouteservice": "200",
   "registered": 4,
   "version": "v1",
   "uptime": 412 seconds
}
```

#Extra:
- If snowing - notify 15 minutes before planned take-off. If not, sleep and notify at take-off.
- 
