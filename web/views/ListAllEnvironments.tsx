"use client";

import { Alert, AlertTitle, Box } from "@mui/material";
import { DataGrid, type GridColDef } from "@mui/x-data-grid";
import type React from "react";
import { useGetAllEnvironments } from "@/api/hooks";
import Spinner from "@/components/Spinner";
import type { EnvironmentResponseItem } from "@/types/api/environments";

const columns: GridColDef<EnvironmentResponseItem>[] = [
    { field: "title", headerName: "Title", flex: 1 },
    { field: "environmentName", headerName: "Environment Name", flex: 1 },
    { field: "environmentKey", headerName: "Environment Key", flex: 3 },
];

const ListAllEnvironments: React.FC = () => {
    const { data, error, isLoading } = useGetAllEnvironments();

    if (error) {
        return (
            <Alert severity="error">
                <AlertTitle>Could not load environments</AlertTitle>
                {error?.error}
            </Alert>
        );
    }
    if (isLoading) {
        return <Spinner />;
    }

    if (!data) {
        return (
            <Alert severity="warning">
                <AlertTitle>No data present</AlertTitle>
                The environments data was not found.
            </Alert>
        );
    }

    return (
        <Box sx={{ width: "100%", maxWidth: 1000 }}>
            <DataGrid rows={data} columns={columns} getRowId={row => row.environmentKey} disableRowSelectionOnClick />
        </Box>
    );
};

export default ListAllEnvironments;
