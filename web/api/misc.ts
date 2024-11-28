"use client";

import { type ApiResponse, call } from "@/api/apibase";
import { type SearchResults, SearchResultsScheme } from "@/types/api/search";

// biome-ignore lint/complexity/noStaticOnlyClass: desired here
export class MiscApiCalls {
    // Due to string params a separate helper
    public static async Search(search: string): Promise<ApiResponse<SearchResults>> {
        const endpoint = `/user/search?search=${encodeURIComponent(search)}`;
        return await call("GET", endpoint, SearchResultsScheme, true);
    }
}
