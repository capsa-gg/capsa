import React from "react";
import { Alert, AlertTitle, Badge, Divider, IconButton, Typography } from "@mui/material";
import { Stack } from "@mui/system";
import Box from "@mui/material/Box";
import FilterAltIcon from "@mui/icons-material/FilterAlt";
import { LogMetadataList } from "@/views/SingleLogView/LogMetadataTopBar/LogMetadataTopBar.components";
import { useSingleLogContext } from "@/context/SingleLogContext/SingleLogContext";
import LogProcessor from "@/data/LogProcessor/LogProcessor";

const LogMetadataTopBar: React.FC = () => {
    const {
        drawerState: [drawerOpen, setDrawerOpen],
        metadata: { error, isLoading, data },
    } = useSingleLogContext();

    if (error) {
        return (
            <Alert severity="error">
                <AlertTitle>Error loading log metadata</AlertTitle>
                <Typography>
                    <b>{error.title}</b>
                </Typography>
                <Divider variant="fullWidth" />
                <Typography>{error.error}</Typography>
                <Typography>{error.additionalData}</Typography>
                <Typography>{error.rawError}</Typography>
            </Alert>
        );
    }

    if (isLoading || !data) {
        return <Typography>Loading metadata...</Typography>;
    }

    return (
        <Stack
            direction="row"
            spacing={2}
            sx={{
                justifyContent: "space-between",
                alignItems: "center",
            }}
        >
            <LogMetadataList logMetadata={data} />
            <Box>
                <Badge color="primary">
                    <IconButton onClick={() => setDrawerOpen(true)}>
                        <FilterAltIcon />
                    </IconButton>
                </Badge>
            </Box>
        </Stack>
    );
};

export default LogMetadataTopBar;
