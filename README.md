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

# Extra:
- If snowing - notify 15 minutes before planned take-off. If not, sleep and notify at take-off.
- 

#Project report:
###Original project plan:
Brief description of the original project plan, and what has and has not been achieved/changed in the final product.

Our first plan was to make an entertainment service to access information about movies, Spotify music, and comics. 
In addition, the user should be able to register webhooks to get notified about new movies, Spotify music or comics, of their interest. 
There were some problems with the use of the Spotify API together with GO, so we had to change the topic. We then decided to make a 
"car" travel planner, where we give a user information about the route they are going to travel, for instance where the nearest charging 
stations are, the shortest path from one destination to another, possibility to get notification about weather conditions or car accidents on 
their route when these are registered for instance, and also more information (some of these are given that we get access to the Norwegian 
Public Roads Administration (Statens vegvesen) API). We had to change the topic of the project again from Statens vegvesen API to Open route service API, 
developer API and position stack API. Because we were not able to get the Statens vegvesen API. We therefore landed on creating a route planner 
that will help you find the best route to your destinations. At the same time, it will help you find EV stations, filling stations, avoid traffic/road work,
describe the weather, restaurants, hotels and roads attractions on your route to your destination. 

The technology that we were planning to use were Open Stack, Firebase and Cloud functions. We ended up using that technology, and we were also able 
to implement Docker. We have been able to implement the wanted functions to our route planner such as finding the route to the users' destination, find EV stations,
filling stations, avoid traffic, describe the weather and points of interest a long your route such as restaurants, hotels and road attractions. 

###Reflection of what went well and what went wrong with the project:
There was a lot of back and forth in the beginning of our project, when it comes to deciding the project topic of our project. The problem was that there was 
a few problems with different APIs we need to be able to create the wanted functionalities for our early plans on project topics. Once we had decided to go for 
the route planner we had problems with the geo API, but luckily it went very well to change to another API. 
**Ealy start, coded early to find api problems, met frequently to work and discuss, use of git branches**



If we were to do to something different with our project we would have chosen to use the GitLab tool milestone. 
By using milestones we would have been able to track better our issues. Since the milestone tool groups issues that corresponds to a project, 
feature or time period. Thus, it would be easier to set different time periods on the different issues to when they needed to be complete and this would
have given us a better overview of the different parts of the project that needs to be completed. 
**More frequent issue updates, proper meeting summary discussion, split workload properly, document more (what we tried what failed when coding), decide on topic earlier, use a more test driven approach when coding**

Communication within the group is something that went well in the project. Every group member were good at creating issues in GitLab and 
thus it made it easy for everyone to see what has been implemented and what has not been done yet. **Nja egt ikke?** Creating the different endpoints we needed
for the implementation of our application went really well. The learning outcome of the project part of the course has been rewarding for us in getting a better 
understanding of what we have learned during the course. In the project part we have had to use a lot of the knowledge we got from 
assignment 1 and assignment 2, and we have been able to improve our skills and knowledge in parts like webhooks and error handling due to working on this more in the project. 
Another thing that also worked well was using branches in Gitlab and pushing and pulling the different work the group members had done. And at the end 
merging it all together.

###Reflection on the hard aspects of the project:
The hard aspects of the project has been a few things. First we struggled with the webhook implementation. The hard part has been to understand what we wanted the users to get notifications on and how to 
implement this to our project. Further we tried to implement client API to our project which also resulted in some difficulties, the problem of using a client was that the user had to refresh the site to receive
notifications. This is not very user-friendly, and annoying for the user. Therefore, we changed it, so that the user gets a Slack notification with messages. Then the user will get the messages in
time on their phone, computer or clock. Another hard aspect of the project was the caching. Due to the file structure in our project the persistence of location data was problematic to implement, 
but we were able to implement it in the end by fixing the file structure. The unit-testing has also been difficult because we wanted to test the webhook invoke, and to do so we needed to ship over 
the sleep function which was hard finding a solution to. We have also struggled with implementing Docker, and needed to do some research on that topic. We studied the examples on Docker
provided in the course, to get a better understanding of how to use it. 

Another hard part of the project was that everyone has their own way of coding. We all implement functions differently and uses different syntax. Therefore, it can be hard sometimes to understand what
other group members have done. Thus, we have been strict on the commenting part of code quality, so that it is easy for everyone else to understand different functions and variables group members 
have made. 




###What new has the group learned?:
We have become better in collaborating on programming tasks with other students. The GitLab issues has been a great tool for the communication in our group.
By creating issues in Git and labeling and assigning them to group members, we have been able to see what people are working on. We have also been 
able to see what needs to be done and what is finished. We have written meeting summaries for every meeting we have had during the project in Cloud.
This has given us the possibility to read a short summary of what we did the last time and what needs to be done. These summaries are provided in 
the project WIKI. We have learned that good group communication is very important for us to be able to finish the project, and we also became
better in using GitLabs different tools like issues. 

We have also learned a lot more programming especially the webhooks implementation part has been where we have learned the most during this project. 
 **Skriv om unit testing, docker og docker compose, using gitlab, caching data from api's**
 
###Total work hours dedicated to the project cumulatively by the group:
The total hours the group has worked on the project has been hours. 
