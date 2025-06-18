"use client";

import { ThemeProvider } from "@mui/material";
import type React from "react";
import useUser from "@/context/UserContext";
import getTheme from "@/styles/theme";

const ThemeProviderWithMode: React.FC<ThemeProviderWithModeProps> = ({ children }) => {
    const {
        preferences: {
            preferences: { darkMode },
        },
    } = useUser();

    const theme = getTheme(darkMode);

    return <ThemeProvider theme={theme}>{children}</ThemeProvider>;
};

export default ThemeProviderWithMode;

interface ThemeProviderWithModeProps {
    children: React.ReactNode;
}
