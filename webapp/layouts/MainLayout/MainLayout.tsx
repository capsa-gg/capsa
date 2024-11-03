import React from "react";
import TopNav, { topNavSidePaddingPx } from "./TopNav";
import SideNav from "./SideNav/SideNav";
import { Box } from "@mui/material";
import { sideNavWidthPx } from "@/layouts/MainLayout/SideNav/SideNav.data";

const MainLayout: React.FC<MainLayoutProps> = ({ children }) => (
    <>
        <SideNav />
        <Box
            component="main"
            sx={{ marginLeft: `${sideNavWidthPx + topNavSidePaddingPx}px` }}
        >
            <TopNav />
            {children}
        </Box>
    </>
);

export default MainLayout;

export interface MainLayoutProps {
    children: React.ReactNode;
}
