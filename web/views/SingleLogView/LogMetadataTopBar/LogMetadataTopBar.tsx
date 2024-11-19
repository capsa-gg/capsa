import React from "react";
import { Alert, AlertTitle, Divider, Typography } from "@mui/material";
import { LogMetadataList } from "@/views/SingleLogView/LogMetadataTopBar/LogMetadataTopBar.components";
import { useSingleLogContext } from "@/context/SingleLogContext/SingleLogContext";

const LogMetadataTopBar: React.FC = () => {
    const {
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

    return <LogMetadataList logMetadata={data} />;
};

export default LogMetadataTopBar;
