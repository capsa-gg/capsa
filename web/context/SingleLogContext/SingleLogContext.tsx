"use client";

import React, { createContext, useCallback, useContext, useEffect, useReducer, useState } from "react";
import {
    SingleLogContextProviderProps,
    SingleLogContextData,
    FilterState,
} from "@/context/SingleLogContext/SingleLogContext.types";
import { useGetSingleLogMetadata } from "@/api/hooks";
import useLogProcessor from "@/hooks/useLogProcessor/useLogProcessor";
import {
    defaultFilterReducerState,
    filterReducer,
    getFilterReducerLocalInitialState,
} from "@/context/SingleLogContext/SingleLogContext.reducers";
import { setSingleLogFilterStateInLocalStorage } from "@/context/SingleLogContext/SingleLogContext.storage";
import { ReadonlyURLSearchParams, usePathname, useRouter, useSearchParams } from "next/navigation";

const searchParamIncludedSeverities = "included_severities";
const searchParamIncludedCategories = "included_categories";
const searchParamExcludedCategories = "excluded_categories";

//@ts-ignore // This is fine, we are checking with the use hook
export const SingleLogContext = createContext<SingleLogContextData>(undefined);

export const SingleLogContextProvider: React.FC<SingleLogContextProviderProps> = ({ logUUID, children }) => {
    const searchParams = useSearchParams();
    const router = useRouter();
    const pathname = usePathname();
    const drawerState = useState(false);
    const filters = useReducer(filterReducer, null, () => getInitialFilterState(searchParams));
    const metadata = useGetSingleLogMetadata(logUUID);
    const logProcessor = useLogProcessor();

    const [filterState] = filters;

    const saveFilters = () => {
        setSingleLogFilterStateInLocalStorage(filterState);

        const params = generateUrlParamString(filterState);
        router.push(`${pathname}?${params}`);

        drawerState[1](false); // Hide drawer
        logProcessor.stopFetchingLog; // Stop fetching current log
        logProcessor.startFetchingLog(logUUID, filterState); // Start fetching log with set filters
    };

    // Start loading logs with the initial filter state on load
    useEffect(() => {
        saveFilters();
    }, []);

    const context: SingleLogContextData = {
        drawerState,
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

    if (!severityParam && !includedCategoryParam && !excludedCategoryParam) {
        console.log("[getInitialFilterState]: no search parameters set, loading local");
        return getFilterReducerLocalInitialState();
    }
    console.log("[getInitialFilterState]: search parameters set, using from url bar");

    const filterStateFromUrl = JSON.parse(JSON.stringify(defaultFilterReducerState)) as FilterState; // Deep clone to prevent object being updated

    if (severityParam) {
        const enabledSeverities = severityParam.split(",");
        filterStateFromUrl.severities = Object.keys(defaultFilterReducerState.severities).reduce(
            (acc, sev) => ({
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

    return filterStateFromUrl;
};

const generateUrlParamString = (filterState: FilterState) => {
    const params = {
        [searchParamIncludedSeverities]: "",
        [searchParamIncludedCategories]: "",
        [searchParamExcludedCategories]: "",
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

    const urlParams = new URLSearchParams(params);
    return urlParams.toString();
};
