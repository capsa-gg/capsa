"use client";

import { NotificationSnackbarItem } from "@/containers/NotificationSnackbar/NotificationSnackbar.components";
import { useNotificationsContext } from "@/context/NotificationsContext/NotificationsContext";
import type React from "react";

const NotificationSnackbar: React.FC = () => {
    const { notifications } = useNotificationsContext();

    return (
        <>
            {notifications.map(n => (
                <NotificationSnackbarItem key={n.id} notification={n} />
            ))}
        </>
    );
};

export default NotificationSnackbar;
