package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/products", app.createProductHandler)
	router.HandlerFunc(http.MethodGet, "/v1/products/:id", app.showProductHandler)

	return router
}

