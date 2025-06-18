import CloseIcon from "@mui/icons-material/Close";
import ContentCopyIcon from "@mui/icons-material/ContentCopy";
import { Alert, AlertTitle, IconButton, Snackbar, Tooltip, Typography } from "@mui/material";
import type React from "react";
import { useState } from "react";
import {
    DEFAULT_DURATION_ERRORS_SECONDS,
    DEFAULT_DURATION_NOTIFICATION_SECONDS,
} from "@/context/NotificationsContext/NotificationsContext";
import type { Notification } from "@/context/NotificationsContext/NotificationsContext.types";

export const NotificationSnackbarItem: React.FC<{ notification: Notification }> = ({ notification }) => {
    const [open, setOpen] = useState(true);
    const [copySuccess, setCopySuccess] = useState(false);

    const handleClose = () => {
        setOpen(false);
    };

    const handleCopy = () => {
        const jsonString = JSON.stringify(notification, null, 2);
        navigator.clipboard.writeText(jsonString).then(() => {
            setCopySuccess(true);
            setTimeout(() => setCopySuccess(false), 2000);
        });
    };

    const AlertBody = () => {
        if (notification.severity === "error") {
            const err = notification.error;
            return (
                <>
                    <AlertTitle>{err.title}</AlertTitle>
                    {err.error}
                    {err.details && (
                        <Typography variant="caption">
                            <br />
                            Details: {capitalizeFirstLetter(err.details)}
                        </Typography>
                    )}
                    {err.additionalData && (
                        <Typography variant="caption">
                            <br />
                            Additional data: {capitalizeFirstLetter(err.additionalData)}
                        </Typography>
                    )}
                    {err.rawError && (
                        <Typography variant="caption">
                            <br />
                            Raw error: {err.rawError}
                        </Typography>
                    )}
                </>
            );
        }

        // Non-error
        return (
            <>
                <AlertTitle>{notification.title}</AlertTitle>
                <Typography variant="caption">
                    <br />
                    {notification.message}
                </Typography>
            </>
        );
    };

    const autoHideMs =
        notification.severity === "error"
            ? DEFAULT_DURATION_ERRORS_SECONDS * 1000 - 500
            : DEFAULT_DURATION_NOTIFICATION_SECONDS * 1000 - 500;

    return (
        <Snackbar
            open={open}
            autoHideDuration={autoHideMs}
            onClose={handleClose}
            anchorOrigin={{ vertical: "bottom", horizontal: "right" }}
        >
            <Alert
                onClose={handleClose}
                severity={notification.severity}
                action={
                    <>
                        <Tooltip title={copySuccess ? "Copied!" : "Copy JSON"}>
                            <IconButton aria-label="copy" color="inherit" size="small" onClick={handleCopy}>
                                <ContentCopyIcon fontSize="inherit" />
                            </IconButton>
                        </Tooltip>
                        <Tooltip title="Hide notification">
                            <IconButton aria-label="copy" color="inherit" size="small" onClick={handleClose}>
                                <CloseIcon fontSize="inherit" />
                            </IconButton>
                        </Tooltip>
                    </>
                }
            >
                <AlertBody />
            </Alert>
        </Snackbar>
    );
};

const capitalizeFirstLetter = (s: string) => {
    if (!s) return s;
    return s.charAt(0).toUpperCase() + s.slice(1);
};
