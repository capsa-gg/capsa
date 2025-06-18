"use client";

import AddIcon from "@mui/icons-material/Add";
import { Alert, AlertTitle, Button, Stack, Typography } from "@mui/material";
import Box from "@mui/material/Box";
import { useState } from "react";
import NewUserDialog from "@/app/users/NewUserDialog";
import UserManagement from "@/app/users/UserManagement";
import useAdminUsers from "@/app/users/useAdminUsers";

const Errors = () => {
    const { errors } = useAdminUsers();
    return (
        <>
            {errors
                .filter(e => !!e)
                .map(e => (
                    <Alert key={e?.name} severity="error">
                        <AlertTitle>{e.name}</AlertTitle>
                        {e.message}
                    </Alert>
                ))}
        </>
    );
};

const Users = () => {
    const [newUserDialogOpen, setNewUserDialogOpen] = useState(false);

    return (
        <Stack mr="40px">
            <Stack mb={4} sx={{ display: "flex", flexDirection: "row", justifyContent: "space-between" }}>
                <Typography variant="h4">User management</Typography>
                <Box>
                    <Button variant="contained" startIcon={<AddIcon />} onClick={() => setNewUserDialogOpen(true)}>
                        New User
                    </Button>
                </Box>
            </Stack>
            <Errors />
            <UserManagement />
            <NewUserDialog open={newUserDialogOpen} onClose={() => setNewUserDialogOpen(false)} />
        </Stack>
    );
};

export default Users;
