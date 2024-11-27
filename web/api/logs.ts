"use client";

import { type ApiResponse, call } from "@/api/apibase";
import {
    type LogMetadata,
    LogMetadataSchema,
    type LogOverviewResponse,
    LogOverviewResponseSchema,
} from "@/types/api/logs";

// biome-ignore lint/complexity/noStaticOnlyClass: desired here
export class Logs {
    public static async ListAllLogs(): Promise<ApiResponse<LogOverviewResponse>> {
        return await call("GET", "/user/logs", LogOverviewResponseSchema, true);
    }

    public static async GetSingleLogMetadata(logUuid: string): Promise<ApiResponse<LogMetadata>> {
        return await call("GET", `/user/logs/${logUuid}/metadata`, LogMetadataSchema, true);
    }
}
