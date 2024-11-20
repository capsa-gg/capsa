import React from "react";
import { Badge, Typography, Stack, Box } from "@mui/material";
import { SingleLogContextProvider } from "@/context/SingleLogContext/SingleLogContext";
import CopyUrlButton from "@/components/CopyUrlButton";

const SingleLogLayout = async ({ params, children }: SingleLogLayoutProps) => {
    const { id } = await params;
    return (
        <>
            <Stack
                direction="row"
                gap={4}
                sx={{
                    alignItems: "center",
                    mb: 4,
                }}
            >
                <Typography variant="h6">Log ID: {id}</Typography>
                <Box>
                    <Badge color="primary">
                        <CopyUrlButton />
                    </Badge>
                </Box>
            </Stack>
            <SingleLogContextProvider logUUID={id}>{children}</SingleLogContextProvider>
        </>
    );
};

export default SingleLogLayout;

export interface SingleLogLayoutProps {
    children: React.ReactNode;
    params: Promise<{ id: string }>;
}
