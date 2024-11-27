"use client";

import SingleLogView from "@/views/SingleLogView/SingleLogView";
import React from "react";

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
