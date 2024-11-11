import * as React from "react";
import CircularProgress from "@mui/material/CircularProgress";
import Box from "@mui/material/Box";

const Spinner = () => (
    <Box sx={{ display: "flex", p: 3 }}>
        <CircularProgress size={5} />
    </Box>
);

export default Spinner;
