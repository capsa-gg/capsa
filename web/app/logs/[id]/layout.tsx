import React from "react";
import { Typography } from "@mui/material";
import { SingleLogContextProvider } from "@/context/SingleLogContext/SingleLogContext";

const SingleLogLayout = async ({ params, children }: SingleLogLayoutProps) => {
    const { id } = await params;
    return (
        <>
            <Typography variant="h6" mb={4}>
                Log ID: {id}
            </Typography>
            <SingleLogContextProvider logUUID={id}>{children}</SingleLogContextProvider>
        </>
    );
};

export default SingleLogLayout;

export interface SingleLogLayoutProps {
    children: React.ReactNode;
    params: Promise<{ id: string }>;
}
