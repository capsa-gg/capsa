"use client";

import { useState, useEffect, useCallback } from "react";
import LogProcessor from "@/data/LogProcessor/LogProcessor";
import { UseLogProcessorHook } from "@/hooks/useLogProcessor/useLogProcessor.types";

const useLogProcessor: UseLogProcessorHook = () => {
    const [logWorkerManager] = useState(() => new LogProcessor());
    const [fullLog, setFullLog] = useState("Log not loaded");
    const [error, setError] = useState<string | null>(null);
    const [isDone, setIsDone] = useState(false);

    useEffect(() => {
        logWorkerManager.setOnLogContentsUpdated(newFullLog => {
            setFullLog(newFullLog);
        });

        logWorkerManager.setOnError(newError => {
            setError(newError);
        });

        logWorkerManager.setOnDone(() => {
            setIsDone(true);
        });

        return () => {
            logWorkerManager.stopFetchingLog();
        };
    }, [logWorkerManager]);

    const startFetchingLog = useCallback(
        (id: string) => {
            setError(null);
            setIsDone(false);
            logWorkerManager.startFetchingLog(id);
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
        startFetchingLog,
        stopFetchingLog,
    };
};

export default useLogProcessor;
