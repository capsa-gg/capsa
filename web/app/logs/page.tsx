"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { Alert, AlertTitle, Box, Link, Typography } from "@mui/material";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import Spinner from "@/components/Spinner";
import { useGetAllLogs } from "@/api/hooks";
import { LogOverviewItem } from "@/types/api/logs";
import ColoredSeverities from "@/components/ColoredSeverities";
import { formatDateString } from "@/util/formatDateString";

const columns: GridColDef<LogOverviewItem>[] = [
    { field: "id", headerName: "ID", maxWidth: 300, flex: 4, renderCell: row => <LogLink id={row.row.id} /> },
    { field: "logType", headerName: "Type", maxWidth: 100, flex: 1 },
    { field: "platform", headerName: "Platform", maxWidth: 100, flex: 1 },
    { field: "lineCount", headerName: "Lines", maxWidth: 100, flex: 1 },
    {
        field: "severitiesCount",
        headerName: "Severities",
        flex: 3,
        minWidth: 130,
        maxWidth: 180,
        sortable: false,
        renderCell: row => <ColoredSeverities severities={row.row.severitiesCounts} />,
    },
    {
        field: "categoriesCounts",
        headerName: "Categories",
        flex: 1,
        maxWidth: 100,
        valueFormatter: value => Object.keys(value).length,
    },
    {
        field: "tsFirstLine",
        headerName: "First timestamp",
        flex: 3,
        maxWidth: 200,
        valueFormatter: value => formatDateString(value),
    },
    {
        field: "tsLastLine",
        headerName: "Last timestamp",
        flex: 3,
        maxWidth: 200,
        valueFormatter: value => formatDateString(value),
    },
];

const LogsOverviewPage = () => {
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
            <Box sx={{ width: "100%", maxWidth: "1300px" }}>
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

export default LogsOverviewPage;

const LogLink: React.FC<{ id: string }> = ({ id }) => {
    const router = useRouter();

    return (
        // eslint-disable-next-line jsx-a11y/anchor-is-valid
        <Link sx={{ cursor: "pointer" }} onClick={() => router.push(`/logs/${id}`)}>
            {id}
        </Link>
    );
};
