"use client";

import { type ApiResponse, call } from "@/api/apibase";
import { type ListAllEnvironmentsResponse, ListAllEnvironmentsResponseSchema } from "@/types/api/environments";

// biome-ignore lint/complexity/noStaticOnlyClass: desired here
export class Environments {
    public static async ListAllEnvironments(): Promise<ApiResponse<ListAllEnvironmentsResponse>> {
        return await call("GET", "/user/environments", ListAllEnvironmentsResponseSchema, true);
    }
}
