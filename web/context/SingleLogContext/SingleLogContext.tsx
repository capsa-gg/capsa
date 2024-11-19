"use client";

import React, { createContext, useContext } from "react";
import { SingleLogContextProviderProps, SingleLogContextData } from "@/context/SingleLogContext/SingleLogContext.types";
import { useGetSingleLogMetadata } from "@/api/hooks";

//@ts-ignore // This is fine, we are checking with the use hook
export const SingleLogContext = createContext<SingleLogContextData>(undefined);

export const SingleLogContextProvider: React.FC<SingleLogContextProviderProps> = ({ logUUID, children }) => {
    const metadata = useGetSingleLogMetadata(logUUID);

    const context: SingleLogContextData = {
        metadata: metadata,
    };

    return <SingleLogContext.Provider value={context}>{children}</SingleLogContext.Provider>;
};

export const useSingleLogContext = () => {
    const context = useContext(SingleLogContext);
    if (!context) {
        throw new Error("useSingleLogContext must be used within an SingleLogContextProvider");
    }
    return context;
};
