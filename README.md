# CloudProject

Welcome to the project for the PROG2005 Cloud Technologies course. 

This project is made by Group 2: *implement group name here*.

In this project the focus is on the use of a REST web application created with the use of third party APIs, and the
creation and use of webhooks. It is created with the programming language Golang and then deployed to a server using
OpenStack, which provides the client with the possibility to retrieve information about movies, music or comics.
___
###The REST services in use are:

***Related movies API:***
- https://tastedive.com/read/api

***Movie metadata:***
- https://tivovideometadata.api-docs.io/v3/content-enrichment/id-lookup

***Spotify related API:***
- https://developer.spotify.com/documentation/web-api/reference/#endpoint-get-track

***Lyrics API (might be used):***
- https://lyricsovh.docs.apiary.io/#reference/0/lyrics-of-a-song/search?console=1

***IMDB Movie API***
- https://rapidapi.com/blog/movie-api/
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

##Movies
The ***Movies***-endpoint focuses on returning movies based on the **title** of the movie, or a specific **actor**.
The user is able to search for a specific movie based on title, and get information/metadata about that movie. 
When the user searches by an actor, all movies the actor has appeared in will be returned.

Further the user are able to ask to get similar movies to the search result, which will return information about those movies.
Lastly the user are able to get movie providers and where they can see the chosen movie.

### Request
```
Method: GET
Path: /{root_endpoint_path}/v1/{endpoint_1}/{movie_title}/
```

```
Method: GET
Path: /{root_endpoint_path}/v1/{endpoint_1}/{movie_actor}/
```

`{movie_title}` refers to the English name of the movie provided by the ***Implement the API link here***.

`{movie_actor}` refers to the English name of the actor provided by the ***Implement the API link here***.

Example request 1: `/{root_endpoint_path}/v1/movies/Bad Boys 2/` \
Example request 1: `/{root_endpoint_path}/v1/movies/Will Smith/` 

### Response
***Implement response***

##Music
### Request
***Implement request***

### Response
***Implement response***

##Comics
### Request
***Implement request***

### Response
***Implement response***
