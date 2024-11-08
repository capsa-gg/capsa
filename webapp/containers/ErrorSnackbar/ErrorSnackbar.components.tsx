import { Alert, AlertTitle, Snackbar, Typography } from "@mui/material";
import React, { useEffect, useState } from "react";
import { ApplicationError } from "@/types/api/error";

export const ErrorSnackbarItem: React.FC<{ err: ApplicationError }> = ({ err }) => {
    const [open, setOpen] = useState(true);

    const handleClose = () => {
        setOpen(false);
    };

    return (
        <Snackbar
            open={open}
            autoHideDuration={8000}
            onClose={handleClose}
            anchorOrigin={{ vertical: "bottom", horizontal: "right" }}
        >
            <Alert onClose={handleClose} severity={err.severity ?? "info"}>
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
            </Alert>
        </Snackbar>
    );
};

const capitalizeFirstLetter = (s: string) => {
    if (!s) return s;
    return s.charAt(0).toUpperCase() + s.slice(1);
};
