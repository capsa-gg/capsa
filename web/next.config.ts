import type { NextConfig } from "next";

const nextConfig: NextConfig = {
    output: "standalone",
    productionBrowserSourceMaps: true,
    reactStrictMode: true,
    // TODO: https://stackoverflow.com/questions/76124346/how-can-i-pass-an-env-variable-to-next-js-app-running-inside-a-docker-container
    // TODO: headers for security
};

export default nextConfig;
