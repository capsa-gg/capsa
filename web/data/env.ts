"use server";

export const getEnv = async (): Promise<Env> => ({
    // biome-ignore lint/style/noNonNullAssertion: safe here
    serverUrl: process.env.SERVER_URL!,
});

export interface Env {
    serverUrl: string;
}
