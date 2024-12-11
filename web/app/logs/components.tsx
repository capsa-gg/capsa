import type { EnvironmentResponseItem, ListAllEnvironmentsResponse } from "@/types/api/environments";
import { Box, FormControl, InputLabel, Link, MenuItem, Select, type SelectChangeEvent, TextField } from "@mui/material";
import { useRouter, useSearchParams } from "next/navigation";
import type React from "react";
import { type ChangeEvent, type FormEvent, useState } from "react";

export const LogLink: React.FC<{ id: string }> = ({ id }) => {
    const router = useRouter();

    return (
        <Link sx={{ cursor: "pointer" }} onClick={() => router.push(`/logs/${id}`)}>
            {id}
        </Link>
    );
};

export const PlatformInput = () => {
    const searchParams = useSearchParams();
    const router = useRouter();
    const [platform, setPlatform] = useState("");
    const handleInputChange = (event: ChangeEvent<HTMLInputElement>) => {
        setPlatform(event.target.value);
    };
    const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        const params = new URLSearchParams(searchParams);
        params.set("platform", platform);
        router.push(`/logs?${params.toString()}`);
    };
    return (
        <Box component="form" onSubmit={handleSubmit}>
            <TextField label="Platform" variant="outlined" value={platform} onChange={handleInputChange} />
        </Box>
    );
};

export const LogTypeSelector: React.FC = () => {
    const router = useRouter();
    const searchParams = useSearchParams();
    const currentEnv = searchParams.get("type") || "all";

    const handleChange = (event: SelectChangeEvent) => {
        const newEnv = event.target.value;
        const params = new URLSearchParams(searchParams);
        params.set("type", newEnv);
        router.push(`/logs?${params.toString()}`);
    };

    return (
        <FormControl sx={{ minWidth: 100 }}>
            <InputLabel>Log Type</InputLabel>
            <Select value={currentEnv} label="Log Type" onChange={handleChange}>
                <MenuItem value="all">All</MenuItem>
                <MenuItem value="Client">Client</MenuItem>
                <MenuItem value="Game">Game</MenuItem>
                <MenuItem value="Editor">Editor</MenuItem>
                <MenuItem value="Server">Server</MenuItem>
            </Select>
        </FormControl>
    );
};

export const EnvironmentSelector: React.FC<{ envs?: ListAllEnvironmentsResponse }> = ({ envs }) => {
    const router = useRouter();
    const searchParams = useSearchParams();
    const currentEnv = searchParams?.get("env") || "all";

    const handleChange = (event: SelectChangeEvent) => {
        const newEnv = event.target.value;
        const params = new URLSearchParams(searchParams);
        params.set("env", newEnv);
        router.push(`/logs?${params.toString()}`);
    };

    if (!envs) return null;

    return (
        <FormControl sx={{ minWidth: 200 }}>
            <InputLabel>Environment</InputLabel>
            <Select value={currentEnv} label="Environment" onChange={handleChange}>
                <MenuItem value="all">All</MenuItem>
                {(JSON.parse(JSON.stringify(envs)) as ListAllEnvironmentsResponse)
                    .sort((a, b) => (envToDisplayString(a) > envToDisplayString(b) ? 1 : -1))
                    .map(env => (
                        <MenuItem key={env.environmentKey} value={env.environmentKey}>
                            {envToDisplayString(env)}
                        </MenuItem>
                    ))}
            </Select>
        </FormControl>
    );
};

const envToDisplayString = (env: EnvironmentResponseItem): string => `${env.title} / ${env.environmentName}`;
