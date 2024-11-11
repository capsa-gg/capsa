"use client";

// TODO: streaming!
import { getRequestUrl } from "@/api/apibase";
import { getJwtFromLocalStorage } from "@/data/jwt/localStorage";

export const loadSingleLogAsString = async (id: string): Promise<string> => {
    const path = `/user/logs/${id}/log`;
    const reqUrl = await getRequestUrl(path);

    const jwtData = getJwtFromLocalStorage();
    const jwt = jwtData?.token;

    const res = await fetch(reqUrl, {
        method: "GET",
        headers: {
            accept: "application/json",
            authorization: jwt ? `Bearer ${jwt}` : "",
        },
    });

    const logContents = await res.text();

    return logContents;
};
