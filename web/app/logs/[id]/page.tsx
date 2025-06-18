"use client";

import React from "react";
import SingleLogView from "@/views/SingleLogView/SingleLogView";

const SingleLog: React.FC<SingleLogProps> = ({ params }) => {
    const paramsUse = React.use(params);

    if (!paramsUse.id) {
        return null;
    }

    return <SingleLogView />;
};

export default SingleLog;

export interface SingleLogProps {
    params: Promise<{ id: string }>;
}
