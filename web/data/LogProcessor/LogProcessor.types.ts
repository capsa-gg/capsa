export type LogMode = "SingleUnfiltered" | "SingleFiltered";

export type WorkerCommandMessage =
    | { type: "START_FETCHING_LOG"; payload: { logUrlBase: string; jwt: string; filters: LogFilters } }
    | { type: "STOP_FETCHING_LOG"; payload: undefined };

export type WorkerPostMessage =
    | { type: "LOG_CONTENTS_UPDATED"; payload: { fullLog: string } }
    | { type: "ERROR"; payload: { error: string } }
    | { type: "LOG_FETCHING_DONE"; payload: undefined };

export interface LogFilters {
    includedSeverities: string[];
    includedCategories: string[];
    excludedCategories: string[];
}
