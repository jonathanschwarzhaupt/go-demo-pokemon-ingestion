package main

import (
	"fmt"
	"github.com/jonathanschwarzhaupt/go-demo-pokemon-ingestion/internal/validator"
	"net/http"
)

type pokemon struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Order          int    `json:"order"`
	IsDefault      bool   `json:"is_default"`
}

// createPokemonHandler - handles POST /v1/pokemon, creating a new pokemon in DB
func (app *application) createPokemonHandler(w http.ResponseWriter, r *http.Request) {
	var input pokemon

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	v := validator.New()

	v.Check(input.Id != 0, "id", "must be provided")
	v.Check(input.Id > 0, "id", "must be a positive integer")
	v.Check(input.Name != "", "name", "must be provided")
	v.Check(len(input.Name) > 2, "name", "must be longer then two characters")
	v.Check(input.BaseExperience != 0, "base_experience", "must be provided")
	v.Check(input.BaseExperience > 0, "base_experience", "must be a positive integer")
	v.Check(input.Height != 0, "height", "must be provided and a positive integer")
	v.Check(input.Height > 0, "height", "must be a positive integer")
	v.Check(input.Weight != 0, "weight", "must be provided")
	v.Check(input.Weight > 0, "weight", "must be a positive integer")
	v.Check(input.Order != 0, "order", "must be provided")
	v.Check(input.Order > 0, "order", "must be a positive integer")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}
