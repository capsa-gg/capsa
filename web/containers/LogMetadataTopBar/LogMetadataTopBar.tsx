import React from "react";
import { useGetSingleLogMetadata } from "@/api/hooks";
import { Alert, AlertTitle, Divider, Typography } from "@mui/material";
import { LogMetadataList } from "@/containers/LogMetadataTopBar/LogMetadataTopBar.components";

const LogMetadataTopBar: React.FC<{ logUUID: string }> = ({ logUUID }) => {
    const { data, isLoading, error } = useGetSingleLogMetadata(logUUID);

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

    return <LogMetadataList logMetadata={data} />;
};

export default LogMetadataTopBar;
