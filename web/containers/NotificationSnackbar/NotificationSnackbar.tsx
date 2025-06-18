"use client";

import type React from "react";
import { NotificationSnackbarItem } from "@/containers/NotificationSnackbar/NotificationSnackbar.components";
import { useNotificationsContext } from "@/context/NotificationsContext/NotificationsContext";

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
