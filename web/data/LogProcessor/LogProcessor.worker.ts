"use client";

import { getJwtFromLocalStorage } from "@/data/jwt/localStorage";
import type { LogMode, WorkerCommandMessage, WorkerPostMessage } from "./LogProcessor.types";

let abortController: AbortController | null = null;

self.addEventListener("message", async (event: MessageEvent<WorkerCommandMessage>) => {
    const { type, payload } = event.data;

    switch (type) {
        case "START_FETCHING_LOG":
            if (payload?.reqUrl) {
                await fetchLog(payload.reqUrl, payload.jwt);
            }
            break;
        case "STOP_FETCHING_LOG":
            if (abortController) {
                abortController.abort();
            }
            break;
    }
});

async function fetchLog(reqUrl: string, jwt: string): Promise<void> {
    abortController = new AbortController();

    try {
        const res = await fetch(reqUrl, {
            method: "GET",
            headers: {
                accept: "text/plain",
                authorization: `Bearer ${jwt}`,
            },
            signal: abortController.signal,
        });

        if (!res.ok) {
            throw new Error(`API call responded with code ${res.status} ${res.statusText}`);
        }

        const logMode = res.headers.get("x-capsa-log-mode") as LogMode;
        if (!logMode) {
            throw new Error("Log mode header not set");
        }

        self.postMessage({ type: "LOG_MODE_RECEIVED", payload: { logMode } });

        const reader = res.body?.getReader();
        if (!reader) {
            throw new Error("ReadableStream not supported");
        }

        const decoder = new TextDecoder();
        let fullLog = "";

        while (true) {
            const { done, value } = await reader.read();
            if (done) {
                break;
            }
            const chunk = decoder.decode(value);

            // Process the chunk based on logMode
            const processedChunk = processChunk(chunk, logMode);
            fullLog += processedChunk;

            const updateMessage: WorkerPostMessage = {
                type: "LOG_CONTENTS_UPDATED",
                payload: { fullLog },
            };
            self.postMessage(updateMessage);
        }

        const doneMessage: WorkerPostMessage = { type: "LOG_FETCHING_DONE", payload: undefined };
        self.postMessage(doneMessage);
    } catch (error) {
        const errorMessage: WorkerPostMessage = { type: "ERROR", payload: { error: `${error}` } };
        self.postMessage(errorMessage);
    } finally {
        abortController = null;
    }
}

function processChunk(chunk: string, logMode: LogMode): string {
    // Implement chunk processing based on logMode
    // This is a placeholder implementation
    return logMode === "SingleFiltered" ? chunk.toUpperCase() : chunk;
}
