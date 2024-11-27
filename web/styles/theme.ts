"use client";

import { amber, green } from "@mui/material/colors";
import { createTheme } from "@mui/material/styles";
import { Roboto } from "next/font/google";

const roboto = Roboto({
    weight: ["300", "400", "500", "700"],
    subsets: ["latin"],
    display: "swap",
});

const theme = createTheme({
    palette: {
        mode: "light", // TODO: Add dark mode support, https://mui.com/material-ui/customization/dark-mode/
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

export default theme;
