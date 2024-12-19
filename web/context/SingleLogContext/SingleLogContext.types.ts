import type { UseLogProcessor } from "@/hooks/useLogProcessor/useLogProcessor.types";
import type { ApplicationError } from "@/types/api/error";
import type { LogMetadata } from "@/types/api/logs";
import type React from "react";
import type { Dispatch, SetStateAction } from "react";
import type { SWRResponse } from "swr";
import { z } from "zod";

export interface SingleLogContextProviderProps {
    logUUID: string;
    children: React.ReactNode;
}

export type SingleLogViewMode = "Log" | "Metadata";

export interface SingleLogContextData {
    viewMode: [SingleLogViewMode, Dispatch<SetStateAction<SingleLogViewMode>>];
    filterDrawerState: [boolean, Dispatch<SetStateAction<boolean>>];
    mergeDrawerState: [boolean, Dispatch<SetStateAction<boolean>>];
    metadata: SWRResponse<LogMetadata, ApplicationError, null>;
    logProcessor: UseLogProcessor;
    filters: [FilterState, Dispatch<FilterAction>];
    saveFilters: () => void;
}

export const FilterStateSchema = z.object({
    severities: z.record(z.string(), z.boolean()),
    includedCategories: z.array(z.string()),
    excludedCategories: z.array(z.string()),
    mergedLogs: z.array(z.string()),
});

export type FilterState = z.infer<typeof FilterStateSchema>;

export type ResetFiltersAction = { type: "RESET_FILTERS" };
export type SwitchSeverityAction = { type: "SWITCH_SEVERITY"; severity: string };
export type CategoryAction =
    | { type: "INCLUDED_CATEGORY_ADD"; category: string }
    | { type: "INCLUDED_CATEGORY_REMOVE"; category: string }
    | { type: "INCLUDED_CATEGORY_CLEAR" }
    | { type: "EXCLUDED_CATEGORY_ADD"; category: string }
    | { type: "EXCLUDED_CATEGORY_REMOVE"; category: string }
    | { type: "EXCLUDED_CATEGORY_CLEAR" };
export type MergedLogActionAction = { type: "MERGED_LOGS_SWITCH"; log: string } | { type: "MERGED_LOGS_CLEAR" };

export type FilterAction = ResetFiltersAction | SwitchSeverityAction | CategoryAction | MergedLogActionAction;
