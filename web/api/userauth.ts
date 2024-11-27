"use client";

import {
    type ApiResponse,
    call,
    formatUnexpectedErrorResponse,
    getRequestUrl,
    returnErrorResponseData,
} from "@/api/apibase";
import {
    type LoginRequest,
    type UserJwtData,
    UserJwtDataSchema,
    type UserPasswordResetConfirmation,
} from "@/types/api/auth";

// biome-ignore lint/complexity/noStaticOnlyClass: desired here
export class UserAuth {
    public static async Login(req: LoginRequest): Promise<ApiResponse<UserJwtData>> {
        return await call("POST", "/user/auth/login", UserJwtDataSchema, false, req);
    }

    // Due to string params a separate helper
    // TODO: make a generic helper for this
    public static async PasswordResetStart({ email }: { email: string }): Promise<ApiResponse<null>> {
        const endpoint = "/user/auth/password-reset";
        const reqUrl = await getRequestUrl(endpoint);
        const url = new URL(reqUrl);
        url.searchParams.append("email", email);

        try {
            const res = await fetch(url);
            if (!res.ok) {
                return returnErrorResponseData(res, endpoint);
            }
            return [null, null];
        } catch (e) {
            return formatUnexpectedErrorResponse(endpoint, e);
        }
    }

    public static async PasswordResetComplete(req: UserPasswordResetConfirmation): Promise<ApiResponse<null>> {
        return await call("POST", "/user/auth/password-reset", null, false, req);
    }
}
