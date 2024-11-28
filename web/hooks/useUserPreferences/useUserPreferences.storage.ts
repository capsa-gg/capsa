"use client";

import { type UserPreferences, UserPreferencesSchema } from "./useUserPreferences.types";

const localStorageKey = "capsa-user-preferences-v1";

export const getUserPreferencesFromLocalStorage = (): UserPreferences | null => {
    if (typeof localStorage === "undefined") {
        return null;
    }

    const userPreferences = localStorage.getItem(localStorageKey);
    if (!userPreferences) {
        return null;
    }

    const { data, success, error } = UserPreferencesSchema.safeParse(JSON.parse(userPreferences));
    if (!success) {
        console.error("error parsing user preferences from local storage", error);
        removeUserPreferencesFromLocalStorage();
        return null;
    }

    return data;
};

export const setUserPreferencesInLocalStorage = (data: UserPreferences) => {
    if (typeof localStorage === "undefined") {
        return;
    }

    localStorage.setItem(localStorageKey, JSON.stringify(data));
};

export const removeUserPreferencesFromLocalStorage = () => {
    if (typeof localStorage === "undefined") {
        return;
    }

    localStorage.removeItem(localStorageKey);
};
