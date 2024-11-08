"use client";

import { ApiResponse } from "@/api/apibase";
import useSWRMutation, { type SWRMutationResponse } from "swr/mutation";
import { ApplicationError } from "@/types/api/error";
import { UserAuth } from "@/api/userauth";
import { LoginRequest, UserJwtData, UserPasswordResetConfirmation } from "@/types/api/auth";

type ApiCallFunc<Req, Res> = (req: Req) => Promise<ApiResponse<Res>>;

const apiCallToSwrMutation = <Req, Res>(
    id: string,
    apiCallFunc: ApiCallFunc<Req, Res>,
): (() => SWRMutationResponse<Res, ApplicationError, string, Req>) => {
    const fetcher = async (_: string, { arg }: { arg: Req }) => {
        const [res, err] = await apiCallFunc(arg);
        if (err) {
            throw new ApplicationError(err);
        }
        return res;
    };

    return () => useSWRMutation(id, fetcher);
};

// Add hooks below

export const useUserLogin = apiCallToSwrMutation("userlogin", UserAuth.Login);
export const useUserPasswordResetRequest = apiCallToSwrMutation("passwordforgot", UserAuth.PasswordResetStart);
export const useUserPasswordResetComplete = apiCallToSwrMutation("passwordreset", UserAuth.PasswordResetComplete);
