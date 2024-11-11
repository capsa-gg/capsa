import type { NextConfig } from "next";

const nextConfig: NextConfig = {
    output: "standalone",
    productionBrowserSourceMaps: true,
    reactStrictMode: true,
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
                        value: "camera=(); battery=(); geolocation=(); microphone=()",
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

export default nextConfig;
