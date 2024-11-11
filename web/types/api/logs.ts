import { z } from "zod";

export const LogOverviewItemSchema = z.object({
    id: z.string(),
    lineCount: z.number(),
    logType: z.string(),
    platform: z.string(),
    tsFirstLine: z.nullable(z.string()),
    tsLastLine: z.nullable(z.string()),
});

export type LogOverviewItem = z.infer<typeof LogOverviewItemSchema>;

export const LogOverviewResponseSchema = z.array(LogOverviewItemSchema);
export type LogOverviewResponse = z.infer<typeof LogOverviewResponseSchema>;
