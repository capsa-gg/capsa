"use client";

import { getRequestUrl } from "@/api/apibase";
import { getJwtFromLocalStorage } from "@/data/jwt/localStorage";
import type { LogFilters, LogMode, WorkerCommandMessage, WorkerPostMessage } from "./LogProcessor.types";

type LogContentsUpdatedCallback = (fullLog: string, addedChunk: string, absoluteLineNumbers: string[] | null) => void;
type LogModeReceivedCallback = (logMode: LogMode | null) => void;
type ErrorCallback = (error: string) => void;
type DoneCallback = () => void;

import * as Comlink from "comlink";

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
