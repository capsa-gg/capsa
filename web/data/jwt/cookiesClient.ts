"use client";

import Cookies from "js-cookie";
import { UserJwtData } from "@/types/api/auth";
import { jwtStorageKey } from "@/data/jwt/jwtData";

export const setJwtCookie = (data: UserJwtData) => {
    Cookies.set(jwtStorageKey, data.token, {
        expires: new Date(data.tokenExpiry),
        secure: true,
        sameSite: "strict",
    });
};

export const getJwtCookie = (): string | undefined => {
    return Cookies.get(jwtStorageKey);
};

export const removeJwtCookie = () => {
    Cookies.remove(jwtStorageKey);
};
