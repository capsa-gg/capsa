"use client";

import { useErrors } from "@/context/ErrorContext";
import type React from "react";
import type { ReactNode } from "react";
import { SWRConfig } from "swr";

export const SWRProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const errors = useErrors();

    return <SWRConfig value={{ onError: error => errors.addError(error) }}>{children}</SWRConfig>;
};
