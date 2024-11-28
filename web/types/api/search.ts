import { z } from "zod";

export const SearchResultSchema = z.object({
    type: z.string(),
    match: z.string(),
    description: z.string(),
    details: z.string(),
});

export type SearchResultItem = z.infer<typeof SearchResultSchema>;

export const SearchResultsScheme = z.object({
    hasMore: z.boolean(),
    results: z.array(SearchResultSchema),
});

export type SearchResults = z.infer<typeof SearchResultsScheme>;
