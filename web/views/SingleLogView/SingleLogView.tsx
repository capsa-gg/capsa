"use client";

import LogViewer from "@/containers/LogViewer/LogViewer";
import { useSingleLogContext } from "@/context/SingleLogContext/SingleLogContext";
import LogLineFilterDrawer from "@/views/SingleLogView/LogLineFilterDrawer/LogLineFilterDrawer";
import LogMetadataTopBar from "@/views/SingleLogView/LogMetadataTopBar/LogMetadataTopBar";
import Box from "@mui/material/Box";
import type React from "react";

const SingleLogView: React.FC = () => {
    const { logProcessor } = useSingleLogContext();

    const logViewerData = () => {
        if (logProcessor.error) return `${logProcessor.error}`;
        return logProcessor.fullLog;
    };

    return (
        <Box width="97%" height="calc(100vh - 240px)">
            <LogLineFilterDrawer />
            <LogMetadataTopBar />
            <LogViewer data={logViewerData()} absoluteLineNumbers={logProcessor.absoluteLineNumbers} />
        </Box>
    );
};

export default SingleLogView;
