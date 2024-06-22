-- name: CreatePokemon :one
INSERT INTO pokemon (id, name, height, weight, base_experience)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: GetPokemonByID :one
SELECT * FROM pokemon WHERE id = ?;

-- name: GetPokemons :many
SELECT * FROM pokemon;

-- name: DeletePokemon :one
DELETE FROM pokemon WHERE id = ?
RETURNING *;
