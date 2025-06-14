# cross-module-migration-example

## Point

Example of how to set up database migration process in modular go project.
We will assume that our codebase contains database migrations in different modules,
those migrations follow all rules of building a loosely coupled system - therefore
they do not know of each other.

I will use [pressly/goose](github.com/pressly/goose), because it provides API as a library,
if your tool does not provide itself as a library use a scripting language.

# How to run the example

```sh
docker-compose up -d # start the PostgreSQL container

# Setting up env variables, you might use dotenv or task for that purpose.
POSTGRES_HOST=127.0.0.1 \
POSTGRES_PORT=5432 \
POSTGRES_USER=user \
POSTGRES_PASSWORD=password \
POSTGRES_DB=db \
go run ./cmd/migration up # run the migration 
```

To learn how it's done, check out `cmd/migration/main.go` - it's pretty simple!
