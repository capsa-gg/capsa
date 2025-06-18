"use client";

import Cookies from "js-cookie";
import { jwtStorageKey } from "@/data/jwt/jwtData";
import type { UserJwtData } from "@/types/api/auth";

export const setJwtCookie = (data: UserJwtData) => {
    Cookies.set(jwtStorageKey, data.token, {
        expires: new Date(data.tokenExpiry),
        secure: true,
        sameSite: "Strict",
        httpOnly: false,
    });
};

export const removeJwtCookie = () => {
    Cookies.remove(jwtStorageKey);
};
