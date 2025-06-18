"use client";

import { Typography } from "@mui/material";
import Box from "@mui/material/Box";
import type React from "react";
import useUser from "@/context/UserContext";

const AdminOnly: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    const {
        userInfo: { user },
    } = useUser();

    if (!user) {
        return null;
    }

    if (user.role !== "Admin") {
        return (
            <Box mt={8}>
                <Typography variant="h4">This page is for admins only.</Typography>
            </Box>
        );
    }

    return <>{children}</>;
};

export default AdminOnly;
