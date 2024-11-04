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

For all commands, run `make`.
