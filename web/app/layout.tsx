import CssBaseline from "@mui/material/CssBaseline";
import { AppRouterCacheProvider } from "@mui/material-nextjs/v15-appRouter";
import type { Metadata } from "next";
import type React from "react";
import { SWRProvider } from "@/app/swr-provider";
import ErrorSnackbar from "@/containers/NotificationSnackbar/NotificationSnackbar";
import { NotificationsContextProvider } from "@/context/NotificationsContext/NotificationsContext";
import MainLayout from "@/layouts/MainLayout/MainLayout";
import ScreenTooSmall from "@/layouts/ScreenTooSmall/ScreenTooSmall";
import "./globals.css";
import ThemeProviderWithMode from "@/app/ThemeProviderWithMode";
import { AppContextProvider } from "@/context/AppContext/AppContext";
import { UserProvider } from "@/context/UserContext";

export const metadata: Metadata = {
    title: "Capsa Webapp",
    description: "Capsa WebApp",
};

const RootLayout: React.FC<RootLayoutProps> = ({ children }) => (
    <html lang="en">
        <body>
            <AppRouterCacheProvider>
                <NotificationsContextProvider>
                    <AppContextProvider>
                        <UserProvider>
                            <ThemeProviderWithMode>
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
                            </ThemeProviderWithMode>
                        </UserProvider>
                    </AppContextProvider>
                </NotificationsContextProvider>
            </AppRouterCacheProvider>
        </body>
    </html>
);

export default RootLayout;

type RootLayoutProps = Readonly<{
    children: React.ReactNode;
}>;
