package data

import "github.com/jonathanschwarzhaupt/go-demo-pokemon-ingestion/internal/validator"

type Pokemon struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Height         int64  `json:"height"`
	Weight         int64  `json:"weight"`
	BaseExperience int64  `json:"base_experience"`
}

func ValidatePokemon(v *validator.Validator, pokemon *Pokemon) {
	v.Check(pokemon.Id != 0, "id", "must be provided")
	v.Check(pokemon.Id > 0, "id", "must be a positive integer")

	v.Check(pokemon.Name != "", "name", "must be provided")
	v.Check(len(pokemon.Name) > 2, "name", "must be longer then two characters")

	v.Check(pokemon.Height != 0, "height", "must be provided and a positive integer")
	v.Check(pokemon.Height > 0, "height", "must be a positive integer")

	v.Check(pokemon.Weight != 0, "weight", "must be provided")
	v.Check(pokemon.Weight > 0, "weight", "must be a positive integer")

	v.Check(pokemon.BaseExperience != 0, "base_experience", "must be provided")
	v.Check(pokemon.BaseExperience > 0, "base_experience", "must be a positive integer")
}
