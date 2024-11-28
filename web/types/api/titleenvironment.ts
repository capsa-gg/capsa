import { z } from "zod";

export const ListAllTitlesResponseSchema = z.array(z.object({ title: z.string(), createdOn: z.string() }));

export type ListAllTitlesResponse = z.infer<typeof ListAllTitlesResponseSchema>;

export interface AddEnvironmentRequest {
    title: string;
    environment: string;
}

export interface AddTitleRequest {
    title: string;
}

export const AddTitleResponseSchema = z.array(z.object({ title: z.string(), createdOn: z.string() }));

export type AddTitleResponse = z.infer<typeof AddTitleResponseSchema>;
