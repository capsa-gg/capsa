import { Box } from "@mui/material";
import type React from "react";
import SideNav from "./SideNav/SideNav";
import TopNav from "./TopNav";

const MainLayout: React.FC<MainLayoutProps> = ({ children }) => (
    <>
        <SideNav />
        <Box component="main" sx={{ marginLeft: "240px" }}>
            <TopNav />
            <Box mt={4}>{children}</Box>
        </Box>
    </>
);

export default MainLayout;

export interface MainLayoutProps {
    children: React.ReactNode;
}
