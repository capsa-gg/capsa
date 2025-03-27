import type { OverridableComponent } from "@mui/material/OverridableComponent";
import type { SvgIconTypeMap } from "@mui/material/SvgIcon";

export interface SideNavItemProps {
    title: string;
    href: string;
    selected: boolean;
    disabled: boolean;
    Icon: OverridableComponent<SvgIconTypeMap> & { muiName: string };
}
