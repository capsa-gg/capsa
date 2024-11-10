"use client";

import React, { createContext, useContext, useEffect, useState } from "react";
import { Env, getEnv } from "@/data/env";
import { AppContextData } from "@/context/AppContext/AppContext.types";

export const AppContext = createContext<AppContextData>({ env: null });

export const AppContextProvider = ({ children }: React.FC<{ children: React.ReactNode }>) => {
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
