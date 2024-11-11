"use client";

import React from "react";
import moment from "moment";
import { useRouter } from "next/navigation";
import { Alert, AlertTitle, Box, Link, Typography } from "@mui/material";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import Spinner from "@/components/Spinner";
import { useGetAllLogs } from "@/api/hooks";
import { LogOverviewItem } from "@/types/api/logs";

const columns: GridColDef<LogOverviewItem>[] = [
    { field: "id", headerName: "ID", flex: 3, renderCell: row => <LogLink id={row.row.id} /> },
    { field: "logType", headerName: "Type", flex: 1 },
    { field: "platform", headerName: "Platform", flex: 1 },
    { field: "lineCount", headerName: "Lines", flex: 1 },
    {
        field: "tsFirstLine",
        headerName: "First timestamp",
        flex: 3,
        valueFormatter: value => moment(value).format("MMMM Do YYYY, h:mm:ss a"),
    },
    {
        field: "tsLastLine",
        headerName: "Last timestamp",
        flex: 3,
        valueFormatter: value => moment(value).format("MMMM Do YYYY, h:mm:ss a"),
    },
];

const Home = () => {
    const { data, error, isLoading } = useGetAllLogs();

    // eslint-disable-next-line react/no-unstable-nested-components
    const LogsOverview = () => {
        if (error) {
            return (
                <Alert severity="error">
                    <AlertTitle>Could not load logs</AlertTitle>
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
                    <AlertTitle>No logs present</AlertTitle>
                    No logs were found
                </Alert>
            );
        }

        return (
            <Box sx={{ width: "100%" }}>
                <DataGrid rows={data} columns={columns} getRowId={row => row.id} disableRowSelectionOnClick />
            </Box>
        );
    };

    return (
        <>
            <Typography variant="h4" mb={2}>
                Available logs
            </Typography>
            <LogsOverview />
        </>
    );
};

export default Home;

const LogLink: React.FC<{ id: string }> = ({ id }) => {
    const router = useRouter();

    return (
        // eslint-disable-next-line jsx-a11y/anchor-is-valid
        <Link sx={{ cursor: "pointer" }} onClick={() => router.push(`/logs/${id}`)}>
            {id}
        </Link>
    );
};
