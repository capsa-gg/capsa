"use client";

import type { ApplicationError } from "@/types/api/error";
import type React from "react";
import { type ReactNode, createContext, useContext, useState } from "react";
import type { Notification, NotificationsContextType } from "./NotificationsContext.types";

export const DEFAULT_DURATION_NOTIFICATION_SECONDS = 5;
export const DEFAULT_DURATION_ERRORS_SECONDS = 9;

const NotificationsContext = createContext<NotificationsContextType | null>(null);

export const NotificationsContextProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [notifications, setNotifications] = useState<Notification[]>([]);

    const addError = (error: ApplicationError) => {
        const id = generateNotificationID();
        const notification: Notification = {
            id,
            severity: "error",
            error,
        };

        addNotification(notification);

        const timeoutMs = getTimeoutDuration(notification);
        if (timeoutMs > 0) {
            setTimeout(() => {
                removeNotification(id);
            }, timeoutMs);
        }
    };

    const addNotification = (notification: Notification) => {
        if (!notification.id) notification.id = generateNotificationID();
        setNotifications(prevNotifications => [...prevNotifications, notification]);
    };

    const removeNotification = (id: string) => {
        setNotifications(prevNotifications => prevNotifications.filter(e => e.id !== id));
    };

    const value: NotificationsContextType = {
        notifications,
        addNotification,
        addError,
        removeNotification,
    };

    return <NotificationsContext.Provider value={value}>{children}</NotificationsContext.Provider>;
};

export const useNotificationsContext = (): NotificationsContextType => {
    const context = useContext(NotificationsContext);
    if (!context) {
        throw new Error("useNotificationsContext must be used within a NotificationsContextProvider");
    }
    return context;
};

const generateNotificationID = (): string => `${new Date().getTime()}`;

const getTimeoutDuration = (notification: Notification): number => {
    if (notification.durationSecondsOverride !== undefined) return notification.durationSecondsOverride * 1000;

    if (notification.severity === "error") {
        return DEFAULT_DURATION_ERRORS_SECONDS * 1000;
    }

    return DEFAULT_DURATION_NOTIFICATION_SECONDS * 1000;
};
