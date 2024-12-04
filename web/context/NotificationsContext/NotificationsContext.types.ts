import type { ApplicationError } from "@/types/api/error";

export type Notification =
    | { id?: string; severity: "error"; error: ApplicationError; durationSecondsOverride?: number }
    | { id?: string; severity: "warning"; title: string; message: string; durationSecondsOverride?: number }
    | { id?: string; severity: "info"; title: string; message: string; durationSecondsOverride?: number }
    | { id?: string; severity: "success"; title: string; message: string; durationSecondsOverride?: number };

export interface NotificationsContextType {
    notifications: Notification[];
    addError: (error: ApplicationError) => void;
    addNotification: (notification: Notification) => void;
    removeNotification: (id: string) => void;
}
