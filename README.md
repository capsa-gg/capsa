# Capsa

[![License: AGPL v3](https://img.shields.io/badge/License-AGPL_v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)
[![Server](https://github.com/capsa-gg/capsa/actions/workflows/main-server.yml/badge.svg)](https://github.com/capsa-gg/capsa/actions/workflows/main-server.yml)
[![Web](https://github.com/capsa-gg/capsa/actions/workflows/main-web.yml/badge.svg)](https://github.com/capsa-gg/capsa/actions/workflows/main-webapp.yml)

## Getting started

_TODO: write_

## Development setup

For setting up the development environments, please check the folders `server` and `web` for their respective development setup instructions and requirements.

It is advisable to install auto formatters and .editorconfig support in your editor/IDE.

## Development guidelines

### Commit messages

Commit messages should follow the 'conventional commit' structure `<type>(<scope>): description`, for example:

```
feat(server): implement transactional emails
```

Merge commits are an exception for this, you don't have to modify these to fit the format.

The commit messages are also checked by the CI. To see the available types and scopes, please see `.commitlintrc.js` in the repo root, this file contains documentation about the correct usage. It is also recommended to install a code editor/IDE plugin for conventional commits.

The commit messages will be used to generate a changelog for each release.

### Code generation

This project tries to use code generation where possible.

Generated code means less time maintaining it and is often less error-prone than hand-rolled code.

For Typescript, code generation is not as prevalent as with Go. Therefore it relies more on helper functions that hide the complexity but offer a nice and simple API for users.

### Strict linting

No-one likes arguing about tabs/spaces or casing systems. Therefore, the linters for this project are set up quite strict.

The linting determines the code style, so the focus of reviews can be on the logic instead of style.

Make sure to run the linting locally before pushing commits to prevent the CI from failing.

### Secrets

Secrets should _never_ be pushed to the Git repository. These files should always be gitignored.

When writing tests using logs, make sure that these logs are allowed to be made public at some point, don't break NDAs or leak codenames and have no identifiable data in them. It is best to only use internal Companion Group project logs for this, with some search-replace logic to make the logs generic.

## Changing API types

Whenever changing the API req/res types on either the server or the client, please make sure these match.

## Release instructions

To create a new release, go to the [Semantic Release actions workflow](https://github.com/capsa-gg/capsa/actions/workflows/release.yml). On the top right, press `Run workflow`, confirm running the workflow and kick off semantic release. Make sure the web and server CIs on main have passed before kicking off a release. Do not change the branch, as this workflow will only work correctly on `main`.

The release commit will not run the default main branch CI workflows. Due to the `:latest` tag being pushed for the Docker images, they will be deployed.

Until we are ready to move to v0.1, all of the releases will be on patch versions.

After moving to v0.1, all breaking changes and features will increment the minor version, until the project is complete enough to be marked as v1.0.0, from where we will follow the regular semantic release patterns.
