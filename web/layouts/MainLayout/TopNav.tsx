"use client";

import { Box, Button, TextField, Toolbar } from "@mui/material";
import AppBar from "@mui/material/AppBar";
import { useRouter } from "next/navigation";
import { type ChangeEvent, type FormEvent, useState } from "react";
import useUser from "@/context/UserContext";
import { removeJwtCookie } from "@/data/jwt/cookiesClient";
import { removeJwtFromLocalStorage } from "@/data/jwt/localStorage";

export const topNavSidePaddingPx = 20;

const TopNav = () => {
    const router = useRouter();
    const {
        userInfo: { isLoggedIn },
    } = useUser();
    const signOut = () => {
        removeJwtCookie();
        removeJwtFromLocalStorage();
        router.refresh();
    };

    return (
        <AppBar
            position="static"
            sx={{
                paddingRight: "40px",
                background: "transparent",
                boxShadow: "none",
            }}
        >
            <Toolbar sx={{ justifyContent: "space-between" }} disableGutters>
                <Box>
                    <SearchField />
                </Box>
                <Box>
                    <Button variant="outlined" size="small" disabled={!isLoggedIn} onClick={signOut}>
                        Sign out
                    </Button>
                </Box>
            </Toolbar>
        </AppBar>
    );
};

export default TopNav;

const SearchField = () => {
    const [query, setQuery] = useState("");
    const router = useRouter();
    const handleInputChange = (event: ChangeEvent<HTMLInputElement>) => {
        setQuery(event.target.value);
    };
    const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        if (query) {
            router.push(`/search?query=${encodeURIComponent(query)}`);
        }
    };
    return (
        <Box component="form" onSubmit={handleSubmit} sx={{ display: "flex", alignItems: "center" }}>
            <TextField
                label="Search"
                variant="outlined"
                size="small"
                value={query}
                onChange={handleInputChange}
                sx={{ marginRight: 1 }}
            />
        </Box>
    );
};
