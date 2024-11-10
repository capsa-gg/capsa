"use client";

import z, { ZodSchema } from "zod";
import ApiError, { ErrorResponseSchema } from "@/types/api/error";
import { getJwtFromLocalStorage } from "@/data/jwt/localStorage";
import { getEnv } from "@/data/env";

class BaseUrl {
    private static _baseUrl?: string;

    public static async getBaseUrl(): Promise<string> {
        if (this._baseUrl) return this._baseUrl;

        const env = await getEnv();
        if (!env || !env.serverUrl) {
            throw new Error("cannot get server url");
        }

        this._baseUrl = env.serverUrl;
        return env.serverUrl;
    }
}

export const getRequestUrl = async (path: string) => {
    const baseUrl = await BaseUrl.getBaseUrl();
    return `${baseUrl}/v1${path}`;
};

export const returnErrorResponseData = async (res: Response, reqUrl: string): Promise<ApiResponseError> => {
    try {
        const resErr = await res.json();
        const resParsed = ErrorResponseSchema.parse(resErr);
        const error: ApiError = {
            title: `Request to endpoint ${reqUrl} failed`,
            severity: "error",
            ...resParsed,
        };
        return [null, error];
    } catch (e) {
        console.error(e);
        return formatUnexpectedErrorResponse(reqUrl, e);
    }
};

// Throws on error
const callWithResponse = async <TReq, TResSchema extends ZodSchema>(
    method: Method,
    path: string,
    resZodSchema: TResSchema,
    body?: TReq,
    jwt?: string,
): Promise<ApiResponse<z.infer<TResSchema>>> => {
    const reqUrl = await getRequestUrl(path);

    const res = await fetch(reqUrl, {
        method,
        headers: {
            accept: "application/json",
            authorization: jwt ? `Bearer ${jwt}` : "",
        },
        body: body ? JSON.stringify(body) : null,
    });

    // Error from API
    if (!res.ok) {
        return await returnErrorResponseData(res, reqUrl);
    }

    // Parse response body and verify types
    const resJson = await res.json();
    const { success, data, error } = resZodSchema.safeParse(resJson);

    // Process zod results and return
    if (!success) {
        const errorRes: ApiError = {
            title: `Request to endpoint ${reqUrl} failed`,
            severity: "error",
            error: "Could not decode response JSON payload",
            details: error.toString(),
        };
        return [null, errorRes];
    }

    return [data, null];
};

const callWithoutResponse = async <TReq>(
    method: Method,
    path: string,
    body?: TReq,
    jwt?: string,
): Promise<ApiResponse<null>> => {
    const reqUrl = await getRequestUrl(path);

    const res = await fetch(reqUrl, {
        method,
        headers: {
            accept: "application/json",
            authorization: jwt ? `Bearer ${jwt}` : "",
        },
        body: body ? JSON.stringify(body) : null,
    });

    // Error from API
    if (!res.ok) {
        return await returnErrorResponseData(res, reqUrl);
    }

    return [null, null];
};

export const call = async <TReq, TResSchema extends ZodSchema>(
    method: Method,
    path: string,
    resZodSchema: TResSchema | null,
    authenticate: boolean,
    body?: TReq,
): Promise<ApiResponse<z.infer<TResSchema>>> => {
    try {
        if (resZodSchema) {
            if (authenticate) {
                const jwtData = getJwtFromLocalStorage();
                const jwt = jwtData?.token;
                return await callWithResponse(method, path, resZodSchema, body, jwt);
            } else {
                return await callWithResponse(method, path, resZodSchema, body);
            }
        }

        if (authenticate) {
            const jwtData = getJwtFromLocalStorage();
            const jwt = jwtData?.token;
            return await callWithoutResponse(method, path, body, jwt);
        } else {
            return await callWithoutResponse(method, path, body);
        }
    } catch (e) {
        console.error(e);
        return formatUnexpectedErrorResponse(path, e);
    }
};

export const formatUnexpectedErrorResponse = (path: string, e: any): ApiResponseError => {
    const error: ApiError = {
        title: `Request to endpoint ${path} failed`,
        severity: "error",
        error: "Unexpected error occurred",
        details: `${e}`,
    };
    return [null, error];
};

export type Method = "GET" | "POST" | "PUT" | "PATCH" | "DELETE";

export type ApiResponseSuccess<T> = [T, null];
export type ApiResponseError = [null, ApiError];
export type ApiResponse<T> = ApiResponseSuccess<T> | ApiResponseError;
