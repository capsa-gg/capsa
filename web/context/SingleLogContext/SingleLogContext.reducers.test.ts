import { filterReducer, defaultFilterReducerState } from "./SingleLogContext.reducers";
import { FilterAction, FilterState } from "@/context/SingleLogContext/SingleLogContext.types";
import logSeverities from "@/types/logSeverities";

describe("filterReducer", () => {
    let initialState: FilterState;

    beforeEach(() => {
        initialState = defaultFilterReducerState;
    });

    it("should reset to default state on RESET_FILTERS", () => {
        const actionSwitch: FilterAction = { type: "SWITCH_SEVERITY", severity: logSeverities[0] };
        const newState = filterReducer(initialState, actionSwitch);
        expect(newState).not.toEqual(initialState);

        const action: FilterAction = { type: "RESET_FILTERS" };
        const resetState = filterReducer(newState, action);
        expect(resetState).toEqual(initialState);
    });

    it("should toggle severity on SWITCH_SEVERITY", () => {
        const action: FilterAction = { type: "SWITCH_SEVERITY", severity: logSeverities[0] };
        const newState = filterReducer(initialState, action);

        expect(newState.severities[logSeverities[0]]).toBe(false);
        expect(newState.severities).not.toEqual(initialState.severities);
    });

    it("should add a category to includedCategories on INCLUDED_CATEGORY_ADD", () => {
        const action: FilterAction = { type: "INCLUDED_CATEGORY_ADD", category: "LogCategoryTest" };
        const newState = filterReducer(initialState, action);

        expect(newState.includedCategories).toContain("LogCategoryTest");
        expect(newState.includedCategories.length).toBe(1);
    });

    it("should remove a category from includedCategories on INCLUDED_CATEGORY_REMOVE", () => {
        const stateWithCategories: FilterState = {
            ...initialState,
            includedCategories: ["LogCategoryTest"],
        };
        const action: FilterAction = { type: "INCLUDED_CATEGORY_REMOVE", category: "LogCategoryTest" };
        const newState = filterReducer(stateWithCategories, action);

        expect(newState.includedCategories).not.toContain("LogCategoryTest");
        expect(newState.includedCategories.length).toBe(0);
    });

    it("should clear categories from includedCategories on INCLUDED_CATEGORY_CLEAR", () => {
        const stateWithIncluded: FilterState = {
            ...initialState,
            includedCategories: ["LogCategoryTest", "LogCategoryTestTwo"],
        };
        const action: FilterAction = { type: "INCLUDED_CATEGORY_CLEAR" };
        const newState = filterReducer(stateWithIncluded, action);
        expect(newState.includedCategories.length).toBe(0);
    });

    it("should add a category to excludedCategories on EXCLUDED_CATEGORY_ADD", () => {
        const action: FilterAction = { type: "EXCLUDED_CATEGORY_ADD", category: "LogCategoryTest" };
        const newState = filterReducer(initialState, action);

        expect(newState.excludedCategories).toContain("LogCategoryTest");
        expect(newState.excludedCategories.length).toBe(1);
    });

    it("should remove a category from excludedCategories on EXCLUDED_CATEGORY_REMOVE", () => {
        const stateWithExcluded: FilterState = {
            ...initialState,
            excludedCategories: ["LogCategoryTest"],
        };
        const action: FilterAction = { type: "EXCLUDED_CATEGORY_REMOVE", category: "LogCategoryTest" };
        const newState = filterReducer(stateWithExcluded, action);

        expect(newState.excludedCategories).not.toContain("LogCategoryTest");
        expect(newState.excludedCategories.length).toBe(0);
    });

    it("should clear categories from excludedCategories on EXCLUDED_CATEGORY_CLEAR", () => {
        const stateWithExcluded: FilterState = {
            ...initialState,
            excludedCategories: ["LogCategoryTest", "LogCategoryTestTwo"],
        };
        const action: FilterAction = { type: "EXCLUDED_CATEGORY_CLEAR" };
        const newState = filterReducer(stateWithExcluded, action);
        expect(newState.excludedCategories.length).toBe(0);
    });
});
