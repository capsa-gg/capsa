import React, { Dispatch, Reducer, SetStateAction } from "react";
import { SWRResponse } from "swr";
import { ApplicationError } from "@/types/api/error";
import { LogMetadata } from "@/types/api/logs";
import { UseLogProcessor } from "@/hooks/useLogProcessor/useLogProcessor.types";

export interface SingleLogContextProviderProps {
    logUUID: string;
    children: React.ReactNode;
}

export interface SingleLogContextData {
    drawerState: [boolean, Dispatch<SetStateAction<boolean>>];
    metadata: SWRResponse<LogMetadata, ApplicationError, null>;
    logProcessor: UseLogProcessor;
    filters: [FilterState, Dispatch<FilterAction>];
    saveFilters: () => void;
}

export interface FilterState {
    severities: Record<string, boolean>;
}

export type ResetFiltersAction = { type: "RESET_FILTERS" };
export type SwitchSeverityAction = { type: "SWITCH_SEVERITY"; severity: string };

export type FilterAction = ResetFiltersAction | SwitchSeverityAction;
