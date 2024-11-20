import { LogMetadata } from "@/types/api/logs";
import { ApplicationError } from "@/types/api/error";
import React from "react";
import { SWRResponse } from "swr";

export interface SingleLogContextProviderProps {
    logUUID: string;
    children: React.ReactNode;
}

export interface SingleLogContextData {
    metadata: SWRResponse<LogMetadata, ApplicationError, null>;
}
