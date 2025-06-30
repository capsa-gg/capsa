"use client";

import * as Comlink from "comlink";
import logSeverities from "@/types/logSeverities";
import type { LogFilters, LogMode } from "./LogProcessor.types";

const fetchLog = async (
    logUrlBase: string,
    jwt: string,
    filters: LogFilters,
    onUpdate: (chunk: string, lineNumbers: string[] | null) => void,
    onMode: (mode: LogMode) => void,
): Promise<void> => {
    const reqUrl = generateLogUrlWithParams(logUrlBase, filters);
    const res = await fetch(reqUrl, {
        method: "GET",
        headers: {
            accept: "text/plain",
            authorization: `Bearer ${jwt}`,
        },
    });

    if (!res.ok) throw new Error(`API call responded with ${res.status}`);

    const logMode = res.headers.get("x-capsa-log-mode") as LogMode;
    if (!logMode) throw new Error("Log mode header not set");

    onMode(logMode);

    const reader = res.body?.getReader();
    if (!reader) throw new Error("ReadableStream not supported");

    const decoder = new TextDecoder("utf-8");
    let fullLog = "";
    const absoluteLineNumbers: string[] = [];

    let leftoverBytes = new Uint8Array(0);

    // TODO: This is horrible
    while (true) {
        const { done, value } = await reader.read();
        if (done) break;

        // Combine leftover + new chunk
        const combined = new Uint8Array(leftoverBytes.length + value.length);
        combined.set(leftoverBytes);
        combined.set(value, leftoverBytes.length);

        let start = 0;
        for (let i = 0; i < combined.length; ++i) {
            if (combined[i] === 0x0a /* \n */) {
                const lineBytes = combined.subarray(start, i + 1); // include \n
                const line = decoder.decode(lineBytes, { stream: true });

                switch (logMode) {
                    case "SingleUnfiltered":
                        fullLog += line;
                        break;
                    case "SingleFiltered": {
                        const [text, lines] = await processFilteredChunkSingleLog(line);
                        fullLog += text;
                        absoluteLineNumbers.push(...lines);
                        break;
                    }
                    case "MergedFiltered": {
                        const [text, lines] = await processFilteredChunkMergedLog(line);
                        fullLog += text;
                        absoluteLineNumbers.push(...lines);
                        break;
                    }
                    default: {
                        throw new Error(`Log mode not supported: ${logMode}`);
                    }
                }

                start = i + 1;
            }
        }

        // Save leftover bytes (incomplete line)
        leftoverBytes = combined.subarray(start);

        // Send update
        onUpdate(fullLog, logMode === "SingleUnfiltered" ? null : [...absoluteLineNumbers, " "]);
    }

    // Flush any remaining bytes
    if (leftoverBytes.length > 0) {
        const line = decoder.decode(leftoverBytes, { stream: false });

        switch (logMode) {
            case "SingleUnfiltered":
                fullLog += line;
                break;
            case "SingleFiltered": {
                const [text, lines] = await processFilteredChunkSingleLog(line);
                fullLog += text;
                absoluteLineNumbers.push(...lines);
                break;
            }
            case "MergedFiltered": {
                const [text, lines] = await processFilteredChunkMergedLog(line);
                fullLog += text;
                absoluteLineNumbers.push(...lines);
                break;
            }
            default: {
                throw new Error(`Log mode not supported: ${logMode}`);
            }
        }

        onUpdate(fullLog, logMode === "SingleUnfiltered" ? null : [...absoluteLineNumbers, " "]);
    }
};

Comlink.expose({ fetchLog });

async function processFilteredChunkSingleLog(chunk: string): Promise<[string, string[]]> {
    const absoluteLineNumbers = [];
    const cleanedLines = [];

    const regex = /\{(\d+)}(.*)/g;

    // biome-ignore lint/suspicious/noImplicitAnyLet: this is fine here
    let match;

    // biome-ignore lint/suspicious/noAssignInExpressions: this is preferred
    while ((match = regex.exec(chunk)) !== null) {
        absoluteLineNumbers.push(match[1]);
        cleanedLines.push(`${match[2]}\n`);
    }

    return [cleanedLines.join(""), absoluteLineNumbers];
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
        cleanedLines.push(`${match[3]}\n`);
    }

    return [cleanedLines.join(""), absoluteLineNumbers];
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
