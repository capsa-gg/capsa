"use client";

import { LogFilters, WorkerCommandMessage, WorkerPostMessage } from "./LogProcessor.types";
import { getRequestUrl } from "@/api/apibase";
import { getJwtFromLocalStorage } from "@/data/jwt/localStorage";

type LogContentsUpdatedCallback = (fullLog: string) => void;
type ErrorCallback = (error: string) => void;
type DoneCallback = () => void;

export default class LogProcessor {
    private worker: Worker;

    private logContentsUpdatedCallback: LogContentsUpdatedCallback | null = null;
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
                    if (this.logContentsUpdatedCallback) {
                        this.logContentsUpdatedCallback(payload.fullLog);
                    }
                    break;
                case "ERROR":
                    if (this.errorCallback) {
                        this.errorCallback(payload.error);
                    }
                    break;
                case "LOG_FETCHING_DONE":
                    if (this.doneCallback) {
                        this.doneCallback();
                    }
                    break;
            }
        });
    }

    public async startFetchingLog(id: string, filters: LogFilters) {
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
        this.worker.postMessage(message);
    }

    public setOnLogContentsUpdated(callback: LogContentsUpdatedCallback) {
        this.logContentsUpdatedCallback = callback;
    }

    public setOnError(callback: ErrorCallback) {
        this.errorCallback = callback;
    }

    public setOnDone(callback: DoneCallback) {
        this.doneCallback = callback;
    }
}
