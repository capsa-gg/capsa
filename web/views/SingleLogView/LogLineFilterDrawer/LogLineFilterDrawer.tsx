import * as React from "react";
import { Divider, Drawer, Typography } from "@mui/material";
import Box from "@mui/material/Box";
import { useSingleLogContext } from "@/context/SingleLogContext/SingleLogContext";
import { Stack } from "@mui/system";
import {
    FilterButtons,
    SeveritiesFilters,
} from "@/views/SingleLogView/LogLineFilterDrawer/LogLineFilterDrawer.components";

const LogLineFilterDrawer: React.FC = () => {
    const {
        drawerState: [drawerOpen, setDrawerOpen],
        filters: [filterState, filterDispatch],
    } = useSingleLogContext();

    return (
        <Drawer anchor="right" open={drawerOpen} onClose={() => setDrawerOpen(false)}>
            <Box sx={{ width: 800, p: 5 }} role="presentation">
                <Typography variant="h3">Log line filters</Typography>
                <Divider sx={{ mt: 2, mb: 5 }} />
                <Stack direction="column" justifyContent="space-between" sx={{ height: "100%" }}>
                    <Stack direction="row">
                        <SeveritiesFilters />
                    </Stack>
                    <FilterButtons />
                </Stack>
            </Box>
        </Drawer>
    );
};

export default LogLineFilterDrawer;
