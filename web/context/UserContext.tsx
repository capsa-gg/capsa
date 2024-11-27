// FIXME: write

"use client";

import { getJwtFromLocalStorage } from "@/data/jwt/localStorage";
import type { UserJwtData } from "@/types/api/auth";
import type React from "react";
import { type ReactNode, createContext, useContext, useState } from "react";

export type UserInfo = { isLoggedIn: true; user: UserJwtData } | { isLoggedIn: false; user: null };

interface UserContextType {
    reloadInfoFromLocalStorage: () => void;
    userInfo: UserInfo;
}

const UserContext = createContext<UserContextType | undefined>(undefined);

export const UserProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [userInfo, setUserInfo] = useState<UserInfo>(getUserInfo());

    const reloadInfoFromLocalStorage = () => setUserInfo(getUserInfo());

    return <UserContext.Provider value={{ reloadInfoFromLocalStorage, userInfo }}>{children}</UserContext.Provider>;
};

const useUser = (): UserContextType => {
    const userHook = useContext(UserContext);

    if (!userHook) {
        throw new Error("useUser can only be used in within UserProvider");
    }

    return userHook;
};

export default useUser;

const getUserInfo = (): UserInfo => {
    const userDataLocalStorage = getJwtFromLocalStorage();

    if (userDataLocalStorage) {
        return {
            isLoggedIn: true,
            user: userDataLocalStorage,
        };
    }
    return { isLoggedIn: false, user: null };
};
