import { Button, ButtonGroup, Checkbox, Divider, Drawer, FormControlLabel, FormGroup, Typography } from "@mui/material";
import Box from "@mui/material/Box";
import { Stack } from "@mui/system";
import type * as React from "react";
import { useSingleLogContext } from "@/context/SingleLogContext/SingleLogContext";

const MergeLogsDrawer: React.FC = () => {
    const {
        mergeDrawerState: [drawerOpen, setDrawerOpen],
        filters: [filterState, filterDispatch],
        metadata: { data: metadata },
        saveFilters,
    } = useSingleLogContext();

    if (!metadata) return null;

    // TODO: This should be a proper table
    // TODO: Show the labels for the merged logs, see what G1, E1, C1 etc. are
    const MergedLogs = () =>
        metadata.linkedLogs.map(ll => (
            <Stack direction="row" justifyContent="space-between" alignItems="center" gap={6} key={ll.linkedLog}>
                <Box>
                    <Typography>
                        <b>{ll.linkedLog}</b> ({ll.description})
                    </Typography>
                </Box>
                <Box>
                    <FormGroup>
                        <FormControlLabel
                            control={
                                <Checkbox
                                    checked={filterState.mergedLogs.includes(ll.linkedLog)}
                                    onChange={() => filterDispatch({ type: "MERGED_LOGS_SWITCH", log: ll.linkedLog })}
                                    inputProps={{ "aria-label": "controlled" }}
                                />
                            }
                            label="Include"
                        />
                    </FormGroup>
                </Box>
            </Stack>
        ));

    return (
        <Drawer anchor="right" open={drawerOpen} onClose={() => setDrawerOpen(false)}>
            <Box sx={{ width: 800, p: 5 }} role="presentation">
                <Typography variant="h3">Merged logs</Typography>
                <Divider sx={{ mt: 2, mb: 5 }} />
                <Stack direction="column" justifyContent="space-between" sx={{ height: "100%" }}>
                    <Stack direction="column" gap={1}>
                        <MergedLogs />
                    </Stack>
                    <ButtonGroup color="primary" aria-label="Log filter buttons">
                        <Button variant="contained" onClick={saveFilters}>
                            Save
                        </Button>
                        <Button variant="outlined" onClick={() => filterDispatch({ type: "MERGED_LOGS_CLEAR" })}>
                            Clear
                        </Button>
                    </ButtonGroup>
                </Stack>
            </Box>
        </Drawer>
    );
};

export default MergeLogsDrawer;
