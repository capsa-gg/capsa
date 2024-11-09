"use client";

import React, { ReactNode } from "react";
import { SWRConfig } from "swr";
import { useErrors } from "@/context/ErrorContext";

export const SWRProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const errors = useErrors();

    return <SWRConfig value={{ onError: error => errors.addError(error) }}>{children}</SWRConfig>;
};
