import { z } from "zod";

export const JwkSchema = z.object({
    use: z.string(),
    kty: z.string(),
    kid: z.string(),
    alg: z.string(),
    n: z.string(),
    e: z.string(),
});

export type JwkJwt = z.infer<typeof JwkSchema>;

export const JwkResponseSchema = z.object({
    keys: z.array(JwkSchema),
});

export type JwkResponse = z.infer<typeof JwkResponseSchema>;
