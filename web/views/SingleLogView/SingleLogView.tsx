"use client";

import React from "react";
import Box from "@mui/material/Box";
import LogViewer from "@/containers/LogViewer/LogViewer";
import LogMetadataTopBar from "@/views/SingleLogView/LogMetadataTopBar/LogMetadataTopBar";
import { useSingleLogContext } from "@/context/SingleLogContext/SingleLogContext";
import LogLineFilterDrawer from "@/views/SingleLogView/LogLineFilterDrawer/LogLineFilterDrawer";

const SingleLogView: React.FC = () => {
    const { logProcessor } = useSingleLogContext();

    const logViewerData = () => {
        if (logProcessor.error) return `${logProcessor.error}`;
        return logProcessor.fullLog;
    };

    return (
        <Box width="97%" height="calc(100vh - 200px)">
            <LogLineFilterDrawer />
            <LogMetadataTopBar />
            <LogViewer data={logViewerData()} />
        </Box>
    );
};

export default SingleLogView;
