import { getSingleLogFilterStateFromLocalStorage } from "@/context/SingleLogContext/SingleLogContext.storage";
import type { FilterAction, FilterState } from "@/context/SingleLogContext/SingleLogContext.types";
import logSeverities from "@/types/logSeverities";

export const filterReducer = (state: FilterState, action: FilterAction) => {
    switch (action.type) {
        case "RESET_FILTERS": {
            return defaultFilterReducerState;
        }
        case "SWITCH_SEVERITY": {
            const sev = action.severity;

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
                includedCategories: state.includedCategories.filter(cat => cat !== category),
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
                excludedCategories: state.includedCategories.filter(cat => cat !== category),
            };
        }
        case "EXCLUDED_CATEGORY_CLEAR": {
            return {
                ...state,
                excludedCategories: [],
            };
        }
        default:
            console.error("[SingleLogContext.reducers]: unknown action", action);
    }

    // In case an action is not supported, return the state to prevent runtime errors
    return state;
};

export const getFilterReducerLocalInitialState = (): FilterState => {
    const stateFromLocalStorage = getSingleLogFilterStateFromLocalStorage();
    return stateFromLocalStorage ?? defaultFilterReducerState;
};

export const defaultFilterReducerState: FilterState = {
    // biome-ignore lint/performance/noAccumulatingSpread: ok here
    severities: logSeverities.reduce((acc, sev) => ({ ...acc, [sev]: true }), {}),
    includedCategories: [],
    excludedCategories: [],
};
