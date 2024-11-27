"use client";

import { useGetAllLogs } from "@/api/hooks";
import ColoredSeverities from "@/components/ColoredSeverities";
import Spinner from "@/components/Spinner";
import type { LogOverviewItem } from "@/types/api/logs";
import { formatDate } from "@/util/formatDate";
import { Alert, AlertTitle, Box, Link, Typography } from "@mui/material";
import { DataGrid, type GridColDef } from "@mui/x-data-grid";
import { useRouter } from "next/navigation";
import type React from "react";

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
        field: "chunkCount",
        headerName: "Chunks",
        flex: 1,
        maxWidth: 100,
    },
    {
        field: "tsFirstLine",
        headerName: "First timestamp",
        flex: 3,
        maxWidth: 200,
        valueFormatter: value => formatDate(value),
    },
    {
        field: "tsLastLine",
        headerName: "Last timestamp",
        flex: 3,
        maxWidth: 200,
        valueFormatter: value => formatDate(value),
    },
];

const LogsOverviewPage = () => {
    const { data, error, isLoading } = useGetAllLogs();

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
            <Box sx={{ width: "100%", maxWidth: "1400px" }}>
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
        <Link sx={{ cursor: "pointer" }} onClick={() => router.push(`/logs/${id}`)}>
            {id}
        </Link>
    );
};
