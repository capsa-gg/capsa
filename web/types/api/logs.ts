import { z } from "zod";

export type LogOverviewItem = z.infer<typeof LogOverviewItemSchema>;
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

export const LogOverviewResponseSchema = z.array(LogOverviewItemSchema);
export type LogOverviewResponse = z.infer<typeof LogOverviewResponseSchema>;

export type LogAdditionalMetadata = z.infer<typeof LogAdditionalMetadataSchema>;
export const LogAdditionalMetadataSchema = z.object({
    savedOn: z.string(),
    metadata: z.record(z.string(), z.any()),
});

export type LinkedLog = z.infer<typeof LinkedLogSchema>;
export const LinkedLogSchema = z.object({
    linkedLog: z.string(),
    description: z.string(),
});

export type LogMetadata = z.infer<typeof LogMetadataSchema>;
export const LogMetadataSchema = z.object({
    logData: LogOverviewItemSchema,
    additionalMetadata: z.array(LogAdditionalMetadataSchema),
    linkedLogs: z.array(LinkedLogSchema),
});
