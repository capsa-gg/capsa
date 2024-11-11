"use client";

import React, { useEffect, useState } from "react";
import { loadSingleLogAsString } from "@/api/singlelog";

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

    return <code style={{ whiteSpace: "pre" }}>{logText}</code>;
};

export default SingleLog;

export interface SingleLogProps {
    params: Promise<{ id: string }>;
}
