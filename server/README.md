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

The server requires a public/private key set to run. To do this, run `make jwkset`. We use public/private keys for JWTs instead of the commonly used secret value, to make it possible for the webapp server to validate tokens without needing access to the signing secret.

To add a new database and set up a superuser for your local database, run:

```psql
CREATE DATABASE capsadb;
CREATE USER capsauser WITH PASSWORD 'capsapass';
ALTER USER capsauser WITH SUPERUSER;
```

Make sure to migrate up your database before running the application.

## Development and commands

For all commands, run `make`, which will print a help screen.

For development on Mac and Linux, you can run `make dev` to run the application with hot-reloading. Do keep in mind that this hot reloading does _not_ generate the Swagger or SQL files, you need to manually run the commands for this. Hot-reloading under WSL works with caveats.

## Database code generation

The Go code to access the database is code-generated with SQLc. Database queries are written in SQL directly, and then it's converted to Go.

After modifying the SQL code in `./sql`, run `make sql` to generate the Go code. These changes must be committed to Git. For new SQL queries, please follow the existing casing and style patterns.

For details on how to use SQLc, see the [SQLc documentation](https://docs.sqlc.dev/en/latest/).

The database methods should be following the pattern `[Verb][Entity][OptionalArgument][OptionalBy]`, for example `GetUserById`, `ListAvailableLogs` or `GetChunksForLogByUuid`. For verbs, please use the correctly matching verbs:

* `Add`: add one or multiple entries to the database
* `List`: list all entries without filtering
* `Get`: list single or multiple entries with filtering
* `Initialize`: initialize a flow, often touching multiple data, or with more logic than an update
* `Update`: update an entry
* `Delete`: delete one or multiple entries

## Swagger documentation

API documentation is generated from the comments on the HTTP handlers. Therefore, it is important that these comments are kept up-to-date and are added for every endpoint.

When working in HTTP handlers, make sure to generate the Swagger documentation before commiting your code. Most likely, new req/res bodies will be defined in the `bodies` package, though sometimes it makes more sense to define them as `entities`.

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

## Architecture

This architecture is loosely inspired by Clean Architecture and Domain Driven Design, but adopted to be a bit more flexible.

### Import rules

Please make sure to adhere to the current import rules. The rules here are not overly strict. It is fine to import a package if you need to refer to types for example, but not to call functions from them. If you need to access instances, they should be passed as arguments, but not created.

Specific guidelines need to be established, but for now it's best to follow existing patterns. For example: `domain` only imports types but does not create new instances.

### Domains

Domains contain the business logic of the application. Think of authentication handling and log parsing. Domain instances are not created, but the public functions are called by passing a services interactor to use different systems.

The functions in domains should be seen as the heart of the application logic, where no input/output validation should take place and where logic should be as clean and pure as possible.

### Configuration and services interactor

Two very commonly used data types are `*entities.Config` and `*interactor.Services`. These are used as a form of dependency injection for the called code to call instances created in the calling code. This services interactor also includes an instance of the generated database module, which can be used for database access, so domain code does not need to import the database directly or manage connections.

The configuration is loaded on application start and validated. With the configuration, different services are created, the instances of which are used to create the interactor.

For small utilities, it is best to use `*entities.Config` if you just need a configuration value to access outside services or configuration values. For situations like the http server, you want to use `*interactor.Services` so you can use instances of all services created during the application start.

For easier iteration, the `*interactor.Services` contains direct references to structs and not to interfaces. At some point, when there will be different implementations of the same service (for example different file storage systems), these can be moved to the `interactor` package. For now, that is not worth doing as it would be quite cumbersome to have to change signatures in different places for very little benefit.

In modules where the configuration or services are often used, the pointer to the instance can be stored as a struct member.

### HTTP handlers and middleware

The HTTP routes are defined in `server/routes.go`. These, in turn, call the with Swagger comments annotated handler methods in `server/handlers`.

The handlers are responsible for validating the request data (which can be done by using the `extractBodyJSON` method and validation tags in the request types) and returning the correct HTTP codes (which can be done with the `sendErrorResponse` method).

All request and response types should have struct tags for JSON (de)marshaling.

Authenticated routes in the `routes.go` file should be placed in router groups and use a middleware handler to validate the request before processing. In case the middleware does not succeed, it _must_ call `c.Abort()` to prevent the handler from being executed. The handlers themselves should not check token validation.

### Returning the correct errors

Application logic should return `entities.DomainError` where possible, with the correct values set (this is ensured by using the `NewDomainError` functions). The `sendErrorResponse` method of the `handlers` package will ensure the correct error code will be sent.

Database errors are also handled in the same method and should be wrapped, rather than returning `DomainError`.

### Tests

Simple tests can be written with the standard library or with `testify/require`, according to preference.

Larger tests can best be written using `testify`.

## Miscellaneous development notes

### JSON field naming convention

For client routes, use snake_case, for user routes, use camelCase. This way, we can use camelCase in the web app.

### Error wrapping

When handling error blocks, try to always wrap errors with extra context.

For example, when receiving an error coming from the database when fetching a user, use:
```go
return nil, fmt.Errorf("error fetching user from database: %w", err)
```

The `%w` here is very important, this will add context, while keeping the error type the same (for `errors.Is` checks). Using `%s` in the `fmt.Errorf` call does _not_ work.
