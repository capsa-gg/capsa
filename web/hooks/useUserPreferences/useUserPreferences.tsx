"use client";

import { useState } from "react";
import { getUserPreferencesFromLocalStorage, setUserPreferencesInLocalStorage } from "./useUserPreferences.storage";
import { defaultUserPreferences, type UserPreferences, type UseUserPreferencesHook } from "./useUserPreferences.types";

// Note: this hook should only be used from useUser, not directly in components
const useUserPreferences: UseUserPreferencesHook = () => {
    const [preferences, setPreferences] = useState(getUserPreferencesState());

    const updatePreferences = (newPreferences: UserPreferences) => {
        setPreferences(newPreferences);
        setUserPreferencesInLocalStorage(newPreferences);
    };

    const setDarkMode = (on: boolean) => {
        const newPreferences = {
            ...preferences,
            darkMode: on,
        };
        updatePreferences(newPreferences);
    };

    return {
        preferences,
        setDarkMode,
    };
};

export default useUserPreferences;

const getUserPreferencesState = (): UserPreferences => {
    const fromLocalStorage = getUserPreferencesFromLocalStorage();
    if (fromLocalStorage) return fromLocalStorage;

    // Not found in local storage, set and return defaults
    setUserPreferencesInLocalStorage(defaultUserPreferences);
    return defaultUserPreferences;
};
