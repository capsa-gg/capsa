"use client";

import React, { useEffect, useState } from "react";
import Box from "@mui/material/Box";
import { loadSingleLogAsString } from "@/api/singlelog";
import LogViewer from "@/containers/LogViewer/LogViewer";

// eslint-disable-next-line no-unused-vars
const loadSingleLog = async (id: string, setter: (logText: string) => void) => {
    const res = await loadSingleLogAsString(id);
    setter(res);
};

const SingleLog: React.FC<SingleLogProps> = ({ params }) => {
    const paramsUse = React.use(params);
    const [logText, setLogText] = useState("Log text loading...");

    useEffect(() => {
        loadSingleLog(paramsUse.id, setLogText);
    }, [paramsUse.id]);

    return (
        <Box width="97%" height="calc(100vh - 200px)">
            <LogViewer data={logText} />
        </Box>
    );
};

export default SingleLog;

export interface SingleLogProps {
    params: Promise<{ id: string }>;
}
