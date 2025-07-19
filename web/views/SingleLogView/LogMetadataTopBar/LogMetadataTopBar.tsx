import DataObjectIcon from "@mui/icons-material/DataObject";
import FilterAltIcon from "@mui/icons-material/FilterAlt";
import MenuBookIcon from "@mui/icons-material/MenuBook";
import MergeIcon from "@mui/icons-material/Merge";
import { Alert, AlertTitle, Badge, Divider, IconButton, Tooltip, Typography } from "@mui/material";
import Box from "@mui/material/Box";
import { Stack } from "@mui/system";
import type React from "react";
import { useSingleLogContext } from "@/context/SingleLogContext/SingleLogContext";
import { LogMetadataList } from "@/views/SingleLogView/LogMetadataTopBar/LogMetadataTopBar.components";
import DownloadFullLogButton from "@/views/SingleLogView/DownloadFullLogButton/DownloadFullLogButton";

const LogMetadataTopBar: React.FC = () => {
    const {
        viewMode: [viewMode, setViewMode],
        filterDrawerState: [_filterDrawerOpen, setFilterDrawerOpen],
        mergeDrawerState: [_mergeDrawerOpen, setMergeDrawerOpen],
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
                    <Tooltip title="Download full log">
                        <DownloadFullLogButton />
                    </Tooltip>
                </Badge>
                <Badge color="primary">
                    <Tooltip title="Switch view mode between log and metadata">
                        {viewMode === "Log" ? (
                            <IconButton onClick={() => setViewMode("Metadata")}>
                                <DataObjectIcon />
                            </IconButton>
                        ) : (
                            <IconButton onClick={() => setViewMode("Log")}>
                                <MenuBookIcon />
                            </IconButton>
                        )}
                    </Tooltip>
                </Badge>
                <Badge color="primary">
                    <Tooltip title="Open merged logs drawer">
                        <IconButton onClick={() => setMergeDrawerOpen(true)}>
                            <MergeIcon />
                        </IconButton>
                    </Tooltip>
                </Badge>
                <Badge color="primary">
                    <Tooltip title="Open log line filters drawer">
                        <IconButton onClick={() => setFilterDrawerOpen(true)}>
                            <FilterAltIcon />
                        </IconButton>
                    </Tooltip>
                </Badge>
            </Box>
        </Stack>
    );
};

export default LogMetadataTopBar;
