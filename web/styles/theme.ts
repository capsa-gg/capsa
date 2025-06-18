"use client";

import { amber, green } from "@mui/material/colors";
import { createTheme, type Theme } from "@mui/material/styles";
import { Roboto } from "next/font/google";

const roboto = Roboto({
    weight: ["300", "400", "500", "700"],
    subsets: ["latin"],
    display: "swap",
});

const themeLight = createTheme({
    palette: {
        mode: "light",
        primary: {
            main: green[300],
        },
        secondary: {
            main: amber[200],
        },
        background: {
            default: "#eee",
        },
    },
    typography: {
        fontFamily: roboto.style.fontFamily,
    },
});

const themeDark = createTheme({
    palette: {
        mode: "dark",
        primary: {
            main: green[300],
        },
        secondary: {
            main: amber[200],
        },
        background: {
            default: "#222",
        },
    },
    typography: {
        fontFamily: roboto.style.fontFamily,
    },
    components: {
        MuiDrawer: { styleOverrides: { paper: { backgroundColor: "#282626" } } },
        MuiPaper: { styleOverrides: { root: { backgroundColor: "#282626" } } },
    },
});

const getTheme = (darkMode: boolean): Theme => (darkMode ? themeDark : themeLight);

export default getTheme;
