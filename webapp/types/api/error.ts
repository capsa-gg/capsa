import { z } from "zod";

export const ErrorResponseSchema = z.object({
    // Error message
    error: z.string(),

    // Used mostly in development
    details: z.string().optional(),

    // Only used in development
    raw_error: z.string().optional(),
});
export type ErrorResponse = z.infer<typeof ErrorResponseSchema>;

// Application extended schema
export const ApiErrorSchema = ErrorResponseSchema.extend({
    title: z.string(),
    severity: z.enum(["warning", "error"]),
    additionalData: z.string().optional(),
});
type ApiError = z.infer<typeof ApiErrorSchema>;

z.never();

export default ApiError;

export class ApplicationError extends Error {
    public title: string;
    public severity: "error" | "warning";
    public error: string;
    public details?: string;
    public additionalData?: string;
    public rawError?: string;

    public constructor(err: ApiError) {
        super(err.title);
        this.title = err.title;
        this.severity = err.severity;
        this.error = err.error;
        this.details = err.details;
        this.additionalData = err.additionalData;
        this.rawError = err.raw_error;
    }
}
