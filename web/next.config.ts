import { withSentryConfig } from "@sentry/nextjs";
import type { NextConfig } from "next";

const nextConfig: NextConfig = {
    output: "standalone",
    productionBrowserSourceMaps: true,
    reactStrictMode: true,
    cacheMaxMemorySize: 0, // disable default in-memory caching
    logging: {
        fetches: {
            fullUrl: true,
            hmrRefreshes: true,
        },
    },
    experimental: {
        serverComponentsHmrCache: false, // defaults to true
    },

    async headers() {
        return [
            {
                source: "/(.*)",
                headers: [
                    {
                        key: "Strict-Transport-Security",
                        value: `max-age=${60 * 60 * 24 * 365 /* 1y */}; includeSubDomains; preload`,
                    },
                    {
                        key: "Permissions-Policy",
                        value: "camera=(), geolocation=(), microphone=()",
                    },
                    {
                        key: "Referrer-Policy",
                        value: "origin-when-cross-origin",
                    },
                    {
                        key: "X-Frame-Options",
                        value: "DENY",
                    },
                    {
                        key: "X-Content-Type-Options",
                        value: "nosniff",
                    },
                ],
            },
        ];
    },
};

const getNextConfig = (): NextConfig => {
    const sentryAuthToken = process.env.SENTRY_AUTH_TOKEN;

    if (!sentryAuthToken) {
        return nextConfig;
    }

    return withSentryConfig(nextConfig, {
        // For source map upload in CI
        org: process.env.SENTRY_ORG,
        project: process.env.SENTRY_PROJECT,
        authToken: sentryAuthToken,

        widenClientFileUpload: true,
        reactComponentAnnotation: {
            enabled: true,
        },

        tunnelRoute: "/monitoring",

        disableLogger: true,
    });
};

export default getNextConfig();
