package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {

	log.Println("Listening on port: " + getPort())
	handlers()
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func handlers() {
	http.HandleFunc("charge", evStations)
	log.Println(http.ListenAndServe(getPort(), nil))
}
