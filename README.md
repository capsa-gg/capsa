# Capsa

## Getting started

_TODO: write_

## Development setup

For setting up the development environments, please check the folders `server` and `webapp` for their respective development setup instructions and requirements.

It is advisable to install auto formatters and .editorconfig support in your editor/IDE.

## Development guidelines

### Commit messages

Commit messages should follow the structure `<type>(<scope>): description`, for example: `feat(server): implement transactional emails`. Merge commits are an exception for this, you don't have to modify these.

The commit messages are also checked by the CI. To see the available types and scopes, please see `.commitlintrc.js` in the repo root.

The commit messages will be used to generate a changelog for each release.

### Code generation

### Strict linting

No-one likes arguing about tabs/spaces or casing systems. Therefore, the linters for this project are set up quite strict.

The linting determines the code style, so the focus of reviews can be on the logic instead of style.

Make sure to run the linting locally before pushing commits to prevent the CI from failing.

### Secrets

Secrets should _never_ be pushed to the Git repository. These files should always be gitignored.

When writing tests using logs, make sure that these logs are allowed to be made public at some point, don't break NDAs or leak codenames and have no identifiable data in them. It is best to only use internal Companion Group project logs for this, with some search-replace logic to make the logs generic.

## Release instructions

_TODO: write_
