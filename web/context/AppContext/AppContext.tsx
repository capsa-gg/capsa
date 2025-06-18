"use client";

import type React from "react";
import { createContext, useContext, useEffect, useState } from "react";
import type { AppContextData } from "@/context/AppContext/AppContext.types";
import { type Env, getEnv } from "@/data/env";

export const AppContext = createContext<AppContextData>({ env: null });

export const AppContextProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    const [env, setEnv] = useState<Env | null>(null);

    useEffect(() => {
        getEnv().then(env => {
            setEnv(env);
        });
    }, []);

    const context: AppContextData = {
        env,
    };

    return <AppContext.Provider value={context}>{children}</AppContext.Provider>;
};

export const useAppContext = () => {
    const context = useContext(AppContext);
    if (!context) {
        throw new Error("useAppContext must be used within an AppContextProvider");
    }
    return context;
};
