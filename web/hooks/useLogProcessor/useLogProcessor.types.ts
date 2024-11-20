import { FilterState } from "@/context/SingleLogContext/SingleLogContext.types";

export type UseLogProcessorHook = () => UseLogProcessor;

export interface UseLogProcessor {
    fullLog: string;
    error: string | null;
    isDone: boolean;
    startFetchingLog: (logUUID: string, filters: FilterState) => void;
    stopFetchingLog: () => void;
}
