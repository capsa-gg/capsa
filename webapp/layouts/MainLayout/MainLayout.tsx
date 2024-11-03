import React from "react";
import TopNav from "./TopNav";
import SideNav from "./SideNav";

const MainLayout: React.FC<MainLayoutProps> = ({ children }) => (
    <>
        <TopNav />
        <SideNav />
        <main>{children}</main>
    </>
);

export default MainLayout;

export interface MainLayoutProps {
    children: React.ReactNode;
}
