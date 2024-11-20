import logSeverities from "@/types/logSeverities";
import { FilterAction, FilterState } from "@/context/SingleLogContext/SingleLogContext.types";
import { getSingleLogFilterStateFromLocalStorage } from "@/context/SingleLogContext/SingleLogContext.storage";

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
        case "INCLUDED_CATEGORY_ADD": {
            const category = action.category;
            return {
                ...state,
                includedCategories: [...state.includedCategories, category],
            };
        }
        case "INCLUDED_CATEGORY_REMOVE": {
            const category = action.category;
            return {
                ...state,
                includedCategories: state.includedCategories.filter(cat => cat != category),
            };
        }
        case "INCLUDED_CATEGORY_CLEAR": {
            return {
                ...state,
                includedCategories: [],
            };
        }
        case "EXCLUDED_CATEGORY_ADD": {
            const category = action.category;
            return {
                ...state,
                excludedCategories: [...state.excludedCategories, category],
            };
        }
        case "EXCLUDED_CATEGORY_REMOVE": {
            const category = action.category;
            return {
                ...state,
                excludedCategories: state.includedCategories.filter(cat => cat != category),
            };
        }
        case "EXCLUDED_CATEGORY_CLEAR": {
            return {
                ...state,
                excludedCategories: [],
            };
        }
    }

    // In case an action is not supported, return the state to prevent runtime errors
    return state;
};

export const getFilterReducerInitialState = (): FilterState => {
    const stateFromLocalStorage = getSingleLogFilterStateFromLocalStorage();
    return stateFromLocalStorage ?? defaultFilterReducerState;
};

export const defaultFilterReducerState: FilterState = {
    severities: logSeverities.reduce((acc, sev) => ({ ...acc, [sev]: true }), {}),
    includedCategories: [],
    excludedCategories: [],
};
