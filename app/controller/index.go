package controller

import (
	"net/http"
)

//Routetest a test handler function for the / endpoint.
//feel free to modify as needed.
func Routetest(w http.ResponseWriter, r *http.Request) {
	// log.Println(r)

	w.Write([]byte("route is functioning"))
}
