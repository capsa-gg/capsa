import type { SideNavItemProps } from "@/layouts/MainLayout/SideNav/SideNav.types";
import { ListItem, ListItemButton, ListItemText } from "@mui/material";
import Link from "next/link";
import type React from "react";

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
