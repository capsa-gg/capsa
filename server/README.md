# Capsa Server

## Requirements

* Go >= v1.23
* PostgreSQL >= 16
* Make
* Go dependencies installed via `make bootstrap`

### Supported operating systems

The tool will work on all operating systems and architectures supported by the Golang compiler.

Development for the server should be done on Linux(/WSL) or MacOS. This is due to the shell scripts used. Development on Windows is in theory possible, by manually performing the actions that the shell scripts automate.

## Configuration

To configure the application, use `cp config.example.yml config.yml` and set the required values.

To add a new database and set up a superuser for your local database, run:

```psql
CREATE DATABASE capsadb;
CREATE USER capsauser WITH PASSWORD 'capsapass';
ALTER USER capsauser WITH SUPERUSER;
```

Make sure to migrate up your database before running the application.

## Development and commands

For all commands, run `make`, which will print a help screen.

## Database code generation

The Go code to access the database is code-generated with SQLc. Database queries are written in SQL directly, and then it's converted to Go.

After modifying the SQL code in `./sql`, run `make sql` to generate the Go code. These changes must be committed to Git.

For details on how to use SQLc, see the [SQLc documentation](https://docs.sqlc.dev/en/latest/).

## Swagger documentation

API documentation is generated from the comments on the HTTP handlers. Therefore it is important that these comments are kept up-to-date and are added for every endpoint.

When working in HTTP handlers, make sure to generate the Swagger documentation before commiting your code.

To generate the Swagger documentation, run `make swagger`.

For details on how to add proper documentation to route handlers, see [swaggo/swag on GitHub](https://github.com/swaggo/swag?tab=readme-ov-file#api-operation).

## Creating new database migrations

Database migrations are written in SQL and defined in `./migrations`.

To add a new migration script, run `make migration` and follow the instructions on screen.

Database migrations should have both up and down scripts.

## Running tasks locally to know the CI will pass

First, run the formatting/linter with `make fmt`.

The commands the CI will use for testing:

```sh
make build-all
make ci
```

If these commands succeed locally, it is quite probable the CI will pass as well.
