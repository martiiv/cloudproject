# CloudProject

Welcome to the project for the PROG2005 Cloud Technologies course. 

This project is made by Group 2: *implement group name here*.

In this project the focus is on the use of a REST web application created with the use of third party APIs, and the
creation and use of webhooks. It is created with the programming language Golang and then deployed to a server using
OpenStack **SHOULD WE USE DOCKER?**, which provides the client with the possibility to retrieve information about a 
travel plan using car. Here the client can get information about *route*, *nearby filling- and EV stations*, *traffic 
incidents and flow*. **UPDATE IF MORE INFORMATION IS USED**


<h1>Project Report</h1>

<h3>Orginal project plan</h3>
<h4>Startup</h4>

Our initial idea was to make an entertainment hub allowing a user to access information about movies, tv-shows, music, and comics. We wanted to implment webhooks and firestore letting the user register which movie/music genres and get notified when new entertainment gets released. Sadly we encountered a problem when trying to implment the music api in golang. We decided that a workaround would be time consuming and decided to change topic. 

<h4>The road-trip-planner</h4>

We decided on a road-trip travel planner, where we give a user information about the route they are going to travel, for instance where the nearest charging stations are, the shortest path from one destination to another, it will calculate a time of departure depending on the weather conditions. It will also provide the user with points of interest if desired. We had to change the "gps" api of the project from the Statens vegvesen API to the tomtom api since Statens Vegvesen did not respond to our email and since their login functionality did not work. 

<h4>Technologies</h4>

The technologies that we decided to implement were Open Stack, Firebase, Cloud functions and Docker. We have been able to implement the wanted functions to our route planner such as finding the route to the users' destination, find EV stations,
filling stations, avoid traffic, describe the weather and points of interest a long your route such as restaurants, hotels and road attractions. 

**Our final product has almost all the functionality we sought out to implement skriv her når koden er ferdig ALEKS OG TORMOD PLIS HJÆLP ikkeno filter for bensin og diesel**

<h3>Reflection</h3>

There was a lot of back and forth in the early stages of our project. No one had any specific ideas in mind which resulted in us having a hard time collectively deciding on something. This on top of the problems with different APIs, keys and paywalls. Once we had decided to go for the route planner we had problems with the location api we used however we quickly found another location api providing us with accurate coordinates. 

One of the things the group benefited from was starting the project early, arranging meetings and discussing topics. Our decision to start coding early on benefited us when trying to find good api's. When we decided to use an api we started poking it immediately, sending requests and understanding it's structure. This made troubleshooting early on easy since we were able to find problems with the api's fast. 

We held frequent meetings with the entire group present on to two times a week in order to maintain an overview of the project, what people were working on, problems encountere and to discuss intended functionality. To maintain a solid project structure, we decided to give each member a git branch to not overwrite eachothers code in addition to separate different aspects of the project and prevent buggy/breaking code in the master branch. We also used git issues and the issue board frequently to document workflow. By using this board each group member could easily check the status of the project. Creating the different endpoints we needed for the implementation of our application went really well since we thoroughly tested each api before implementation.

If we were to do to something differently in our project we would have chosen to use milestones. By using milestones we would have been able to track the overall progression of the project. Since the milestone tool in GitLab, groups issues corresponding to the project. Thus, it would have been easier to set different time periods on the different issues to when they needed to be complete. This would have given us a better overview of the different parts of the project that needs to be completed.  

We would also have more informative meeting summaries where each group member documented what they have been doing since the last session in addition to defining what parts of the project are missing. Splitting workload properly and clearly defining tasks for each member using our issue tracker is also something we would have done. We also think using test driven development would have benefited us in terms of time, project structure and general scope in the project.
**Mer her alle?**

Our biggest struggle throughout the project has been communication in terms of functionality, clearly defining what an endpoint is supposed to do and what it currently does has been hard and prevented us from getting a clear overview of the project making the final stages of development hectic and confusing.  

