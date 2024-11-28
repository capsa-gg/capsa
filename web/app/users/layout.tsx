import { AdminUsersContextProvider } from "@/app/users/useAdminUsers";
import AdminOnly from "@/util/AdminOnly";
import { Stack } from "@mui/material";
import type React from "react";

const UsersLayout: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    return (
        <AdminOnly>
            <AdminUsersContextProvider>
                <Stack>{children}</Stack>
            </AdminUsersContextProvider>
        </AdminOnly>
    );
};

export default UsersLayout;
