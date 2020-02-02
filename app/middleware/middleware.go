package middleware

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

type NegroniWrapper struct {
	CustomMiddleware []http.HandlerFunc
}

func (mw *NegroniWrapper) LoadMiddlewares() *negroni.Negroni {
	middlewares := negroni.New() //Classic()
	for _, m := range mw.CustomMiddleware {
		middlewares.UseHandler(m)

	}
	return middlewares
}

func GetCustomMiddlewares() *NegroniWrapper {
	mw1 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Testing middleware one : mw1")
	})

	mw2 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Testing middleware two : mw2")
	})

	negroniWrapper := &NegroniWrapper{
		CustomMiddleware: []http.HandlerFunc{mw1, mw2},
	}

	return negroniWrapper

}
