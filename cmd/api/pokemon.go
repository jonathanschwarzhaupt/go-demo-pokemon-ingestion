package main

import (
	"errors"
	"fmt"
	"github.com/jonathanschwarzhaupt/go-demo-pokemon-ingestion/internal/data"
	"github.com/jonathanschwarzhaupt/go-demo-pokemon-ingestion/internal/database"
	"github.com/jonathanschwarzhaupt/go-demo-pokemon-ingestion/internal/validator"
	"net/http"
	"strings"
)

// createPokemonHandler - handles POST /v1/pokemon, creating a new pokemon in DB
func (app *application) createPokemonHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Id             int64  `json:"id"`
		Name           string `json:"name"`
		Height         int64  `json:"height"`
		Weight         int64  `json:"weight"`
		BaseExperience int64  `json:"base_experience"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	pokemon := &data.Pokemon{
		Id:             input.Id,
		Name:           input.Name,
		Height:         input.Height,
		Weight:         input.Weight,
		BaseExperience: input.BaseExperience,
	}

	v := validator.New()

	if data.ValidatePokemon(v, pokemon); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	_, err = app.DB.CreatePokemon(r.Context(), database.CreatePokemonParams{
		ID:             pokemon.Id,
		Name:           pokemon.Name,
		Height:         pokemon.Height,
		Weight:         pokemon.Weight,
		BaseExperience: pokemon.BaseExperience,
	})
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			app.badRequestResponse(w, r, errors.New("pokemon already exists"))
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	// Set location header for the client to find resource
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/pokemon/%d", pokemon.Id))

	err = app.writeJSON(w, http.StatusCreated, envelope{"pokemon": pokemon}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// getPokemonHandler - handles GET /v1/pokemon returning all from DB
func (app *application) getPokemonHandler(w http.ResponseWriter, r *http.Request) {
	dbPokes, err := app.DB.GetPokemons(r.Context())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var resPokes []data.Pokemon
	for _, poke := range dbPokes {
		tmp := data.Pokemon{
			Id:             poke.ID,
			Name:           poke.Name,
			Height:         poke.Height,
			Weight:         poke.Weight,
			BaseExperience: poke.BaseExperience,
		}
		resPokes = append(resPokes, tmp)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"data": resPokes}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// getPokemonByIDHandler - handles GET /v1/pokemon/{id} returning it from DB
func (app *application) getPokemonByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	pokeDB, err := app.DB.GetPokemonByID(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}
	pokeRes := data.Pokemon{
		Id:             pokeDB.ID,
		Name:           pokeDB.Name,
		Height:         pokeDB.Height,
		Weight:         pokeDB.Weight,
		BaseExperience: pokeDB.BaseExperience,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"data": pokeRes}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// deletePokemonByIDHandler - deletes the pokemon from DB
func (app *application) deletePokemonByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	_, err = app.DB.DeletePokemon(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "pokemon deleted successfully"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
