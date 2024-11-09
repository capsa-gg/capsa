import { z } from "zod";

export const LoginRequestSchema = z.object({ email: z.string(), password: z.string() });

export type LoginRequest = z.infer<typeof LoginRequestSchema>;

export const UserJwtDataSchema = z.object({
    token: z.string(),
    firstName: z.string(),
    lastName: z.string(),
    email: z.string(),
    userUUID: z.string(),
    tokenExpiry: z.string(),
});

export type UserJwtData = z.infer<typeof UserJwtDataSchema>;

export const UserPasswordResetConfirmationSchema = z.object({
    resetToken: z.string(),
    password: z.string(),
});

export type UserPasswordResetConfirmation = z.infer<typeof UserPasswordResetConfirmationSchema>;
