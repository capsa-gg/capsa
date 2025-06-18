import HomeIcon from "@mui/icons-material/Home";
import MenuBookIcon from "@mui/icons-material/MenuBook";
import PersonIcon from "@mui/icons-material/Person";
import SportsEsportsIcon from "@mui/icons-material/SportsEsports";
import type { SideNavItemProps } from "@/layouts/MainLayout/SideNav/SideNav.types";

export const sideNavWidthPx = 200;

export const sideNavItemsDataUsers: Omit<SideNavItemProps, "selected" | "disabled">[] = [
    {
        title: "Home",
        href: "/",
        Icon: HomeIcon,
    },
    {
        title: "Logs",
        href: "/logs",
        Icon: MenuBookIcon,
    },
];

export const sideNavItemsDataAdmin: Omit<SideNavItemProps, "selected" | "disabled">[] = [
    {
        title: "Users",
        href: "/users",
        Icon: PersonIcon,
    },
    {
        title: "Environments",
        href: "/environments",
        Icon: SportsEsportsIcon,
    },
];
