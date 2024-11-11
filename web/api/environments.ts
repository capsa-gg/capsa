"use client";

import { ApiResponse, call } from "@/api/apibase";
import { ListAllEnvironmentsResponse, ListAllEnvironmentsResponseSchema } from "@/types/api/environments";

export class Environments {
    public static async ListAllEnvironments(): Promise<ApiResponse<ListAllEnvironmentsResponse>> {
        return await call("GET", "/user/environments", ListAllEnvironmentsResponseSchema, true);
    }
}
