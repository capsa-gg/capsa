import { z } from "zod";

export const LogOverviewItemSchema = z.object({
    id: z.string(),
    platform: z.string(),
    logType: z.string(),
    lineCount: z.number(),
    tsFirstLine: z.string(),
    tsLastLine: z.string(),
    categoriesCounts: z.record(z.string(), z.number()),
    severitiesCounts: z.record(z.string(), z.number()),
});

export type LogOverviewItem = z.infer<typeof LogOverviewItemSchema>;

export const LogOverviewResponseSchema = z.array(LogOverviewItemSchema);
export type LogOverviewResponse = z.infer<typeof LogOverviewResponseSchema>;
