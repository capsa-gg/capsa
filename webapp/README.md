This is a [Next.js](https://nextjs.org) project bootstrapped with [`create-next-app`](https://nextjs.org/docs/app/api-reference/cli/create-next-app).

## Requirements

* NodeJS 22

## Getting Started

Run `cp .env.example .env.local` and set the correct variables.

After running `npm install`, you can run the development server with

```bash
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

The development server will have hot reloading set

## API types

Typescript does not validate Typescript types during runtime. To mitigate this, we use type validation using Zod when possible.

For converting JSON types to the types in the webapp, you can use [JSON to Zod](https://transform.tools/json-to-zod) to generate Zod definitions, then use `z.infer<typeof YourZodSchema>` to generate a Typescript type.

## Authentication

## Deployment

## Development guidelines
