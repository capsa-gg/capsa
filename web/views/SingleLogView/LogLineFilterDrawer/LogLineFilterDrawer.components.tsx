import * as React from "react";
import { useSingleLogContext } from "@/context/SingleLogContext/SingleLogContext";
import Box from "@mui/material/Box";
import { Button, ButtonGroup, ToggleButton, ToggleButtonGroup, Typography } from "@mui/material";

export const SeveritiesFilters: React.FC = () => {
    const {
        filters: [filterState, filterDispatch],
    } = useSingleLogContext();

    return (
        <Box>
            <Typography variant="h5">Severities</Typography>
            <ToggleButtonGroup
                color="primary"
                aria-label="Severties"
                value={Object.entries(filterState.severities)
                    .filter(([_, enabled]) => enabled)
                    .map(([sev, _]) => sev)}
            >
                {Object.entries(filterState.severities).map(([severity, enabled]) => (
                    <ToggleButton
                        key={severity}
                        value={severity}
                        onClick={() =>
                            filterDispatch({
                                type: "SWITCH_SEVERITY",
                                severity: severity,
                            })
                        }
                    >
                        {severity}
                    </ToggleButton>
                ))}
            </ToggleButtonGroup>
        </Box>
    );
};

export const FilterButtons: React.FC = () => {
    const {
        filters: [_, filterDispatch],
        saveFilters,
    } = useSingleLogContext();

    return (
        <ButtonGroup color="primary" aria-label="Log filter buttons">
            <Button variant="contained" onClick={saveFilters}>
                Save
            </Button>
            <Button variant="outlined" onClick={() => filterDispatch({ type: "RESET_FILTERS" })}>
                Reset
            </Button>
            ,
        </ButtonGroup>
    );
};
