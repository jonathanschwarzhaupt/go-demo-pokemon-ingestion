# Pokémon JSON API

## Motivation
I designed and created this little project as part of our team's hackathon session during the department strategy day in June 2024. Pokémon were specifically chosen as the data domain because they are accessible, fun, and contain nested data structures, providing an interesting yet manageable challenge for practicing ETL (Extract, Transform, Load) skills.

## Project Overview
This repository contains the source code for a JSON API developed using Go, specifically for this hackathon session. The API allows users to ingest Pokémon data fetched from external sources, designed to interact with ETL scripts created by teammates. Participants created Kubernetes cron jobs to periodically extract data from the Pokémon API and load it into this JSON API.

## API Features
- **Data Validation:** Ensures incoming Pokémon data matches the defined schema.
- **Health Check:** Endpoint provided to verify the API's operational status.
- **SQL Migrations:** Migrations against a remote database to manage data schema changes.

## Available Routes
- Refer to **`endpoints.md`** for a list of available resources

## Technologies & Learnings
- Written in Go, tried to leverage best practices for building robust and maintainable JSON APIs.
- Provided a practical introduction to JSON APIs, specifically data validation, API design, and database migrations.

## Database Setup (Turso)
To set up your remote database:
1. Sign up for a free account on [Turso](https://turso.tech/).
2. Create a new database and generate a connection token.
3. Copy the provided connection string and add it to a `.env` file in the project's root directory:
   ```
   DB_URL=libsql://...
   ```
Refer to `.env.example` for an example. 

## Deployment
1. **Run Database Migrations**
   ```bash
   ./scripts/migrate.sh
   ```

2. **Local Deployment**
   ```bash
   ./scripts/build-run-local.sh
   ```

The API will now be running locally and ready to accept Pokémon data from your ETL scripts on port 4000, by default.

## Further Resources
- [PokéAPI](https://pokeapi.co/) - Pokémon data source used by ETL scripts.
- [Boot.dev](https://boot.dev) - for an excellent introduction to backend development using Go.
- [Let's Go & Let's Go Further](https://www.alexedwards.net/) - two excellent books on building web applications in Go by Alex Edwards.

