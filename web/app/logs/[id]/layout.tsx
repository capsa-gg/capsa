import React from "react";
import { Typography } from "@mui/material";

const SingleLogLayout = async ({ params, children }: SingleLogLayoutProps) => {
    const { id } = await params;
    return (
        <>
            <Typography variant="h6" mb={4}>
                Log ID: {id}
            </Typography>
            {children}
        </>
    );
};

export default SingleLogLayout;

export interface SingleLogLayoutProps {
    children: React.ReactNode;
    params: Promise<{ id: string }>;
}
