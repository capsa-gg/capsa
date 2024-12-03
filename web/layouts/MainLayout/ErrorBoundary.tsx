"use client";

import ErrorFallback from "@/layouts/MainLayout/ErrorFallback";
import { Box } from "@mui/material";
import type React from "react";
import { ErrorBoundary as ErrorBoundaryComponent } from "react-error-boundary";

const ErrorBoundary: React.FC<{ children: React.ReactNode }> = ({ children }) => (
    <ErrorBoundaryComponent
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
