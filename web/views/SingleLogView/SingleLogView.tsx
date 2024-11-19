"use client";

import React, { useEffect, useState } from "react";
import Box from "@mui/material/Box";
import { loadSingleLogAsString } from "@/api/singlelog";
import LogViewer from "@/containers/LogViewer/LogViewer";
import LogMetadataTopBar from "@/views/SingleLogView/LogMetadataTopBar/LogMetadataTopBar";

// eslint-disable-next-line no-unused-vars
const loadSingleLog = async (id: string, setter: (logText: string) => void) => {
    const res = await loadSingleLogAsString(id);
    setter(res);
};

const SingleLogView: React.FC<{ logUUID: string }> = ({ logUUID }) => {
    const [logText, setLogText] = useState("Log text loading...");

    useEffect(() => {
        loadSingleLog(logUUID, setLogText);
    }, [logUUID]);

    return (
        <Box width="97%" height="calc(100vh - 200px)">
            <LogMetadataTopBar />
            <LogViewer data={logText} />
        </Box>
    );
};

export default SingleLogView;
