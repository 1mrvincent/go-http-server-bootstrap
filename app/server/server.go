package server

import (
	"fmt"
	"net/http"

	"app/route"
)

//StartServer starts the go server
func StartServer() {
	r := route.LoadRoutes()

	PORT := "2200"

	fmt.Println("server Started on port: " + PORT)
	http.ListenAndServe(":"+PORT, r)
}
