"use client";

import Box from "@mui/material/Box";
import type React from "react";
import LogViewer from "@/containers/LogViewer/LogViewer";
import { useSingleLogContext } from "@/context/SingleLogContext/SingleLogContext";
import LogLineFilterDrawer from "@/views/SingleLogView/LogLineFilterDrawer/LogLineFilterDrawer";
import LogMetadataPage from "@/views/SingleLogView/LogMetadataPage/LogMetadataPage";
import LogMetadataTopBar from "@/views/SingleLogView/LogMetadataTopBar/LogMetadataTopBar";
import MergeLogsDrawer from "@/views/SingleLogView/MergeLogsDrawer/MergeLogsDrawer";

const SingleLogView: React.FC = () => {
    const {
        viewMode: [viewMode],
        logProcessor,
    } = useSingleLogContext();

    const logViewerData = () => {
        if (logProcessor.error) return `${logProcessor.error}`;
        return logProcessor.fullLog;
    };

    const ViewBody = () => {
        if (viewMode === "Log") {
            return <LogViewer data={logViewerData()} absoluteLineNumbers={logProcessor.absoluteLineNumbers} />;
        }
        return <LogMetadataPage />;
    };

    return (
        <Box width="97%" height="calc(100vh - 240px)">
            <LogLineFilterDrawer />
            <MergeLogsDrawer />
            <LogMetadataTopBar />
            <ViewBody />
        </Box>
    );
};

export default SingleLogView;
