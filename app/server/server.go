package server

import (
	"fmt"
	"net/http"

	"github.com/MrVxCo/go-http-server/app/middleware"
	"github.com/MrVxCo/go-http-server/app/route"
)

//StartServer starts the go server
func StartServer() {
	r := route.LoadRoutes()
	mw := middleware.InitCustomMiddlewares().LoadMiddlewares()
	mw.UseHandler(r)

	PORT := "2200"

	fmt.Println("server Started on port: " + PORT)
	http.ListenAndServe(":"+PORT, mw)
}
