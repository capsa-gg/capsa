"use client";

import React, { createContext, useContext, useState, useEffect, ReactNode } from "react";
import { ApplicationError } from "@/types/api/error";

interface ErrorContextType {
    errors: ApplicationError[];
    addError: (error: ApplicationError) => void;
    removeError: (title: string) => void; // TODO: We probably want IDs to be used here
}

const ErrorContext = createContext<ErrorContextType | undefined>(undefined);

export const ErrorProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [errors, setErrors] = useState<ApplicationError[]>([]);

    const addError = (error: ApplicationError) => {
        setErrors(prevErrors => [...prevErrors, error]);
        setTimeout(() => {
            setErrors(prevErrors => prevErrors.filter(e => e.title !== error.title));
        }, 10000);
    };

    const removeError = (title: string) => {
        setErrors(prevErrors => prevErrors.filter(e => e.title !== title));
    };

    return <ErrorContext.Provider value={{ errors, addError, removeError }}>{children}</ErrorContext.Provider>;
};

export const useErrors = (): ErrorContextType => {
    const context = useContext(ErrorContext);
    if (!context) {
        throw new Error("useError must be used within an ErrorProvider");
    }
    return context;
};
