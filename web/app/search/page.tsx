"use client";

import { Alert, AlertTitle, Link, Tooltip, Typography } from "@mui/material";
import Box from "@mui/material/Box";
import { DataGrid, type GridColDef } from "@mui/x-data-grid";
import { useRouter } from "next/navigation";
import React from "react";
import { useSearch } from "@/api/hooks";
import Spinner from "@/components/Spinner";
import type { SearchResultItem } from "@/types/api/search";

const RenderMatch: React.FC<{ row: SearchResultItem }> = ({ row: { type, match, details } }) => {
    const router = useRouter();

    if (type === "Logs") {
        const hasChunks = Number.parseInt(details, 10) > 0;

        if (hasChunks) {
            return (
                <Link sx={{ cursor: "pointer" }} onClick={() => router.push(`/logs/${match}`)}>
                    {match}
                </Link>
            );
        }
        return (
            // @ts-expect-error
            <Tooltip title="This log does not have any chunks">{match}</Tooltip>
        );
    }

    return <>{match}</>;
};

const columns: GridColDef<SearchResultItem>[] = [
    { field: "type", headerName: "Type", flex: 1 },
    { field: "match", headerName: "Identifier", flex: 3, renderCell: row => <RenderMatch row={row.row} /> },
    { field: "description", headerName: "Description", flex: 3 },
];

const Search: React.FC<SearchProps> = ({ searchParams }) => {
    const searchParamsUse = React.use(searchParams);
    const { query } = searchParamsUse;
    const { data, isLoading, error } = useSearch((query && (query as string)) as string);

    if (!query) {
        return <Typography variant="h6">Enter a search query in the top bar</Typography>;
    }
    if (error) {
        return (
            <Alert severity="error">
                <AlertTitle>Could not load search results</AlertTitle>
                {error?.error}
            </Alert>
        );
    }
    if (isLoading) {
        return <Spinner />;
    }

    return (
        <Box>
            <Typography variant="h6">Results for {query}</Typography>
            {data?.hasMore && <Alert severity="info">There are more results, please narrow down your search.</Alert>}
            <DataGrid rows={data?.results} columns={columns} getRowId={row => row.match} disableRowSelectionOnClick />
        </Box>
    );
};

export default Search;

export interface SearchProps {
    searchParams: Promise<{ [key: string]: string | string[] | undefined }>;
}
