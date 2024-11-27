"use client";

import { type FilterState, FilterStateSchema } from "@/context/SingleLogContext/SingleLogContext.types";

const localStorageKey = "capsa-single-log-filter-state-v1";

export const getSingleLogFilterStateFromLocalStorage = (): FilterState | null => {
    if (typeof localStorage === "undefined") {
        return null;
    }

    const filterStateJson = localStorage.getItem(localStorageKey);
    if (!filterStateJson) {
        return null;
    }

    try {
        const { data, success, error } = FilterStateSchema.safeParse(JSON.parse(filterStateJson));
        if (!success) {
            console.error("error parsing single log context state", error);
            removeSingleLogFilterStateFromLocalStorage();
            return null;
        }
        return data;
    } catch (e) {
        console.error("error parsing single log context state", e);
        removeSingleLogFilterStateFromLocalStorage();
        return null;
    }
};

export const setSingleLogFilterStateInLocalStorage = (data: FilterState) => {
    localStorage.setItem(localStorageKey, JSON.stringify(data));
};

export const removeSingleLogFilterStateFromLocalStorage = () => {
    localStorage.removeItem(localStorageKey);
};
