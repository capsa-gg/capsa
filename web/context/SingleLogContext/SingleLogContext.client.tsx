"use client";

import dynamic from "next/dynamic";

const SingleLogContextProviderDynamic = dynamic(() =>
    import("./SingleLogContext").then(mod => mod.SingleLogContextProvider),
);

export default SingleLogContextProviderDynamic;
