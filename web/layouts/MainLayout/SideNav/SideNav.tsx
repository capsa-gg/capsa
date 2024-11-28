"use client";

import { useAppContext } from "@/context/AppContext/AppContext";
import useUser from "@/context/UserContext";
import { BottomIcons, SideNavItem } from "@/layouts/MainLayout/SideNav/SideNav.components";
import {
    sideNavItemsDataAdmin,
    sideNavItemsDataUsers,
    sideNavWidthPx,
} from "@/layouts/MainLayout/SideNav/SideNav.data";
import version from "@/version";
import { Box, Divider, Drawer, List, Tooltip, Typography } from "@mui/material";
import { grey } from "@mui/material/colors";
import Image from "next/image";
import { usePathname } from "next/navigation";

// TODO: Add styling
const SideNav = () => {
    const appContext = useAppContext();
    const {
        userInfo: { isLoggedIn, user },
    } = useUser();
    const pathname = usePathname();

    const SideNavListItemsUsers = () =>
        sideNavItemsDataUsers.map(s => (
            <SideNavItem
                key={s.title}
                title={s.title}
                href={s.href}
                disabled={!isLoggedIn}
                Icon={s.Icon}
                selected={s.href === "/" ? pathname === s.href : pathname.includes(s.href)}
            />
        ));

    const SideNavListItemsAdmin = () =>
        user?.role === "Admin" &&
        sideNavItemsDataAdmin.map(s => (
            <SideNavItem
                key={s.title}
                title={s.title}
                href={s.href}
                disabled={!isLoggedIn || user?.role !== "Admin"}
                Icon={s.Icon}
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
                    <SideNavListItemsUsers />
                </List>
                {/* TODO: Add about/help page link */}
                <List sx={{ width: sideNavWidthPx }}>
                    <SideNavListItemsAdmin />
                    <Divider sx={{ mt: 1, mb: 3 }} />
                    <Tooltip
                        title={`User ID: ${user?.userUUID ?? "Loading..."}`}
                        placement="top"
                        suppressHydrationWarning
                    >
                        <Typography
                            component="p"
                            color={grey[400]}
                            variant="caption"
                            align="center"
                            marginBottom={1}
                            sx={{ width: "100%" }}
                        >
                            Welcome {user?.firstName} ({user?.role})
                        </Typography>
                    </Tooltip>
                    <Tooltip
                        title={`Capsa API Server: ${appContext.env?.serverUrl ?? "Loading endpoint..."}`}
                        placement="top"
                        suppressHydrationWarning
                    >
                        <Typography
                            component="p"
                            color={grey[400]}
                            variant="caption"
                            align="center"
                            marginBottom={1}
                            sx={{ width: "100%" }}
                        >
                            Capsa v{version}
                        </Typography>
                    </Tooltip>
                    <BottomIcons />
                </List>
            </Box>
        </Drawer>
    );
};

export default SideNav;
