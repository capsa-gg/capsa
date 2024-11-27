"use client";

import { useAppContext } from "@/context/AppContext/AppContext";
import useUser from "@/context/UserContext";
import { SideNavItem } from "@/layouts/MainLayout/SideNav/SideNav.components";
import { sideNavItemsData, sideNavWidthPx } from "@/layouts/MainLayout/SideNav/SideNav.data";
import version from "@/version";
import { Box, Drawer, List, Tooltip, Typography } from "@mui/material";
import { grey } from "@mui/material/colors";
import Image from "next/image";
import { usePathname } from "next/navigation";

// TODO: Add styling
const SideNav = () => {
    const appContext = useAppContext();
    const {
        userInfo: { isLoggedIn },
    } = useUser();
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
                    <Box
                        sx={{
                            display: "flex",
                            justifyContent: "center",
                            marginTop: "6px",
                            marginBottom: "24px",
                        }}
                    >
                        <Image
                            src={require("../../../public/logo-without-by.png")}
                            alt="Capsa Logo"
                            width={100}
                            height={40}
                        />
                    </Box>
                    <SideNavListItems />
                </List>
                {/* TODO: Add about/help page link */}
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
