import { Typography } from "@mui/material";
import React from "react";

const Search: React.FC<SearchProps> = ({ searchParams }) => {
    const searchParamsUse = React.use(searchParams);
    const { query } = searchParamsUse;

    if (!query) {
        return <Typography variant="h6">Add a search in the top bar</Typography>;
    }

    return <div>Query: {searchParamsUse.query}</div>;
};

export default Search;

export interface SearchProps {
    searchParams: Promise<{ [key: string]: string | string[] | undefined }>;
}
