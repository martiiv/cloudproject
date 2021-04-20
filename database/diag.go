package database

import (
	"fmt"
	"net/http"
)

func diag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movieDB int

	res, err := http.Get("https://api.themoviedb.org/")
	if err != nil {
		http.Error(w, "Cannot connet", http.StatusInternalServerError)
		movieDB = res.StatusCode
	} else {
		movieDB = res.StatusCode
	}

	fmt.Fprintf(w, `{
		"movieDB": "%v",
   		"version": "v1"
 		}`, movieDB)

}
