"use client";

import ListAllEnvironments from "@/views/ListAllEnvironments";
import { Typography } from "@mui/material";

const Home = () => {
    return (
        <>
            <Typography variant="h4" mb={2}>
                Available environments
            </Typography>
            <ListAllEnvironments />
        </>
    );
};

export default Home;
