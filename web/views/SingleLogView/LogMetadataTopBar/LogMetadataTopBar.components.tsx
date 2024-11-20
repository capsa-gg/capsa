import React from "react";
import Box from "@mui/material/Box";
import { Stack, styled } from "@mui/system";
import { Divider, Paper } from "@mui/material";
import { LogMetadata } from "@/types/api/logs";
import { formatDateString } from "@/util/formatDateString";

const Item = styled(Paper)(({ theme }) => ({
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
            <Item>
                Log type: <b>{logMetadata.logData.logType}</b>
            </Item>
            <Item>
                Platform: <b>{logMetadata.logData.platform}</b>
            </Item>
            <Item>
                Line: <b>{logMetadata.logData.lineCount}</b>
            </Item>
            <Item>
                Chunks: <b>{logMetadata.logData.chunkCount}</b>
            </Item>
            <Item>
                Categories: <b>{Object.keys(logMetadata.logData.categoriesCounts).length}</b>
            </Item>
            <Item>
                Start: <b>{formatDateString(logMetadata.logData.tsFirstLine)}</b>
            </Item>
            <Item>
                End: <b>{formatDateString(logMetadata.logData.tsLastLine)}</b>
            </Item>
        </Stack>
    </Box>
);
