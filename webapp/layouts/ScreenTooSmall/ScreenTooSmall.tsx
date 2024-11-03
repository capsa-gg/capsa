import React from "react";
import { Container, Typography } from "@mui/material";

const ScreenTooSmall: React.FC = () => (
    <Container maxWidth="sm" sx={{ marginTop: 24 }}>
        <Typography variant="h2" align="center">
            Please view on desktop
        </Typography>
        <Typography align="center">This webapp is only available on desktop screen sizes.</Typography>
    </Container>
);

export default ScreenTooSmall;
