"use server";

export const getEnv = async (): Promise<Env> => ({
    // biome-ignore lint/style/noNonNullAssertion: safe here
    serverUrl: process.env.SERVER_URL!,

    sentryDsn: process.env.SENTRY_DSN,
});

export interface Env {
    serverUrl: string;
    sentryDsn?: string;
}
