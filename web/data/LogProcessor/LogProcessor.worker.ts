"use client";

import type { LogFilters, LogMode, WorkerCommandMessage, WorkerPostMessage } from "./LogProcessor.types";
import logSeverities from "@/types/logSeverities";

let abortController: AbortController | null = null;

self.addEventListener("message", async (event: MessageEvent<WorkerCommandMessage>) => {
    const { type, payload } = event.data;

    switch (type) {
        case "START_FETCHING_LOG":
            const reqUrl = generateLogUrlWithParams(payload.logUrlBase, payload.filters);
            await fetchLog(reqUrl, payload.jwt);
            break;
        case "STOP_FETCHING_LOG":
            if (abortController) {
                abortController.abort();
            }
            break;
    }
});

async function fetchLog(reqUrl: URL, jwt: string): Promise<void> {
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

function generateLogUrlWithParams(urlBase: string, filters: LogFilters): URL {
    const url = new URL(urlBase);

    const emptyFilters =
        filters.includedSeverities.length === logSeverities.length &&
        filters.includedCategories.length === 0 &&
        filters.excludedCategories.length === 0;
    if (emptyFilters) {
        console.log("Empty log filters, not setting search parameters");
        return url;
    }

    if (filters.includedSeverities.length > 0) {
        url.searchParams.set("included_severities", filters.includedSeverities.join(","));
    }

    if (filters.includedCategories.length > 0) {
        url.searchParams.set("included_categories", filters.includedCategories.join(","));
    }

    if (filters.excludedCategories.length > 0) {
        url.searchParams.set("excluded_categories", filters.excludedCategories.join(","));
    }

    console.log("Filters set, search params: ", url.searchParams.toString());

    return url;
}
