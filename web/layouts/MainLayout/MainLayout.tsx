import React from "react";
import TopNav, { topNavSidePaddingPx } from "./TopNav";
import SideNav from "./SideNav/SideNav";
import { Box } from "@mui/material";

const MainLayout: React.FC<MainLayoutProps> = ({ children }) => (
    <>
        <SideNav />
        <Box component="main" sx={{ marginLeft: `240px` }}>
            <TopNav />
            <Box mt={4}>{children}</Box>
        </Box>
    </>
);

export default MainLayout;

export interface MainLayoutProps {
    children: React.ReactNode;
}
