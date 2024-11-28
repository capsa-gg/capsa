import type { SideNavItemProps } from "@/layouts/MainLayout/SideNav/SideNav.types";
import DarkModeIcon from "@mui/icons-material/DarkMode";
import DescriptionIcon from "@mui/icons-material/Description";
import GitHubIcon from "@mui/icons-material/GitHub";
import { IconButton, ListItem, ListItemButton, ListItemIcon, ListItemText, Tooltip } from "@mui/material";
import Box from "@mui/material/Box";
import Link from "next/link";
import type React from "react";

export const SideNavItem: React.FC<SideNavItemProps> = ({ title, href, selected, disabled, Icon }) => (
    <ListItem>
        <ListItemButton component={Link} href={href} disabled={disabled} selected={selected}>
            <ListItemIcon sx={{ minWidth: 40 }}>
                <Icon />
            </ListItemIcon>
            <ListItemText primary={title} />
        </ListItemButton>
    </ListItem>
);

export const BottomIcons = () => (
    <Box
        sx={{
            display: "flex",
            justifyContent: "center",
        }}
    >
        <Tooltip title={"Switch dark/light mode"} placement="top">
            <IconButton onClick={() => alert("todo")} color="primary" size="small" aria-label="Switch dark/light mode">
                <DarkModeIcon />
            </IconButton>
        </Tooltip>
        <Tooltip title={"Open Capsa on GitHub"} placement="top">
            <IconButton
                color="primary"
                size="small"
                href="https://github.com/capsa-gg"
                target="_blank"
                rel="noopener noreferrer"
                aria-label="GitHub"
            >
                <GitHubIcon fontSize="small" />
            </IconButton>
        </Tooltip>
        <Tooltip title={"Open Capsa documentation"} placement="top">
            <IconButton
                color="primary"
                size="small"
                href="https://capsa.gg"
                target="_blank"
                rel="noopener noreferrer"
                aria-label="Documentation"
            >
                <DescriptionIcon fontSize="small" />
            </IconButton>
        </Tooltip>
    </Box>
);
