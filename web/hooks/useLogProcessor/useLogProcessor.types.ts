export type UseLogProcessorHook = () => UseLogProcessor;

export interface UseLogProcessor {
    fullLog: string;
    error: string | null;
    isDone: boolean;
    startFetchingLog: (logUUID: string) => void;
    stopFetchingLog: () => void;
}
