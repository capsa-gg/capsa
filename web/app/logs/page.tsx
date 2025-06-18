"use client";

import { Alert, AlertTitle, Box, Stack, Typography } from "@mui/material";
import { DataGrid, type GridColDef } from "@mui/x-data-grid";
import { type ReadonlyURLSearchParams, useSearchParams } from "next/navigation";
import { Suspense, useEffect, useState } from "react";
import { getLogsOverview, useGetAllEnvironments } from "@/api/hooks";
import ColoredSeverities from "@/components/ColoredSeverities";
import Spinner from "@/components/Spinner";
import { useNotificationsContext } from "@/context/NotificationsContext/NotificationsContext";
import type { LogOverviewItem } from "@/types/api/logs";
import { formatDate } from "@/util/formatDate";
import { EnvironmentSelector, LogLink, LogTypeSelector, PlatformInput } from "./components";

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
        field: "linkedLogCount",
        headerName: "Links",
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

const searchParamsToString = (searchParams: ReadonlyURLSearchParams): string => {
    const paramsOut = new URLSearchParams();

    searchParams.forEach((val, key) => {
        if (val !== null && val !== "" && val !== "all") {
            paramsOut.set(key, val);
        }
    });

    return paramsOut.toString();
};

const LogsOverview = () => {
    const { addNotification } = useNotificationsContext();
    const searchParams = useSearchParams();
    const [hasNotified, setHasNotified] = useState(false);
    const { data: envs, isLoading: isLoadingEnv, error: errorEnv } = useGetAllEnvironments();
    const {
        data: logs,
        error: errorLogs,
        isLoading: isLoadingLogs,
    } = getLogsOverview(searchParamsToString(searchParams));

    useEffect(() => {
        if (!isLoadingLogs && logs?.hasMore && !hasNotified) {
            addNotification({
                severity: "info",
                title: "Not all logs are shown",
                message: "There are more logs than shown here due to the configured filters.",
            });
            setHasNotified(true);
        }
    }, [isLoadingLogs, logs, logs?.hasMore, addNotification, hasNotified]);

    const LogsOverview = () => {
        if (errorEnv) {
            return (
                <Alert severity="error">
                    <AlertTitle>Could not load environments</AlertTitle>
                    {errorEnv?.error}
                </Alert>
            );
        }
        if (errorLogs) {
            return (
                <Alert severity="error">
                    <AlertTitle>Could not load logs</AlertTitle>
                    {errorLogs?.error}
                </Alert>
            );
        }
        if (isLoadingLogs || isLoadingEnv) {
            return <Spinner />;
        }
        if (!logs) {
            return (
                <Alert severity="warning">
                    <AlertTitle>No logs present</AlertTitle>
                    No logs were found
                </Alert>
            );
        }

        return (
            <Box sx={{ width: "100%", maxWidth: "1400px" }}>
                <DataGrid rows={logs.logs} columns={columns} getRowId={row => row.id} disableRowSelectionOnClick />
            </Box>
        );
    };

    return (
        <Box mr="40px">
            <Stack direction="row" justifyContent="space-between" mb={6}>
                <Typography variant="h4">Available logs</Typography>
                <Stack direction="row" gap={2}>
                    <PlatformInput />
                    <LogTypeSelector />
                    <EnvironmentSelector envs={envs} />
                </Stack>
            </Stack>
            <LogsOverview />
        </Box>
    );
};

const LogsOverviewPage = () => (
    <Suspense>
        <LogsOverview />
    </Suspense>
);

export default LogsOverviewPage;
