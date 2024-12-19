"use client";

// Note(lucianonooijen): this file contains a lot of overly complex gibberish
// There must be a better way to make swr work nicely together with custom fetching functions, but for now this works
// This will need s proper rewrite of the internals while keeping the hook type signatures the same, so I will add a

import { Admin } from "@/api/admin";
import type { ApiResponse } from "@/api/apibase";
import { Environments } from "@/api/environments";
import { Logs } from "@/api/logs";
import { Search } from "@/api/search";
import { UserAuth } from "@/api/userauth";
import type { ListAllEnvironmentsResponse } from "@/types/api/environments";
import { ApplicationError } from "@/types/api/error";
import type { LogMetadata, LogOverviewResponse } from "@/types/api/logs";
import type { SearchResults } from "@/types/api/search";
import type {
    AddEnvironmentRequest,
    AddTitleRequest,
    AddTitleResponse,
    ListAllTitlesResponse,
} from "@/types/api/titleenvironment";
import type { ListAllUsersResponse } from "@/types/api/users";
import useSWR, { type SWRResponse } from "swr";
import useSWRMutation, { type SWRMutationResponse } from "swr/mutation";

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

const apiCallToSwrWithStringArg = <Res>(
    baseName: string,
    apiCallFunc: (arg: string) => Promise<ApiResponse<Res>>,
): ((arg: string) => SWRResponse<Res, ApplicationError, null>) => {
    const fetcher = async (arg: string) => {
        // Remove basename, which is required in the useSWR hook to not have swr cache values
        const argValue = arg.slice(baseName.length + 1); // +1 for the `-` character

        const [res, err] = await apiCallFunc(argValue);
        if (err) {
            throw new ApplicationError(err);
        }
        return res;
    };

    return (id: string) => useSWR(`${baseName}-${id}`, fetcher);
};

/**
 * SWR HOOKS
 */

// Auth
export const useUserLogin = apiCallToSwrMutation("userlogin", UserAuth.Login);
export const useUserPasswordResetRequest = apiCallToSwrMutation("passwordforgot", UserAuth.PasswordResetStart);
export const useUserPasswordResetComplete = apiCallToSwrMutation("passwordreset", UserAuth.PasswordResetComplete);

// Search
export const useSearch = apiCallToSwrWithStringArg<SearchResults>("search", Search.Search);

// Admin
export const useGetAllTitles = apiCallToSwrPlain<ListAllTitlesResponse>("listalltitles", Admin.ListAllTitles);
export const useAddTitle = apiCallToSwrMutation<AddTitleRequest, AddTitleResponse>("addtitle", Admin.AddTitle);
export const useAddEnvironment = apiCallToSwrMutation<AddEnvironmentRequest, ListAllEnvironmentsResponse>(
    "addenvironment",
    Admin.AddEnvironment,
);
export const useGetAllUsers = apiCallToSwrPlain<ListAllUsersResponse>("admingetallusers", Admin.GetAllUsers);

// Environments
export const useGetAllEnvironments = apiCallToSwrPlain<ListAllEnvironmentsResponse>(
    "listallenvironments",
    Environments.ListAllEnvironments,
);

// Logs
export const getLogsOverview = apiCallToSwrWithStringArg<LogOverviewResponse>("getlogsoverview", Logs.ListLogs);
export const useGetSingleLogMetadata = apiCallToSwrWithStringArg<LogMetadata>(
    "singlelogmetadata",
    Logs.GetSingleLogMetadata,
);
