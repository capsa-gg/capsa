"use client";

import logSeverities from "@/types/logSeverities";
import type { LogFilters, LogMode, WorkerCommandMessage, WorkerPostMessage } from "./LogProcessor.types";

let abortController: AbortController | null = null;

self.addEventListener("message", async (event: MessageEvent<WorkerCommandMessage>) => {
    const { type, payload } = event.data;

    switch (type) {
        case "START_FETCHING_LOG": {
            // Abort old request if still loading
            if (abortController) {
                abortController.abort();
            }

            const reqUrl = generateLogUrlWithParams(payload.logUrlBase, payload.filters);
            await fetchLog(reqUrl, payload.jwt);
            break;
        }
        case "STOP_FETCHING_LOG":
            if (abortController) {
                abortController.abort();
            }
            break;
        default:
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

        const message: WorkerPostMessage = { type: "LOG_MODE_RECEIVED", payload: { logMode } };
        self.postMessage(message);

        const reader = res.body?.getReader();
        if (!reader) {
            throw new Error("ReadableStream not supported");
        }

        const decoder = new TextDecoder();

        let fullLog = "";
        let absoluteLineNumbers: string[] = [];

        while (true) {
            const { done, value } = await reader.read();
            if (done) {
                break;
            }
            const chunk = decoder.decode(value);

            switch (logMode) {
                case "SingleUnfiltered": {
                    fullLog += chunk;

                    const updateMessage: WorkerPostMessage = {
                        type: "LOG_CONTENTS_UPDATED",
                        payload: { fullLog, newChunk: chunk, lineNumbers: null },
                    };
                    self.postMessage(updateMessage);

                    break;
                }
                case "SingleFiltered": {
                    const [chunkText, lineNumbers] = await processFilteredChunkSingleLog(chunk);
                    fullLog += chunkText;
                    absoluteLineNumbers = [...absoluteLineNumbers, ...lineNumbers];

                    const updateMessage: WorkerPostMessage = {
                        type: "LOG_CONTENTS_UPDATED",
                        payload: { fullLog, newChunk: chunkText, lineNumbers: absoluteLineNumbers },
                    };
                    self.postMessage(updateMessage);

                    break;
                }
                case "MergedFiltered": {
                    const [chunkText, lineNumbers] = await processFilteredChunkMergedLog(chunk);
                    fullLog += chunkText;
                    absoluteLineNumbers = [...absoluteLineNumbers, ...lineNumbers];

                    const updateMessage: WorkerPostMessage = {
                        type: "LOG_CONTENTS_UPDATED",
                        payload: { fullLog, newChunk: chunkText, lineNumbers: absoluteLineNumbers },
                    };
                    console.log(updateMessage);

                    self.postMessage(updateMessage);

                    break;
                }
                default:
                    throw new Error(`log mode ${logMode} not supported`);
            }
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

async function processFilteredChunkSingleLog(chunk: string): Promise<[string, string[]]> {
    const absoluteLineNumbers = [];
    const cleanedLines = [];

    const regex = /\{(\d+)}(.*)/g;

    // biome-ignore lint/suspicious/noImplicitAnyLet: this is fine here
    let match;

    // biome-ignore lint/suspicious/noAssignInExpressions: this is preferred
    while ((match = regex.exec(chunk)) !== null) {
        absoluteLineNumbers.push(match[1]);
        cleanedLines.push(match[2]);
    }

    return [cleanedLines.join("\n"), absoluteLineNumbers];
}

async function processFilteredChunkMergedLog(chunk: string): Promise<[string, string[]]> {
    const absoluteLineNumbers = [];
    const cleanedLines = [];

    const regex = /\((-+|\w+)\)\{(\d+)}(.*)/g;

    // biome-ignore lint/suspicious/noImplicitAnyLet: this is fine here
    let match;

    // biome-ignore lint/suspicious/noAssignInExpressions: this is preferred
    while ((match = regex.exec(chunk)) !== null) {
        const description = match[1];
        const lineNumber = match[2];
        const spacesCount = 5 - lineNumber.length;

        const newAbsLineNumber = `${description === "--" ? "  " : description}${"&nbsp;".repeat(spacesCount)}${lineNumber}`;

        absoluteLineNumbers.push(newAbsLineNumber);
        cleanedLines.push(match[3]);
    }

    return [cleanedLines.join("\n"), absoluteLineNumbers];
}

function generateLogUrlWithParams(urlBase: string, filters: LogFilters): URL {
    const url = new URL(urlBase);

    const emptyFilters =
        filters.includedSeverities.length === logSeverities.length &&
        filters.includedCategories.length === 0 &&
        filters.excludedCategories.length === 0 &&
        filters.mergeLogs.length === 0;
    if (emptyFilters) {
        console.log("[generateLogUrlWithParams]: Empty log filters, not setting search parameters");
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

    if (filters.mergeLogs.length > 0) {
        url.searchParams.set("merge_logs", filters.mergeLogs.join(","));
    }

    console.log("[generateLogUrlWithParams]: Filters set, search params: ", url.searchParams.toString());

    return url;
}
