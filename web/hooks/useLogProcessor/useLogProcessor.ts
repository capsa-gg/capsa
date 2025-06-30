"use client";

import * as Comlink from "comlink";
import { useCallback, useEffect, useRef, useState } from "react";
import type { FilterState } from "@/context/SingleLogContext/SingleLogContext.types";
import { getEnv } from "@/data/env";
import { getJwtFromLocalStorage } from "@/data/jwt/localStorage";
import { createLogProcessor } from "@/data/LogProcessor/LogProcessor";
import type { LogFilters, LogMode } from "@/data/LogProcessor/LogProcessor.types";

const useLogProcessor = () => {
    const [fullLog, setFullLog] = useState("Log not loaded");
    const [absoluteLineNumbers, setAbsoluteLineNumbers] = useState<string[] | null>(null);
    const [logMode, setLogMode] = useState<LogMode | null>(null);
    const [error, setError] = useState<string | null>(null);
    const [isDone, setIsDone] = useState(false);

    const workerRef = useRef<Worker | null>(null);

    const startFetchingLog = useCallback(async (id: string, filters: FilterState) => {
        setError(null);
        setIsDone(false);
        setFullLog("Loading...");
        setAbsoluteLineNumbers(null);

        try {
            const env = await getEnv();

            const jwt = getJwtFromLocalStorage()?.token;
            if (!jwt) throw new Error("JWT not found");

            const { api, worker } = await createLogProcessor();
            workerRef.current = worker;

            await api.fetchLog(
                `${env.serverUrl}/v1/user/logs/${id}/log`,
                jwt,
                filterStateToWorkerFilters(filters),
                Comlink.proxy((newFullLog, lineNumbers) => {
                    setFullLog(newFullLog);
                    setAbsoluteLineNumbers(lineNumbers);
                }),
                Comlink.proxy(mode => {
                    setLogMode(mode);
                }),
            );

            setIsDone(true);
        } catch (err) {
            setError(`${err}`);
        }
    }, []);

    const stopFetchingLog = useCallback(() => {
        if (workerRef.current) {
            workerRef.current.terminate();
            workerRef.current = null;
        }
    }, []);

    useEffect(() => {
        return () => {
            stopFetchingLog();
        };
    }, [stopFetchingLog]);

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
        mergeLogs: filterState.mergedLogs,
    };
};
