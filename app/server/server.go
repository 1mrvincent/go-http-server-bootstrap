package server

import (
	"fmt"
	"net/http"

	"github.com/MrVxCo/go-http-server/app/route"
	"github.com/MrVxCo/go-http-server/app/middleware"
)

//StartServer starts the go server
func StartServer() {
	r := route.LoadRoutes()
	mw := middleware.GetCustomMiddlewares().LoadMiddlewares()
	mw.UseHandler(r)

	PORT := "2200"

	fmt.Println("server Started on port: " + PORT)
	http.ListenAndServe(":"+PORT, mw)
}
