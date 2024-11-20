import * as React from "react";
import { useSingleLogContext } from "@/context/SingleLogContext/SingleLogContext";
import Box from "@mui/material/Box";
import {
    Autocomplete,
    Button,
    ButtonGroup,
    TextField,
    ToggleButton,
    ToggleButtonGroup,
    Typography,
} from "@mui/material";
import { AutocompleteChangeDetails, AutocompleteChangeReason } from "@mui/material/useAutocomplete/useAutocomplete";

export const SeveritiesFilters: React.FC = () => {
    const {
        filters: [filterState, filterDispatch],
    } = useSingleLogContext();

    return (
        <Box>
            <Typography variant="h5" sx={{ mb: 2 }}>
                Severities
            </Typography>
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

export const IncludedCategoriesFilters: React.FC = () => {
    const {
        metadata,
        filters: [filterState, filterDispatch],
    } = useSingleLogContext();

    const options: string[] = Object.keys(metadata?.data?.logData.categoriesCounts ?? {});

    const handleChange = (
        _e: React.SyntheticEvent,
        _vals: string[],
        reason: AutocompleteChangeReason,
        details?: AutocompleteChangeDetails,
    ) => {
        if (reason === "createOption" || reason === "selectOption") {
            // @ts-ignore // always set for these options
            const value = details.option;
            filterDispatch({ type: "INCLUDED_CATEGORY_ADD", category: value });
        }
        if (reason === "removeOption") {
            // @ts-ignore // always set for these options
            const value = details.option;
            filterDispatch({ type: "INCLUDED_CATEGORY_REMOVE", category: value });
        }
        if (reason === "clear") {
            filterDispatch({ type: "INCLUDED_CATEGORY_CLEAR" });
        }
    };

    return (
        <Box>
            <Typography variant="h5" sx={{ mb: 2 }}>
                Included Categories
            </Typography>
            <Autocomplete
                multiple
                freeSolo
                value={filterState.includedCategories}
                onChange={handleChange}
                options={options}
                renderInput={parameters => <TextField {...parameters} label="Enter categories to include" />}
            />
        </Box>
    );
};

export const ExcludedCategoriesFilters: React.FC = () => {
    const {
        metadata,
        filters: [filterState, filterDispatch],
    } = useSingleLogContext();

    const options: string[] = Object.keys(metadata?.data?.logData.categoriesCounts ?? {});

    const handleChange = (
        _e: React.SyntheticEvent,
        _vals: string[],
        reason: AutocompleteChangeReason,
        details?: AutocompleteChangeDetails,
    ) => {
        if (reason === "createOption" || reason === "selectOption") {
            // @ts-ignore // always set for these options
            const value = details.option;
            filterDispatch({ type: "EXCLUDED_CATEGORY_ADD", category: value });
        }
        if (reason === "removeOption") {
            // @ts-ignore // always set for these options
            const value = details.option;
            filterDispatch({ type: "EXCLUDED_CATEGORY_REMOVE", category: value });
        }
        if (reason === "clear") {
            filterDispatch({ type: "EXCLUDED_CATEGORY_CLEAR" });
        }
    };

    return (
        <Box>
            <Typography variant="h5" sx={{ mb: 2 }}>
                Excluded Categories
            </Typography>
            <Autocomplete
                multiple
                freeSolo
                value={filterState.excludedCategories}
                onChange={handleChange}
                options={options}
                renderInput={parameters => <TextField {...parameters} label="Enter categories to exclude" />}
            />
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
