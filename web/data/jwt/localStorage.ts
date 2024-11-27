"use client";

import { jwtStorageKey } from "@/data/jwt/jwtData";
import { type UserJwtData, UserJwtDataSchema } from "@/types/api/auth";

export const getJwtFromLocalStorage = (): UserJwtData | null => {
    if (typeof localStorage === "undefined") {
        return null;
    }

    const jwtData = localStorage.getItem(jwtStorageKey);
    if (!jwtData) {
        return null;
    }

    const { data, success, error } = UserJwtDataSchema.safeParse(JSON.parse(jwtData));
    if (!success) {
        console.error("error parsing jwt", error);
        return null;
    }

    return data;
};

export const setJwtInLocalStorage = (data: UserJwtData) => {
    localStorage.setItem(jwtStorageKey, JSON.stringify(data));
};

export const removeJwtFromLocalStorage = () => {
    localStorage.removeItem(jwtStorageKey);
};
