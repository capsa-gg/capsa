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
import {
    type CreateUserRequest,
    type ListAllUsersResponse,
    ListAllUsersResponseSchema,
    type SingleUserResponse,
    SingleUserResponseSchema,
    type UpdateUserRequest,
} from "@/types/api/users";

// biome-ignore lint/complexity/noStaticOnlyClass: desired here
export class Admin {
    // Titles and environments
    public static async ListAllTitles(): Promise<ApiResponse<ListAllTitlesResponse>> {
        return await call("GET", "/admin/titles", ListAllTitlesResponseSchema, true);
    }

    public static async AddTitle(req: AddTitleRequest): Promise<ApiResponse<AddTitleResponse>> {
        return await call("POST", "/admin/titles", AddTitleResponseSchema, true, req);
    }

    public static async AddEnvironment(req: AddEnvironmentRequest): Promise<ApiResponse<ListAllEnvironmentsResponse>> {
        return await call("POST", "/admin/environments", ListAllEnvironmentsResponseSchema, true, req);
    }

    public static async GetAllUsers(): Promise<ApiResponse<ListAllUsersResponse>> {
        return await call("GET", "/admin/users", ListAllUsersResponseSchema, true);
    }

    public static async CreateUser(req: CreateUserRequest): Promise<ApiResponse<SingleUserResponse>> {
        return await call("POST", "/admin/users", SingleUserResponseSchema, true, req);
    }

    public static async UpdateUser(userUUID: string, req: UpdateUserRequest): Promise<ApiResponse<SingleUserResponse>> {
        return await call("PUT", `/admin/users/${userUUID}`, SingleUserResponseSchema, true, req);
    }

    public static async DeactivateUser(userUUID: string): Promise<ApiResponse<SingleUserResponse>> {
        return await call("DELETE", `/admin/users/${userUUID}/activation`, SingleUserResponseSchema, true);
    }

    public static async ReactivateUser(userUUID: string): Promise<ApiResponse<SingleUserResponse>> {
        return await call("POST", `/admin/users/${userUUID}/activation`, SingleUserResponseSchema, true);
    }
}
