"use client";

import { Alert, AlertTitle, Button, Collapse, Paper, Typography } from "@mui/material";
import type React from "react";
import { useState } from "react";
import type { FallbackProps } from "react-error-boundary";

const ErrorFallback: React.FC<FallbackProps> = ({ error, resetErrorBoundary }) => {
    const [showStack, setShowStack] = useState(false);

    return (
        <Paper
            elevation={3}
            sx={{
                p: 3,
                maxWidth: 960,
                mx: "auto",
                mt: 4,
            }}
        >
            <Alert severity="error" sx={{ mb: 2 }}>
                <AlertTitle>An error occurred</AlertTitle>
                {error.message}
            </Alert>

            <Button variant="outlined" color="primary" onClick={() => setShowStack(!showStack)} sx={{ mb: 2 }}>
                {showStack ? "Hide" : "Show"} Stack Trace
            </Button>

            <Collapse in={showStack}>
                <Paper variant="outlined" sx={{ p: 2, mb: 2, maxHeight: 200, overflow: "auto" }}>
                    <Typography variant="body2" component="pre" sx={{ whiteSpace: "pre-wrap", wordBreak: "break-all" }}>
                        {error.stack}
                    </Typography>
                </Paper>
            </Collapse>

            <Button variant="contained" color="primary" onClick={resetErrorBoundary}>
                Try Again
            </Button>
        </Paper>
    );
};

export default ErrorFallback;
