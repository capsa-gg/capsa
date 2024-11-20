import logSeverities from "@/types/logSeverities";
import { FilterAction, FilterState } from "@/context/SingleLogContext/SingleLogContext.types";

export const filterReducer = (state: FilterState, action: FilterAction) => {
    switch (action.type) {
        case "RESET_FILTERS": {
            return defaultFilterReducerState;
        }
        case "SWITCH_SEVERITY": {
            const sev = action.severity;
            const hasSev = state.severities[sev];
            if (!hasSev) return state;

            return {
                ...state,
                severities: {
                    ...state.severities,
                    [sev]: !state.severities[sev],
                },
            };
        }
    }
};

export const getFilterReducerInitialState = (): FilterState => {
    // TODO: LocalStorage persistence
    return defaultFilterReducerState;
};

const defaultFilterReducerState: FilterState = {
    severities: logSeverities.reduce((acc, sev) => ({ ...acc, [sev]: true }), {}),
};
