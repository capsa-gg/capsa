"use client";

import type React from "react";
import type { ReactNode } from "react";
import { SWRConfig } from "swr";
import { useNotificationsContext } from "@/context/NotificationsContext/NotificationsContext";
import type { ApplicationError } from "@/types/api/error";

export const SWRProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const notifications = useNotificationsContext();

    const onError = (error: ApplicationError) => notifications.addError(error);

    return <SWRConfig value={{ onError }}>{children}</SWRConfig>;
};
