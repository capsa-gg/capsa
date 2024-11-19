"use client";

import { ApiResponse, call } from "@/api/apibase";
import { ListAllEnvironmentsResponse, ListAllEnvironmentsResponseSchema } from "@/types/api/environments";
import { LogMetadata, LogMetadataSchema, LogOverviewResponse, LogOverviewResponseSchema } from "@/types/api/logs";

export class Logs {
    public static async ListAllLogs(): Promise<ApiResponse<LogOverviewResponse>> {
        return await call("GET", "/user/logs", LogOverviewResponseSchema, true);
    }

    public static async GetSingleLogMetadata(logUuid: string): Promise<ApiResponse<LogMetadata>> {
        return await call("GET", `/user/logs/${logUuid}/metadata`, LogMetadataSchema, true);
    }
}
