// This file configures the initialization of Sentry on the client.
// The config you add here will be used whenever a users loads a page in their browser.
// https://docs.sentry.io/platforms/javascript/guides/nextjs/

import * as Sentry from "@sentry/nextjs";
import { getEnv } from "@/data/env";

const initSentry = async () => {
    const env = await getEnv();

    if (!env.sentryDsn) return;

    Sentry.init({
        dsn: env.sentryDsn,

        integrations: [Sentry.replayIntegration(), Sentry.feedbackIntegration({ colorScheme: "system" })],

        tracesSampleRate: 1,
        replaysSessionSampleRate: 0.1,
        replaysOnErrorSampleRate: 1.0,
        debug: false,
    });
};

initSentry();
