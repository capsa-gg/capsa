import { Badge, Box, Stack, Typography } from "@mui/material";
import type React from "react";
import CopyUrlButton from "@/components/CopyUrlButton";
import { SingleLogContextProvider } from "@/context/SingleLogContext/SingleLogContext";

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
