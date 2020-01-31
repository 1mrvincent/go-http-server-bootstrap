package server

import (
	"fmt"
	"net/http"

	"github.com/MrVxCo/go-http-server/app/route"
	"github.com/urfave/negroni"
)

//StartServer starts the go server
func StartServer() {
	r := route.LoadRoutes()
	middlewares := negroni.New()
	// middlewares.Use()
	middlewares.UseHandler(r)

	PORT := "2200"

	fmt.Println("server Started on port: " + PORT)
	http.ListenAndServe(":"+PORT, r)
}
