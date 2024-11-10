"use client";

import React from "react";
import { usePathname } from "next/navigation";
import { Box, Drawer, List, Tooltip, Typography } from "@mui/material";
import { grey } from "@mui/material/colors";
import { sideNavItemsData, sideNavWidthPx } from "@/layouts/MainLayout/SideNav/SideNav.data";
import { SideNavItem } from "@/layouts/MainLayout/SideNav/SideNav.components";
import version from "@/version";
import useUserInfo from "@/hooks/useUserInfo/useUserInfo";
import { useAppContext } from "@/context/AppContext/AppContext";

// TODO: Add Capsa project title
// TODO: Add styling
const SideNav = () => {
    const appContext = useAppContext();
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
                <Tooltip
                    title={`Server API: ${appContext.env?.serverUrl ?? "Loading endpoint..."}`}
                    placement="top"
                    suppressHydrationWarning
                >
                    <Typography color={grey[400]} variant="caption" align="center" marginBottom={2}>
                        Capsa v{version}
                    </Typography>
                </Tooltip>
            </Box>
        </Drawer>
    );
};

export default SideNav;
