"use client";

import useAdminUsers from "@/app/users/useAdminUsers";
import Spinner from "@/components/Spinner";
import type { SingleUserResponse, UserRole } from "@/types/api/users";
import { formatDate } from "@/util/formatDate";
import ArchiveIcon from "@mui/icons-material/Archive";
import DangerousIcon from "@mui/icons-material/Dangerous";
import GroupAddIcon from "@mui/icons-material/GroupAdd";
import GroupRemoveIcon from "@mui/icons-material/GroupRemove";
import UnarchiveIcon from "@mui/icons-material/Unarchive";
import VerifiedIcon from "@mui/icons-material/Verified";
import { Alert, AlertTitle, Box, ButtonGroup, IconButton, Tooltip } from "@mui/material";
import { DataGrid, type GridColDef } from "@mui/x-data-grid";
import type React from "react";

const columns: GridColDef<SingleUserResponse>[] = [
    { field: "userUuid", headerName: "User UUID", width: 300 },
    { field: "email", headerName: "Email", flex: 3 },
    { field: "firstName", headerName: "First name", flex: 1 },
    { field: "lastName", headerName: "Last name", flex: 1 },
    {
        field: "hasPasswordSet",
        headerName: "Password",
        description: "Shows whether the user has set their password after account creation",
        width: 90,
        renderCell: ({ row: { hasPasswordSet } }) => (
            <Box display="flex" alignItems="center" justifyContent="center" width="100%" height="100%">
                {hasPasswordSet ? <VerifiedIcon color="success" /> : <DangerousIcon color="warning" />}
            </Box>
        ),
    },
    { field: "role", headerName: "Role", width: 80 },
    {
        field: "createdOn",
        headerName: "Created on",
        width: 180,
        renderCell: ({ row }) => <>{formatDate(row.createdAt)}</>,
    },
    {
        field: "deactivatedTs",
        headerName: "Deactivated on",
        width: 180,
        renderCell: ({ row }) => <>{row.deactivatedTs ? formatDate(row.deactivatedTs) : "N/A"}</>,
    },
    {
        field: "actions",
        headerName: "Actions",
        width: 80,
        renderCell: ({ row }) => <UserRowActions row={row} />,
    },
];

const UserRowActions: React.FC<{ row: SingleUserResponse }> = ({ row }) => {
    const { allUsersIsLoading, isUpdating, updateUser, deactivateUser, reactivateUser } = useAdminUsers();
    const disableButtons = allUsersIsLoading || isUpdating;

    const ActivationAction = () => {
        if (row.deactivatedTs === null) {
            return (
                <Tooltip title="Deactivate user">
                    <IconButton size="small" onClick={() => deactivateUser(row.userUuid)} disabled={disableButtons}>
                        <ArchiveIcon fontSize="inherit" color={disableButtons ? "disabled" : "warning"} />
                    </IconButton>
                </Tooltip>
            );
        }
        return (
            <Tooltip title="Reactivate user">
                <IconButton size="small" onClick={() => reactivateUser(row.userUuid)} disabled={disableButtons}>
                    <UnarchiveIcon fontSize="inherit" color={disableButtons ? "disabled" : "success"} />
                </IconButton>
            </Tooltip>
        );
    };

    const RoleAction = () => {
        const withUpdatedRole = (role: UserRole) => ({
            firstName: row.firstName,
            lastName: row.lastName,
            role: role,
        });
        if (row.role === "Admin") {
            return (
                <Tooltip title="Change role to 'User'">
                    <IconButton
                        size="small"
                        onClick={() => updateUser(row.userUuid, withUpdatedRole("User"))}
                        disabled={disableButtons}
                    >
                        <GroupRemoveIcon fontSize="inherit" color={disableButtons ? "disabled" : "warning"} />
                    </IconButton>
                </Tooltip>
            );
        }
        return (
            <Tooltip title="Change role to 'Admin'">
                <IconButton
                    size="small"
                    onClick={() => updateUser(row.userUuid, withUpdatedRole("Admin"))}
                    disabled={disableButtons}
                >
                    <GroupAddIcon fontSize="inherit" color={disableButtons ? "disabled" : "success"} />
                </IconButton>
            </Tooltip>
        );
    };

    return (
        <ButtonGroup variant="outlined">
            <ActivationAction />
            <RoleAction />
        </ButtonGroup>
    );
};

const UserManagement = () => {
    const { allUsers, allUsersIsLoading } = useAdminUsers();

    if (allUsersIsLoading) {
        return <Spinner />;
    }

    if (!allUsers) {
        return (
            <Alert severity="warning">
                <AlertTitle>No data present</AlertTitle>
                The environments data was not found.
            </Alert>
        );
    }

    return (
        <Box sx={{ width: "100%" }}>
            <DataGrid rows={allUsers} columns={columns} getRowId={row => row.userUuid} disableRowSelectionOnClick />
        </Box>
    );
};

export default UserManagement;
