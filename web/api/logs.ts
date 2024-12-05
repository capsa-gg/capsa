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
    public static async ListLogs(env: string): Promise<ApiResponse<LogOverviewResponse>> {
        const endpoint = `/user/logs?env=${encodeURIComponent(env)}`;
        return await call("GET", endpoint, LogOverviewResponseSchema, true);
    }

    public static async GetSingleLogMetadata(logUuid: string): Promise<ApiResponse<LogMetadata>> {
        return await call("GET", `/user/logs/${logUuid}/metadata`, LogMetadataSchema, true);
    }
}
