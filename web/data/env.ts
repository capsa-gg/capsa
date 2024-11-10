"use server";

export const getEnv = async (): Promise<Env> => ({
    serverUrl: process.env.SERVER_URL,
});

export interface Env {
    serverUrl: string;
}
