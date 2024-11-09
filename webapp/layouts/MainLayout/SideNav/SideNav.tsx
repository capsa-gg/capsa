"use client";

import React from "react";
import { Box, Drawer, List, Tooltip, Typography } from "@mui/material";
import { sideNavItemsData, sideNavWidthPx } from "@/layouts/MainLayout/SideNav/SideNav.data";
import { SideNavItem } from "@/layouts/MainLayout/SideNav/SideNav.components";
import { grey } from "@mui/material/colors";
import { usePathname } from "next/navigation";
import useUserInfo from "@/hooks/useUserInfo/useUserInfo";

// TODO: Add Capsa project title
// TODO: Add styling
const SideNav = () => {
    const { isLoggedIn } = useUserInfo();
    const pathname = usePathname();

    const SideNavListItems = () =>
        sideNavItemsData.map(s => (
            <SideNavItem
                key={s.title}
                title={s.title}
                href={s.href}
                disabled={!isLoggedIn}
                selected={s.href === "/" ? pathname === s.href : pathname.includes(s.href)}
            />
        ));

    return (
        <Drawer variant="permanent" open={true}>
            <Box
                sx={{
                    display: "flex",
                    flexDirection: "column",
                    justifyContent: "space-between",
                    height: "100vh",
                    width: sideNavWidthPx,
                }}
            >
                <List sx={{ width: sideNavWidthPx }}>
                    <SideNavListItems />
                </List>
                <Tooltip title={`Server API: ${process.env.NEXT_PUBLIC_SERVER_URL}`} placement="top">
                    <Typography color={grey[400]} variant="caption" align="center" marginBottom={2}>
                        Capsa version 0.0.1
                    </Typography>
                </Tooltip>
            </Box>
        </Drawer>
    );
};

export default SideNav;
