"use client";

import React, { createContext, useCallback, useContext, useEffect, useReducer, useState } from "react";
import { SingleLogContextProviderProps, SingleLogContextData } from "@/context/SingleLogContext/SingleLogContext.types";
import { useGetSingleLogMetadata } from "@/api/hooks";
import useLogProcessor from "@/hooks/useLogProcessor/useLogProcessor";
import { filterReducer, getFilterReducerInitialState } from "@/context/SingleLogContext/SingleLogContext.reducers";
import { setSingleLogFilterStateInLocalStorage } from "@/context/SingleLogContext/SingleLogContext.storage";

//@ts-ignore // This is fine, we are checking with the use hook
export const SingleLogContext = createContext<SingleLogContextData>(undefined);

export const SingleLogContextProvider: React.FC<SingleLogContextProviderProps> = ({ logUUID, children }) => {
    const drawerState = useState(false);
    const filters = useReducer(filterReducer, getFilterReducerInitialState());
    const metadata = useGetSingleLogMetadata(logUUID);
    const logProcessor = useLogProcessor();

    const saveFilters = useCallback(() => {
        const [filterState] = filters;
        setSingleLogFilterStateInLocalStorage(filterState);
        drawerState[1](false); // Hide drawer
        logProcessor.stopFetchingLog; // Stop fetching current log
        logProcessor.startFetchingLog(logUUID, filterState); // Start fetching log with set filters
    }, [logUUID, logProcessor.startFetchingLog, logProcessor.stopFetchingLog, filters]);

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
