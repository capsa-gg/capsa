"use client";

import { UserJwtData } from "@/types/api/auth";
import { getJwtFromLocalStorage } from "@/data/jwt/localStorage";

const useUserInfo: UseUserInfo = () => {
    const userData = getJwtFromLocalStorage();

    if (userData) {
        return { isLoggedIn: true, user: userData };
    }

    return { isLoggedIn: false, user: null };
};

export type UseUserInfo = () => { isLoggedIn: true; user: UserJwtData } | { isLoggedIn: false; user: null };

export default useUserInfo;
