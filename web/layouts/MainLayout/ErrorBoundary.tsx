"use client";

import { Box } from "@mui/material";
import { captureException } from "@sentry/browser";
import type React from "react";
import { ErrorBoundary as ErrorBoundaryComponent, type ErrorBoundaryPropsWithFallback } from "react-error-boundary";
import ErrorFallback from "@/layouts/MainLayout/ErrorFallback";

const onError: ErrorBoundaryPropsWithFallback["onError"] = (error, info) => {
    captureException(error, {
        data: info,
    });
};

const ErrorBoundary: React.FC<{ children: React.ReactNode }> = ({ children }) => (
    <ErrorBoundaryComponent
        onError={onError}
        FallbackComponent={ErrorFallback}
        onReset={() => {
            // Reset the state of your app here
            console.log("Error boundary reset");
        }}
    >
        <Box mt={4}>{children}</Box>
    </ErrorBoundaryComponent>
);

export default ErrorBoundary;
