import { z } from "zod";

export const EnvironmentResponseItemSchema = z.object({
    title: z.string(),
    environmentName: z.string(),
    environmentKey: z.string(),
});

export type EnvironmentResponseItem = z.infer<typeof EnvironmentResponseItemSchema>;

export const ListAllEnvironmentsResponseSchema = z.array(EnvironmentResponseItemSchema);

export type ListAllEnvironmentsResponse = z.infer<typeof ListAllEnvironmentsResponseSchema>;
