"use client";

import { FilterState, FilterStateSchema } from "@/context/SingleLogContext/SingleLogContext.types";

const localStorageKey = "capsa-single-log-filter-state-v1";

export const getSingleLogFilterStateFromLocalStorage = (): FilterState | null => {
    if (typeof localStorage == "undefined") {
        return null;
    }

    const filterStateJson = localStorage.getItem(localStorageKey);
    if (!filterStateJson) {
        return null;
    }

    const { data, success, error } = FilterStateSchema.safeParse(JSON.parse(filterStateJson));
    if (!success) {
        console.error("error parsing single log context state", error);
        return null;
    }

    return data;
};

export const setSingleLogFilterStateInLocalStorage = (data: FilterState) => {
    localStorage.setItem(localStorageKey, JSON.stringify(data));
};

export const removeSingleLogFilterStateFromLocalStorage = () => {
    localStorage.removeItem(localStorageKey);
};
