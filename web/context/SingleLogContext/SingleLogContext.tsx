"use client";

import { type ReadonlyURLSearchParams, usePathname, useRouter, useSearchParams } from "next/navigation";
import type React from "react";
import { createContext, useContext, useEffect, useReducer, useState } from "react";
import { useGetSingleLogMetadata } from "@/api/hooks";
import {
    defaultFilterReducerState,
    filterReducer,
    getFilterReducerLocalInitialState,
} from "@/context/SingleLogContext/SingleLogContext.reducers";
import { setSingleLogFilterStateInLocalStorage } from "@/context/SingleLogContext/SingleLogContext.storage";
import type {
    FilterState,
    SingleLogContextData,
    SingleLogContextProviderProps,
    SingleLogViewMode,
} from "@/context/SingleLogContext/SingleLogContext.types";
import useLogProcessor from "@/hooks/useLogProcessor/useLogProcessor";

const searchParamIncludedSeverities = "included_severities";
const searchParamIncludedCategories = "included_categories";
const searchParamExcludedCategories = "excluded_categories";
const searchParamMergeLogs = "merge_logs";

//@ts-ignore // This is fine, we are checking with the use hook
export const SingleLogContext = createContext<SingleLogContextData>(undefined);

export const SingleLogContextProvider: React.FC<SingleLogContextProviderProps> = ({ logUUID, children }) => {
    const searchParams = useSearchParams();
    const router = useRouter();
    const pathname = usePathname();
    const filterDrawerState = useState(false);
    const mergeDrawerState = useState(false);
    const viewMode = useState<SingleLogViewMode>("Log");
    const filters = useReducer(filterReducer, null, () => getInitialFilterState(searchParams));
    const metadata = useGetSingleLogMetadata(logUUID);
    const logProcessor = useLogProcessor();

    const [filterState] = filters;

    const saveFilters = () => {
        console.log("Saving filters!");
        setSingleLogFilterStateInLocalStorage(filterState);

        const params = generateUrlParamString(filterState);
        router.push(`${pathname}?${params}`);

        filterDrawerState[1](false); // Hide drawer
        mergeDrawerState[1](false); // Hide drawer
        logProcessor.stopFetchingLog(); // Stop fetching current log
        logProcessor.startFetchingLog(logUUID, filterState); // Start fetching log with set filters
    };

    // Start loading logs with the initial filter state on load
    // biome-ignore lint/correctness/useExhaustiveDependencies: fine here, for now
    useEffect(() => {
        // HACK
        setTimeout(saveFilters, 2000);
    }, []);

    const context: SingleLogContextData = {
        filterDrawerState,
        mergeDrawerState,
        viewMode,
        metadata,
        logProcessor,
        filters,
        saveFilters,
    };

    return <SingleLogContext.Provider value={context}>{children}</SingleLogContext.Provider>;
};

export const useSingleLogContext = () => {
    const context = useContext(SingleLogContext);
    if (!context) {
        throw new Error("useSingleLogContext must be used within an SingleLogContextProvider");
    }
    return context;
};

const getInitialFilterState = (searchParams: ReadonlyURLSearchParams): FilterState => {
    const severityParam = searchParams.get(searchParamIncludedSeverities);
    const includedCategoryParam = searchParams.get(searchParamIncludedCategories);
    const excludedCategoryParam = searchParams.get(searchParamExcludedCategories);
    const mergeLogsParam = searchParams.get(searchParamMergeLogs);

    if (!severityParam && !includedCategoryParam && !excludedCategoryParam && !mergeLogsParam) {
        console.log("[getInitialFilterState]: no search parameters set, loading local");
        return getFilterReducerLocalInitialState();
    }
    console.log("[getInitialFilterState]: search parameters set, using from url bar");

    const filterStateFromUrl = JSON.parse(JSON.stringify(defaultFilterReducerState)) as FilterState; // Deep clone to prevent object being updated

    if (severityParam) {
        const enabledSeverities = severityParam.split(",");

        filterStateFromUrl.severities = Object.keys(defaultFilterReducerState.severities).reduce(
            (acc, sev) => ({
                // biome-ignore lint/performance/noAccumulatingSpread: this is fine here
                ...acc,
                [sev]: enabledSeverities.includes(sev),
            }),
            {},
        );
    }

    if (includedCategoryParam) {
        filterStateFromUrl.includedCategories = includedCategoryParam.split(",");
    }

    if (excludedCategoryParam) {
        filterStateFromUrl.excludedCategories = excludedCategoryParam.split(",");
    }

    if (mergeLogsParam) {
        filterStateFromUrl.mergedLogs = mergeLogsParam.split(",");
    }

    return filterStateFromUrl;
};

const generateUrlParamString = (filterState: FilterState) => {
    const params = {
        [searchParamIncludedSeverities]: "",
        [searchParamIncludedCategories]: "",
        [searchParamExcludedCategories]: "",
        [searchParamMergeLogs]: "",
    };

    // At least one disabled category to be added
    if (Object.values(filterState.severities).includes(false)) {
        const severitiesParamString = Object.entries(filterState.severities)
            .filter(([_, enabled]) => enabled)
            .map(([sev, _]) => sev)
            .join(",");

        params[searchParamIncludedSeverities] = severitiesParamString;
    }

    if (filterState.includedCategories.length > 0) {
        params[searchParamIncludedCategories] = filterState.includedCategories.join(",");
    }

    if (filterState.excludedCategories.length > 0) {
        params[searchParamExcludedCategories] = filterState.excludedCategories.join(",");
    }

    if (filterState.mergedLogs.length > 0) {
        params[searchParamMergeLogs] = filterState.mergedLogs.join(",");
    }

    const urlParams = new URLSearchParams(params);
    return urlParams.toString();
};
