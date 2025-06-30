"use client";

import * as Comlink from "comlink";
import type { LogFilters, LogMode } from "./LogProcessor.types";

export type LogProcessorApi = {
    fetchLog: (
        logUrlBase: string,
        jwt: string,
        filters: LogFilters,
        onUpdate: (chunk: string, lineNumbers: string[] | null) => void,
        onMode: (mode: LogMode) => void,
    ) => Promise<void>;
};

export const createLogProcessor = async () => {
    const worker = new Worker(new URL("./LogProcessor.worker.ts", import.meta.url), { type: "module" });
    const api = Comlink.wrap<LogProcessorApi>(worker);
    return { api, worker };
};
