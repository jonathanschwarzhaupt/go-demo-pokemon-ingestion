package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/pokemons", app.createPokemonHandler)
	router.HandlerFunc(http.MethodGet, "/v1/pokemons", app.getPokemonHandler)
	router.HandlerFunc(http.MethodGet, "/v1/pokemons/:id", app.getPokemonByIDHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/pokemons/:id", app.deletePokemonByIDHandler)

	return router
}
