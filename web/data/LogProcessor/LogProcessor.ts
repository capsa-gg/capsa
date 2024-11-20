"use client";

import { LogFilters, LogMode, WorkerCommandMessage, WorkerPostMessage } from "./LogProcessor.types";
import { getRequestUrl } from "@/api/apibase";
import { getJwtFromLocalStorage } from "@/data/jwt/localStorage";

type LogContentsUpdatedCallback = (fullLog: string, absoluteLineNumbers: number[] | null) => void;
type LogModeReceivedCallback = (logMode: LogMode | null) => void;
type ErrorCallback = (error: string) => void;
type DoneCallback = () => void;

export default class LogProcessor {
    private worker: Worker;

    private logContentsUpdatedCallback: LogContentsUpdatedCallback | null = null;
    private logModeReceivedCallback: LogModeReceivedCallback | null = null;
    private errorCallback: ErrorCallback | null = null;
    private doneCallback: DoneCallback | null = null;

    constructor() {
        this.worker = new Worker(new URL("./LogProcessor.worker.ts", import.meta.url));
        this.setupEventListeners();
    }

    private setupEventListeners() {
        this.worker.addEventListener("message", event => {
            const { type, payload } = event.data as WorkerPostMessage;

            switch (type) {
                case "LOG_CONTENTS_UPDATED":
                    console.log(`[LogProcessor]: LOG_CONTENTS_UPDATED`);
                    if (this.logContentsUpdatedCallback) {
                        this.logContentsUpdatedCallback(payload.fullLog, payload.lineNumbers);
                    }
                    break;
                case "LOG_MODE_RECEIVED":
                    console.log(`[LogProcessor]: LOG_MODE_RECEIVED, ${payload.logMode}`);
                    if (this.logModeReceivedCallback) {
                        this.logModeReceivedCallback(payload.logMode);
                    }
                    break;
                case "ERROR":
                    console.log(`[LogProcessor]: ERROR, ${payload.error}`);
                    if (this.errorCallback) {
                        this.errorCallback(payload.error);
                    }
                    break;
                case "LOG_FETCHING_DONE":
                    console.log(`[LogProcessor]: LOG_FETCHING_DONE`);
                    if (this.doneCallback) {
                        this.doneCallback();
                    }
                    break;
            }
        });
    }

    public async startFetchingLog(id: string, filters: LogFilters) {
        if (this.logModeReceivedCallback) {
            this.logModeReceivedCallback(null);
        }
        if (this.logContentsUpdatedCallback) {
            this.logContentsUpdatedCallback("", null);
        }

        const path = `/user/logs/${id}/log`;
        const reqUrl = await getRequestUrl(path);

        const jwtData = getJwtFromLocalStorage();
        const jwt = jwtData?.token;
        if (!jwt) {
            if (this.errorCallback) {
                this.errorCallback("JWT cannot be fetched using getJwtFromLocalStorage");
            }
            return;
        }

        const message: WorkerCommandMessage = {
            type: "START_FETCHING_LOG",
            payload: { logUrlBase: reqUrl, jwt, filters },
        };
        this.worker.postMessage(message);
    }

    public stopFetchingLog() {
        const message: WorkerCommandMessage = { type: "STOP_FETCHING_LOG", payload: undefined };
        if (this.logModeReceivedCallback) {
            this.logModeReceivedCallback(null);
        }
        this.worker.postMessage(message);
    }

    public setOnLogContentsUpdated(callback: LogContentsUpdatedCallback) {
        this.logContentsUpdatedCallback = callback;
    }

    public setOnLogModeReceived(callback: LogModeReceivedCallback) {
        this.logModeReceivedCallback = callback;
    }

    public setOnError(callback: ErrorCallback) {
        this.errorCallback = callback;
    }

    public setOnDone(callback: DoneCallback) {
        this.doneCallback = callback;
    }
}
