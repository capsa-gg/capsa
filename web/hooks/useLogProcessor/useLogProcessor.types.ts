import type { FilterState } from "@/context/SingleLogContext/SingleLogContext.types";

export type UseLogProcessorHook = () => UseLogProcessor;

export interface UseLogProcessor {
    fullLog: string;
    error: string | null;
    isDone: boolean;
    logMode: string | null;
    absoluteLineNumbers: number[] | null;
    startFetchingLog: (logUUID: string, filters: FilterState) => void;
    stopFetchingLog: () => void;
}
