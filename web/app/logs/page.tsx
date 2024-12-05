"use client";

import { getLogsOverview, useGetAllEnvironments } from "@/api/hooks";
import ColoredSeverities from "@/components/ColoredSeverities";
import Spinner from "@/components/Spinner";
import { useNotificationsContext } from "@/context/NotificationsContext/NotificationsContext";
import type { EnvironmentResponseItem, ListAllEnvironmentsResponse } from "@/types/api/environments";
import type { LogOverviewItem } from "@/types/api/logs";
import { formatDate } from "@/util/formatDate";
import {
    Alert,
    AlertTitle,
    Box,
    FormControl,
    InputLabel,
    Link,
    MenuItem,
    Select,
    type SelectChangeEvent,
    Stack,
    Typography,
} from "@mui/material";
import { DataGrid, type GridColDef } from "@mui/x-data-grid";
import { useRouter, useSearchParams } from "next/navigation";
import type React from "react";
import { Suspense } from "react";
import { useEffect, useState } from "react";

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

const LogsOverview = () => {
    const { addNotification } = useNotificationsContext();
    const searchParams = useSearchParams();
    const [hasNotified, setHasNotified] = useState(false);
    const { data: envs, isLoading: isLoadingEnv, error: errorEnv } = useGetAllEnvironments();
    const {
        data: logs,
        error: errorLogs,
        isLoading: isLoadingLogs,
    } = getLogsOverview(searchParams.get("env") === "all" ? "" : (searchParams.get("env") ?? ""));

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
                <EnvironmentSelector envs={envs} />
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

const LogLink: React.FC<{ id: string }> = ({ id }) => {
    const router = useRouter();

    return (
        <Link sx={{ cursor: "pointer" }} onClick={() => router.push(`/logs/${id}`)}>
            {id}
        </Link>
    );
};

const EnvironmentSelector: React.FC<{ envs?: ListAllEnvironmentsResponse }> = ({ envs }) => {
    const router = useRouter();
    const searchParams = useSearchParams();
    const currentEnv = searchParams.get("env") || "all";

    const handleChange = (event: SelectChangeEvent) => {
        const newEnv = event.target.value;
        const params = new URLSearchParams(searchParams);
        params.set("env", newEnv);
        router.push(`/logs?${params.toString()}`);
    };

    if (!envs) return null;

    return (
        <FormControl sx={{ minWidth: 200 }}>
            <InputLabel>Environment</InputLabel>
            <Select value={currentEnv} label="Environment" onChange={handleChange}>
                <MenuItem value="all">All</MenuItem>
                {envs
                    .sort((a, b) => (envToDisplayString(a) > envToDisplayString(b) ? 1 : -1))
                    .map(env => (
                        <MenuItem key={env.environmentKey} value={env.environmentKey}>
                            {envToDisplayString(env)}
                        </MenuItem>
                    ))}
            </Select>
        </FormControl>
    );
};

const envToDisplayString = (env: EnvironmentResponseItem): string => `${env.title} / ${env.environmentName}`;
