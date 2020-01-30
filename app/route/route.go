package route

import (
	"net/http"

	"github.com/MrVxCo/go-http-server/app/controller"

	"github.com/gorilla/mux"
)

//LoadRoutes instanciates all the routes and sends to the server
func LoadRoutes() http.Handler {
	return routes()
}

//Add your routes in the routes function
func routes() http.Handler {
	r := mux.NewRouter()

	//serves static files in the static folder...(.js, .css)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer((http.Dir("static")))))

	//Add http routes here
	r.HandleFunc("/", controller.Routetest).Methods("GET")
	// r.HandleFunc("/login", controller.GetLoginHandler).Methods("GET")

	return r
}
