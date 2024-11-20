import React, { useState } from "react";
import Box from "@mui/material/Box";
import { Stack, styled } from "@mui/system";
import { Badge, Divider, IconButton, Paper } from "@mui/material";
import { LogMetadata } from "@/types/api/logs";
import { formatDateString } from "@/util/formatDateString";

const MetadataItem = styled(Paper)(({ theme }) => ({
    // @ts-ignore
    ...theme.typography.body,
    textAlign: "center",
    color: theme.palette.text.secondary,
    height: 40,
    lineHeight: "40px",
    padding: "0 16px",
}));

export const LogMetadataList: React.FC<{ logMetadata: LogMetadata }> = ({ logMetadata }) => (
    <Box
        sx={{
            borderRadius: 2,
            bgcolor: "background.default",
            gap: 2,
            paddingBottom: 2,
        }}
    >
        <Stack direction="row" spacing={2}>
            <MetadataItem>
                Log type: <b>{logMetadata.logData.logType}</b>
            </MetadataItem>
            <MetadataItem>
                Platform: <b>{logMetadata.logData.platform}</b>
            </MetadataItem>
            <MetadataItem>
                Line: <b>{logMetadata.logData.lineCount}</b>
            </MetadataItem>
            <MetadataItem>
                Chunks: <b>{logMetadata.logData.chunkCount}</b>
            </MetadataItem>
            <MetadataItem>
                Categories: <b>{Object.keys(logMetadata.logData.categoriesCounts).length}</b>
            </MetadataItem>
            <MetadataItem>
                Start: <b>{formatDateString(logMetadata.logData.tsFirstLine)}</b>
            </MetadataItem>
            <MetadataItem>
                End: <b>{formatDateString(logMetadata.logData.tsLastLine)}</b>
            </MetadataItem>
        </Stack>
    </Box>
);
