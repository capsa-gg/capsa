"use client";

import React, { useEffect } from "react";
import Box from "@mui/material/Box";
import LogViewer from "@/containers/LogViewer/LogViewer";
import LogMetadataTopBar from "@/views/SingleLogView/LogMetadataTopBar/LogMetadataTopBar";
import useLogProcessor from "@/hooks/useLogProcessor/useLogProcessor";

const SingleLogView: React.FC<{ logUUID: string }> = ({ logUUID }) => {
    const { fullLog, error, startFetchingLog, stopFetchingLog } = useLogProcessor();

    useEffect(() => {
        startFetchingLog(logUUID);
        return () => {
            stopFetchingLog();
        };
    }, [logUUID, startFetchingLog, stopFetchingLog]);

    const logViewerData = () => {
        if (error) return `${error}`;
        return fullLog;
    };

    return (
        <Box width="97%" height="calc(100vh - 200px)">
            <LogMetadataTopBar />
            <LogViewer data={logViewerData()} />
        </Box>
    );
};

export default SingleLogView;
