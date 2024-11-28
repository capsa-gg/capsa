import Box from "@mui/material/Box";
import CircularProgress from "@mui/material/CircularProgress";

const Spinner = () => (
    <Box sx={{ display: "flex", p: 3 }}>
        <CircularProgress size={64} />
    </Box>
);

export default Spinner;
