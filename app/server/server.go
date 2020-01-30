package server

import (
	"fmt"
	"net/http"

	"github.com/MrVxCo/go-http-server/app/route"
)

//StartServer starts the go server
func StartServer() {
	r := route.LoadRoutes()

	PORT := "2200"

	fmt.Println("server Started on port: " + PORT)
	http.ListenAndServe(":"+PORT, r)
}
