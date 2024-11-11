"use client";

// Note(lucianonooijen): this file contains a lot of overly complex gibberish
// There must be a better way to make swr work nicely together with custom fetching functions, but for now this works
// This will need s proper rewrite of the internals while keeping the hook type signatues the same, so I will add a
// TODO: do this properly

import { ApiResponse } from "@/api/apibase";
import useSWR, { SWRResponse } from "swr";
import useSWRMutation, { type SWRMutationResponse } from "swr/mutation";
import { ApplicationError } from "@/types/api/error";
import { UserAuth } from "@/api/userauth";
import { Environments } from "@/api/environments";
import { ListAllEnvironmentsResponse } from "@/types/api/environments";
import { LogOverviewResponse } from "@/types/api/logs";
import { Logs } from "@/api/logs";

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

const apiCallToSwrPlain = <Res>(
    name: string,
    apiCallFunc: () => Promise<ApiResponse<Res>>,
): (() => SWRResponse<Res, ApplicationError, null>) => {
    const fetcher = async () => {
        const [res, err] = await apiCallFunc();
        if (err) {
            throw new ApplicationError(err);
        }
        return res;
    };

    return () => useSWR(name, fetcher);
};

/**
 * SWR HOOKS
 */

// Auth
export const useUserLogin = apiCallToSwrMutation("userlogin", UserAuth.Login);
export const useUserPasswordResetRequest = apiCallToSwrMutation("passwordforgot", UserAuth.PasswordResetStart);
export const useUserPasswordResetComplete = apiCallToSwrMutation("passwordreset", UserAuth.PasswordResetComplete);

// Environments
export const useGetAllEnvironments = apiCallToSwrPlain<ListAllEnvironmentsResponse>(
    "listallenvironments",
    Environments.ListAllEnvironments,
);

// Logs
export const useGetAllLogs = apiCallToSwrPlain<LogOverviewResponse>("getalllogs", Logs.ListAllLogs);
