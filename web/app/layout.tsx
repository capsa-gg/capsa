import React from "react";
import type { Metadata } from "next";
import { AppRouterCacheProvider } from "@mui/material-nextjs/v15-appRouter";
import { ThemeProvider } from "@mui/material";
import CssBaseline from "@mui/material/CssBaseline";
import theme from "@/styles/theme";
import ScreenTooSmall from "@/layouts/ScreenTooSmall/ScreenTooSmall";
import MainLayout from "@/layouts/MainLayout/MainLayout";
import { SWRProvider } from "@/app/swr-provider";
import { ErrorProvider } from "@/context/ErrorContext";
import ErrorSnackbar from "@/containers/ErrorSnackbar/ErrorSnackbar";
import "./globals.css";
import { AppContextProvider } from "@/context/AppContext/AppContext";
import { UserProvider } from "@/context/UserContext";

export const metadata: Metadata = {
    title: "Capsa Webapp Homepage",
    description: "Capsa WebApp access",
};

// TODO: Error boundary

const RootLayout: React.FC<RootLayoutProps> = ({ children }) => (
    <html lang="en">
        <body>
            <AppRouterCacheProvider>
                <ThemeProvider theme={theme}>
                    <AppContextProvider>
                        <UserProvider>
                            <ErrorProvider>
                                <SWRProvider>
                                    <CssBaseline />
                                    <div className="screen-too-small">
                                        <ScreenTooSmall />
                                    </div>
                                    <div className="main-contents">
                                        <ErrorSnackbar />
                                        <MainLayout>{children}</MainLayout>
                                    </div>
                                </SWRProvider>
                            </ErrorProvider>
                        </UserProvider>
                    </AppContextProvider>
                </ThemeProvider>
            </AppRouterCacheProvider>
        </body>
    </html>
);

export default RootLayout;

type RootLayoutProps = Readonly<{
    children: React.ReactNode;
}>;
