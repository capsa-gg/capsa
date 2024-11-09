import { SideNavItemProps } from "@/layouts/MainLayout/SideNav/SideNav.types";

export const sideNavWidthPx = 200;

export const sideNavItemsData: Omit<SideNavItemProps, "selected" | "disabled">[] = [
    {
        title: "Home",
        href: "/",
    },
    {
        title: "Example",
        href: "/example",
    },
];
