import React from "react";
import type { Metadata } from "next";
import { AppRouterCacheProvider } from "@mui/material-nextjs/v15-appRouter";
import { ThemeProvider } from "@mui/system";
import CssBaseline from "@mui/material/CssBaseline";
import theme from "@/styles/theme";

import "./globals.css";

export const metadata: Metadata = {
    title: "Capsa Webapp Homepage",
    description: "Capsa WebApp access",
};

const RootLayout: React.FC<RootLayoutProps> = ({ children }) => (
    <html lang="en">
        <body>
            <AppRouterCacheProvider>
                <ThemeProvider theme={theme}>
                    <CssBaseline />
                    {children}
                </ThemeProvider>
            </AppRouterCacheProvider>
        </body>
    </html>
);

export default RootLayout;

type RootLayoutProps = Readonly<{
    children: React.ReactNode;
}>;
