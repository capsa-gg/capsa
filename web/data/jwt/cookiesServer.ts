"use server";

import { jwtStorageKey } from "@/data/jwt/jwtData";
import type { NextRequest } from "next/server";

export const getJwtCookieFromRequest = async (req: NextRequest): Promise<string | undefined> => {
    return req.cookies.get(jwtStorageKey)?.value;
};

export const deleteJwtCookie = async (req: NextRequest) => {
    req.cookies.delete(jwtStorageKey);
};
