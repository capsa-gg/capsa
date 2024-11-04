"use client";

import React from "react";
import AppBar from "@mui/material/AppBar";
import { Box, Toolbar, Typography } from "@mui/material";
import { sideNavWidthPx } from "@/layouts/MainLayout/SideNav/SideNav.data";
import { grey } from "@mui/material/colors";

export const topNavSidePaddingPx = 20;

const TopNav = () => (
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
                <Typography color={grey[900]}>
                    Capsa version 0.0.1 (hash) in (mode) mode
                </Typography>
            </Box>
            <Box>
                <Typography color={grey[900]}>Logged in as (user)</Typography>
            </Box>
        </Toolbar>
    </AppBar>
);

export default TopNav;
