"use client";

import { AddEnvironmentForm } from "@/app/environments/AddEnvironmentForm";
import AddTitleForm from "@/app/environments/AddTitleForm";
import AdminOnly from "@/util/AdminOnly";
import ListAllEnvironments from "@/views/ListAllEnvironments";
import { Divider, Stack, Typography } from "@mui/material";
import Box from "@mui/material/Box";
import { useState } from "react";

const EnvironmentsManagement = () => {
    const [addedTitle, setAddedTitle] = useState("");

    return (
        <AdminOnly>
            <Typography variant="h4" mb={6} mt={4}>
                Title and environment management
            </Typography>
            <Stack gap={4}>
                <Box>
                    <Typography variant="h6" mb={2}>
                        Add title
                    </Typography>
                    <AddTitleForm onTitleAdded={setAddedTitle} />
                </Box>
                <Divider />
                <Box>
                    <Typography variant="h6" mb={2}>
                        Add environment
                    </Typography>
                    <AddEnvironmentForm addedTitle={addedTitle} />
                </Box>
                <Divider />
                <Box>
                    <Typography variant="h6" mb={2}>
                        Available environments
                    </Typography>
                    <ListAllEnvironments />
                </Box>
            </Stack>
        </AdminOnly>
    );
};

export default EnvironmentsManagement;
