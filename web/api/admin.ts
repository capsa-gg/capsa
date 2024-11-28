"use client";

import { type ApiResponse, call } from "@/api/apibase";
import { type ListAllEnvironmentsResponse, ListAllEnvironmentsResponseSchema } from "@/types/api/environments";
import {
    type AddEnvironmentRequest,
    type AddTitleRequest,
    type AddTitleResponse,
    AddTitleResponseSchema,
    type ListAllTitlesResponse,
    ListAllTitlesResponseSchema,
} from "@/types/api/titleenvironment";

// biome-ignore lint/complexity/noStaticOnlyClass: desired here
export class Admin {
    public static async ListAllTitles(): Promise<ApiResponse<ListAllTitlesResponse>> {
        return await call("GET", "/admin/titles", ListAllTitlesResponseSchema, true);
    }

    public static async AddTitle(req: AddTitleRequest): Promise<ApiResponse<AddTitleResponse>> {
        return await call("POST", "/admin/titles", AddTitleResponseSchema, true, req);
    }

    public static async AddEnvironment(req: AddEnvironmentRequest): Promise<ApiResponse<ListAllEnvironmentsResponse>> {
        return await call("POST", "/admin/environments", ListAllEnvironmentsResponseSchema, true, req);
    }
}
