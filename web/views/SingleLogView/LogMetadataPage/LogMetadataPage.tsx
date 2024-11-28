"use client";
import Spinner from "@/components/Spinner";
import { useSingleLogContext } from "@/context/SingleLogContext/SingleLogContext";
import type { LinkedLog } from "@/types/api/logs";
import { formatDate } from "@/util/formatDate";
import { Link, Stack, Typography } from "@mui/material";
import Box from "@mui/material/Box";
import { DataGrid, type GridColDef } from "@mui/x-data-grid";
import { useRouter } from "next/navigation";
import type React from "react";

const LogMetadataPage: React.FC = () => {
    const { metadata } = useSingleLogContext();
    const router = useRouter();

    if (!metadata.data) {
        return <Spinner />;
    }

    const columnsLinkedLogs: GridColDef<LinkedLog>[] = [
        {
            field: "linkedLog",
            headerName: "Linked Log",
            flex: 1,
            width: 300,
            renderCell: ({ row }) => (
                <Link sx={{ cursor: "pointer" }} onClick={() => router.push(`/logs/${row.linkedLog}`)}>
                    {row.linkedLog}
                </Link>
            ),
        },
        { field: "description", headerName: "Description", width: 500 },
    ];

    const LinkedLogs = () => {
        if (!metadata.data?.linkedLogs?.length) {
            return <Typography variant="caption">No linked logs</Typography>;
        }
        return (
            <DataGrid
                rows={metadata.data.linkedLogs}
                columns={columnsLinkedLogs}
                getRowId={row => row.linkedLog}
                disableRowSelectionOnClick
                sx={{ maxWidth: 800 }}
            />
        );
    };

    const AdditionalMetadataEntries = () => {
        if (!metadata.data?.additionalMetadata?.length) {
            return <Typography variant="caption">No additional metadata</Typography>;
        }

        const additionalMetadata = metadata.data?.additionalMetadata.flatMap(item =>
            Object.entries(item.metadata).map(([key, value]) => ({
                savedOn: item.savedOn,
                key,
                value,
            })),
        );

        return (
            <DataGrid
                rows={additionalMetadata}
                columns={[
                    { field: "key", headerName: "Key", width: 300 },
                    { field: "value", headerName: "Description", flex: 1 },
                    {
                        field: "savedOn",
                        headerName: "Saved on",
                        width: 230,
                        valueFormatter: value => formatDate(value),
                    },
                ]}
                getRowId={row => `${row.savedOn}-${row.key}`}
                disableRowSelectionOnClick
                sx={{ maxWidth: 1000 }}
            />
        );
    };

    return (
        <Stack gap={4} sx={{ mt: 4 }}>
            <Box>
                <Typography variant="h6" sx={{ mb: 2 }}>
                    Linked logs
                </Typography>
                <LinkedLogs />
            </Box>
            <Box>
                <Typography variant="h6" sx={{ mb: 2 }}>
                    Additional metadata
                </Typography>
                <AdditionalMetadataEntries />
            </Box>
        </Stack>
    );
};

export default LogMetadataPage;
