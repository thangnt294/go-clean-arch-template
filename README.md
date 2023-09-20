# Clean architecture template for Go

## How to run
- `make install-tools` to install all tools
- `make init` to init infrastructures
- `make dev` to start the service

## Project structure
- `cmd` contains the main files to start the program
- `config` contains the config struct to use in the program
- `migrations` contains the migration files for database
- `internal` contains the internal files for use in the program

## Within the internal folder
- `internal/domain` contains all the domain files. Each domain file contains the necessary structs and interfaces for that domain.
- `internal/auth` is the domain-specific folder for the auth domain. Each domain-specific folder contains 3 additional layers: handler, usecase and repository.

## Domain layers
- **Handler** accepts requests, parse the body (or queries, params), use the **usecase** layer to handle business logics, and send back response to client.
- **Usecase** handles business logics of the auth domain, and call **repository** if it needs to work with the database.
- **Repository** works with the databases to create, read, update or delete data.

Reference repo: https://github.com/bxcodec/go-clean-arch
