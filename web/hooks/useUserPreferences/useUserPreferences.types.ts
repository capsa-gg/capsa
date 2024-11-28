import { z } from "zod";

export type UseUserPreferencesHook = () => UseUserPreferences;

export interface UseUserPreferences {
    preferences: UserPreferences;
    setDarkMode: (on: boolean) => void;
}

export const UserPreferencesSchema = z.object({
    darkMode: z.boolean(),
});

export type UserPreferences = z.infer<typeof UserPreferencesSchema>;

export const defaultUserPreferences: UserPreferences = {
    darkMode: false,
};
