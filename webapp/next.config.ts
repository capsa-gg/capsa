import type { NextConfig } from "next";

const nextConfig: NextConfig = {
    output: "standalone",
    productionBrowserSourceMaps: true,
    reactStrictMode: true,
    swcMinify: true,
};

export default nextConfig;
