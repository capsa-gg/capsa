This is a [Next.js](https://nextjs.org) project bootstrapped with [`create-next-app`](https://nextjs.org/docs/app/api-reference/cli/create-next-app).

[![Checked with Biome](https://img.shields.io/badge/Checked_with-Biome-60a5fa?style=flat&logo=biome)](https://biomejs.dev)

## Requirements

- NodeJS 22

## Getting Started

Run `cp .env.example .env.local` and set the correct variables.

After running `npm install`, you can run the development server with

```bash
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

The development server will have hot reloading set

## API types and endpoints

Typescript does not validate Typescript types during runtime. To mitigate this, we use type validation using Zod when possible.

For converting JSON types to the types in the webapp, you can use [JSON to Zod](https://transform.tools/json-to-zod) to generate Zod definitions, then use `z.infer<typeof YourZodSchema>` to generate a Typescript type.

For adding new endpoints, many helper functions are available. With a Zod schema generated, it takes only a few minutes to add the call to the correct endpoint with runtime type checking, making a hook available to React components. You can look at existing hooks for this and how to add new ones with the proper arguments.

## Authentication

This application does not handle the authentication and authorization, but is merely a way to access the server APIs. There is middleware to check the authentication of a user with their cookie, validating the JWT with a JWK and adding redirects if needed. This is the easiest way to keep the auth logic out of the application logic.

## Deployment

Deployment can best be done with Docker. There is a Dockerfile which will produce a standalone instance of this application. After setting the correct environment variables, it can be run in nearly any Docker-supporting environment.

## Development guidelines

For development on the web app, please only make use of the app router. Keep logic server-side when possible, except for API requests, try to set `"use client|server";` where it can only be run on one of them, to make it explicit for readers of the code. The application should not become a SPA, so URLs and URL parameters should update accordingly to make specific views sharable. With the app router, it is possible to separate the page contents from the layout, please try to do so. Using the app router layouts is great for similar pages as well.

State management should be done at the local level when possible. For shared states, please use `Context`s with a custom `useContextName` hook to access these, for complex logic, a custom hook instance provided by a context is a good idea. The providers should be kept at as low level as possible, try to avoid putting them all in the root layout.

For input should be validated and show meaningful errors. An example of this can be found in the auth pages.

## Environment variable access

The `NEXT_PUBLIC_` environment variables should be avoided, as get inlined during the build process. That means that they should only be used for variables that don't change with different deployments, making them just an inferior non-typesafe alternative to a constant in the code.

Use the `data/env.ts` file instead for server environment variable access in the browser. Be cautious with this however, everything in here will be available for clients!
