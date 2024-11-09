import React from "react";
import { ListItem, ListItemButton, ListItemText } from "@mui/material";
import Link from "next/link";
import { SideNavItemProps } from "@/layouts/MainLayout/SideNav/SideNav.types";

// TODO: Include showing active item
// TODO: Styling
export const SideNavItem: React.FC<SideNavItemProps> = ({ title, href, selected, disabled }) => (
    <ListItem>
        <ListItemButton component={Link} href={href} disabled={disabled} selected={selected}>
            {/*TODO: Add ListItemIcon here*/}
            <ListItemText primary={title} />
        </ListItemButton>
    </ListItem>
);
