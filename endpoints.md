# Pokémon JSON API

## Endpoints
All endpoints are unprotected, i.e. no authentication required. Available endpoints are:

| Method 	| Endpoint        	| Purpose                    	|
|--------	|-----------------	|----------------------------	|
| GET    	| /v1/healthcheck 	| Display system information 	|
| GET    	| /v1/pokemons    	| Get list of pokemons       	|
| POST   	| /v1/pokemons    	| Add a new pokemon          	|
| GET    	| /v1/pokemons/id 	| Get a specific pokemon     	|
| DELETE 	| /v1/pokemons/id 	| Delete a specific pokemon  	|

### GET /v1/healthcheck
Method: `GET`

Returns: Health stats

Example:
```
{
    "healthcheck": {
        "environment": "development",
        "status": "available",
        "version": "1.0.0"
    }
}
```

### GET /v1/pokemons
Method: `GET`

Returns: List of `Pokemon`

Example:
```
{
    "data": [
        {
            "id": 1,
            "name": "bulbasaur",
            "height": 7,
            "weight": 69,
            "base_experience": 64
        }
    ]
}
```

### POST /v1/pokemons
Method: `POST`

Returns: `Pokemon`

Request body: `Pokemon`

Example:
```
{
    "id": 1,
    "name": "bulbasaur",
    "height": 7,
    "weight": 69,
    "base_experience": 64
}
```

### GET /v1/pokemons/id
Method: `GET`

Returns: `Pokemon`

Example: `/v1/pokemons/1`
```
{
    "id": 1,
    "name": "bulbasaur",
    "height": 7,
    "weight": 69,
    "base_experience": 64
}
```

### DELETE /v1/pokemons
Method: `DELETE`

Returns: message

Example: `/v1/pokemons/1`
```
{
    "message": "pokeon deleted successfully"
}
```

## Objects
### Pokemon
| Field           	| Type   	|
|-----------------	|--------	|
| id              	| int64  	|
| name            	| string 	|
| height          	| int64  	|
| weight          	| int64  	|
| base_experience 	| int64  	|