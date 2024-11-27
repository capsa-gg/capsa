import { Box, Container, Paper, Typography } from "@mui/material";
import type React from "react";
import type { ReactNode } from "react";

interface AuthLayoutProps {
    children: ReactNode;
}

const AuthLayout: React.FC<AuthLayoutProps> = ({ children }) => (
    <Box
        sx={{
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            minHeight: "calc(100vh - 200px)",
        }}
    >
        <Container
            maxWidth="sm"
            sx={{
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                justifyContent: "center",
            }}
        >
            <Typography variant="h4" align="center" gutterBottom mb={3}>
                Please log in to Capsa
            </Typography>
            <Paper
                elevation={3}
                sx={{
                    padding: 4,
                    bgcolor: "white",
                    borderRadius: 2,
                    width: "100%",
                }}
            >
                {children}
            </Paper>
        </Container>
    </Box>
);

export default AuthLayout;
