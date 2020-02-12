package middleware

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

type negroniWrapper struct {
	CustomMiddleware []http.HandlerFunc
}

//LoadMiddlewares adds the custom middlewares onto the negroni
//middleware stack
func (mw *negroniWrapper) LoadMiddlewares() *negroni.Negroni {
	middlewares := negroni.New() //or negroni.Classic() to add negroni default middleware stack
	for _, m := range mw.CustomMiddleware {
		middlewares.UseHandler(m)

	}
	return middlewares
}

func InitCustomMiddlewares() *negroniWrapper {
	// declare your middlewares here
	mw1 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Testing middleware one : mw1")
	})

	mw2 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Testing middleware two : mw2")
	})

	negroniWrapper := &negroniWrapper{
		CustomMiddleware: []http.HandlerFunc{mw1, mw2},
	}

	return negroniWrapper

}