There have been some other hard aspects in the project: Firstly we struggled with webhook implementation, clearly defining what features we wanted it to have, when to invoke and notify. Furthermore we tried to implement a client in our project which also resulted in some difficulties, our problem was that in order to get a the notification, the user had to refresh the provided url. We concluded that this was not very user-friendly, and annoying. Therefore, we removed it and instead implemented Slack, now the user gets a Slack notification with messages given that they provide a Slack url.

Another hard aspect of the project was the problem of api calls. Due to the way we designed api calls, the api had to call the location api in order to get coordinates for almost all of our endpoints. However we were able to implement caching in the end by creating an additional collection in firebase containing previously used locations. The unit-testing has also been difficult, we wanted to test the webhook invoke and to do so we needed to skip the sleep function which was hard. We  also struggled with implementing Docker and needed to do some research on Dockerfiles and Docker-compose in order to implement this properly. Another hard aspect of the project was how to deal with personal traces in the code, every group member has their own way of coding. Implementing functions differently, and uses different variable names. Therefore, it can be hard sometimes to understand what other group members have done without sufficent commenting.  

<h3>Learning outcome</h3>

We have become better at collaborating on programming tasks with eachother. The GitLab issues has been a great tool for the communication in our group. By creating issues in Git and labeling, assigning them to group members. We have learned that good group communication is very important in order to be able to finish a project like this one, we also became better at using GitLabs different tools like issues, branching, WIKI implementation, merge requests and labels. 

We found that having a good foundation with hard policies and principles is important for project work, especially in our field. Not having a good grasp of our project will result in alot of extra work and confusion. Frequently meeting and discussing the state of the project is crucial for delivering a working application. We have also learned a lot in regards to programming, especially the webhooks and Docker implementation. Expanding on our knowledge of webhooks, how to invoke, cache and structure code to use this technology properly. Learning how to implement Docker, using best practice for Docker compose and Dockerfiles has been challenging but nevertheless rewarding. 
 
<h3>Total work hours dedicated to the project cumulatively by the group:</h3>
The total hours the group has worked on the project has been 78 hours. 

___
### The REST services in use are:

***Related route API:***
- https://openrouteservice.org/

***Traffic news- and filling station API:***
- https://developer.tomtom.com

***Position related API:***
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
The ***Traffic news***-endpoint focuses on returning the traffic news based on the start and end location.
The user is able to enter their destination to the Position API, which then sends their longitude and latitude
to the TrafficNews-API.
From there the client is able to get a detailed description about incidents, slow traffic, stationary traffic and road construction. 
### Request
Main request method:
```
Method: GET
Path: /roadTripPlanner/v1/traffic/{location1}/{location2}/
```

We find the use of an alternative request method necessary due to the possibility of not being on an address accepted by
the Position API. Therefore, the user will be able to manually enter their destination-coordinates.

`{location1}` refers to the location you are travelling from, which can be an address, place name or attraction(Eks: Slottsplassen 1, Washington DC or Eiffel Tower)
provided by the ***Implement the API link here***.

`{location2}` refers to the location you are arriving to, which can be an address, place name or attraction (Eks: Slottsplassen 1, Washington DC or Eiffel Tower)
provided by the ***Implement the API link here***.

Example request 1: `/roadTripPlanner/v1/traffic/Gjøvik/Lillehammer`


The ***FillingStations***-endpoint focuses on returning filling stations based on the start and end location.
The user is able to enter their destination to the Position API, which then sends their longitude and latitude
to the FillingStations-API.
From there the client is able to get a detailed description about filling stations near you during your trip.
### Request
Main request method:
```
Method: GET
Path: /roadTripPlanner/v1/fuel/{fuelType}/{location}/
```

We find the use of an alternative request method necessary due to the possibility of not being on an address accepted by
the Position API. Therefore, the user will be able to manually enter their destination-coordinates.

`{fuelType}` refers to the fuel type you want, which can be diesel or petrol.
provided by the ***Implement the API link here***.

`{location}` refers to the location you are now, which can be an address, place name or attraction (Eks: Slottsplassen 1, Washington DC or Eiffel Tower)
provided by the ***Implement the API link here***.

Example request 1: `/roadTripPlanner/v1/fuel/s`

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
