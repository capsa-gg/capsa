"use client";

import type { FilterState } from "@/context/SingleLogContext/SingleLogContext.types";
import LogProcessor from "@/data/LogProcessor/LogProcessor";
import type { LogFilters, LogMode } from "@/data/LogProcessor/LogProcessor.types";
import type { UseLogProcessorHook } from "@/hooks/useLogProcessor/useLogProcessor.types";
import { useCallback, useEffect, useState } from "react";

const useLogProcessor: UseLogProcessorHook = () => {
    const [logWorkerManager] = useState(() => new LogProcessor());
    const [fullLog, setFullLog] = useState("Log not loaded");
    const [absoluteLineNumbers, setAbsoluteLineNumbers] = useState<number[] | null>(null);
    const [logMode, setLogMode] = useState<LogMode | null>(null);
    const [error, setError] = useState<string | null>(null);
    const [isDone, setIsDone] = useState(false);

    useEffect(() => {
        logWorkerManager.setOnLogContentsUpdated((newFullLog, lineNumbers) => {
            setFullLog(newFullLog);
            setAbsoluteLineNumbers(lineNumbers);
        });

        logWorkerManager.setOnError(newError => {
            setError(newError);
        });

        logWorkerManager.setOnDone(() => {
            setIsDone(true);
        });

        logWorkerManager.setOnLogModeReceived(logModeVal => {
            setLogMode(logModeVal);
        });

        return () => {
            logWorkerManager.stopFetchingLog();
        };
    }, [logWorkerManager]);

    const startFetchingLog = useCallback(
        (id: string, filters: FilterState) => {
            setError(null);
            setIsDone(false);
            logWorkerManager.startFetchingLog(id, filterStateToWorkerFilters(filters));
        },
        [logWorkerManager],
    );

    const stopFetchingLog = useCallback(() => {
        logWorkerManager.stopFetchingLog();
    }, [logWorkerManager]);

    return {
        fullLog,
        error,
        isDone,
        logMode,
        absoluteLineNumbers,
        startFetchingLog,
        stopFetchingLog,
    };
};

export default useLogProcessor;

const filterStateToWorkerFilters = (filterState: FilterState): LogFilters => {
    const getIncludedSeverities = (): string[] => {
        // Note: this prevents the included severities from being set even if all severities are enabled
        // We only want to set this if we need to filter by serverity as well
        const hasDisabledSeverities = Object.values(filterState.severities).filter(enabled => !enabled).length > 0;

        if (!hasDisabledSeverities) {
            return [];
        }

        // We have disabled severities, so we need to build the enabled severities arg
        return Object.entries(filterState.severities)
            .filter(([_, enabled]) => enabled)
            .map(([sev, _]) => sev);
    };

    return {
        includedSeverities: getIncludedSeverities(),
        includedCategories: filterState.includedCategories,
        excludedCategories: filterState.excludedCategories,
    };
};
