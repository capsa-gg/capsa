"use client";

import React from "react";
import { Drawer, List } from "@mui/material";
import {
    sideNavItemsData,
    sideNavWidthPx,
} from "@/layouts/MainLayout/SideNav/SideNav.data";
import { SideNavItem } from "@/layouts/MainLayout/SideNav/SideNav.components";

// TODO: Add Capsa project title
// TODO: Add styling
const SideNav = () => {
    const SideNavListItems = () =>
        sideNavItemsData.map(s => (
            <SideNavItem key={s.title} title={s.title} href={s.href} />
        ));

    return (
        <Drawer variant="permanent" open={true}>
            <List sx={{ width: sideNavWidthPx }}>
                <SideNavListItems />
            </List>
        </Drawer>
    );
};

export default SideNav;
