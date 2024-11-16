"use client";

import React from "react";
import AppBar from "@mui/material/AppBar";
import { Box, Button, Toolbar, Typography } from "@mui/material";
import { grey } from "@mui/material/colors";
import { removeJwtCookie } from "@/data/jwt/cookiesClient";
import { removeJwtFromLocalStorage } from "@/data/jwt/localStorage";
import { useRouter } from "next/navigation";
import useUser from "@/context/UserContext";

export const topNavSidePaddingPx = 20;

const TopNav = () => {
    const router = useRouter();
    const {
        userInfo: { isLoggedIn, user },
    } = useUser();
    const signOut = () => {
        removeJwtCookie();
        removeJwtFromLocalStorage();
        router.refresh();
    };

    const userText = isLoggedIn ? `Welcome, ${user?.firstName}` : "Not logged in";

    return (
        <AppBar
            position="static"
            sx={{
                paddingRight: `40px`,
                background: "transparent",
                boxShadow: "none",
            }}
        >
            <Toolbar sx={{ justifyContent: "space-between" }} disableGutters>
                <Box>
                    <Typography color={grey[900]}>{userText}</Typography>
                </Box>
                <Box>
                    <Button variant="outlined" size="small" disabled={!isLoggedIn} onClick={signOut}>
                        Sign out
                    </Button>
                </Box>
            </Toolbar>
        </AppBar>
    );
};

export default TopNav;
