import { Tooltip, Typography } from "@mui/material";
import type * as React from "react";
import useUser from "@/context/UserContext";
import { getModeLogColors, type LogsColors } from "@/styles/LogsColors";
import logSeverities from "@/types/logSeverities";

const ColoredSeverities: React.FC<{ severities: Record<string, number> }> = ({ severities }) => {
    const severityItems = logSeverities.map(s => <Severity key={s} severity={s} count={severities[s] ?? 0} />);

    return <>{severityItems}</>;
};

export default ColoredSeverities;

const Severity: React.FC<{ severity: string; count: number }> = ({ severity, count }) => {
    const {
        preferences: {
            preferences: { darkMode },
        },
    } = useUser();

    const colorScheme = getModeLogColors(darkMode);
    const color = severityColor(colorScheme, severity);

    return (
        <Tooltip title={severity}>
            <Typography variant="caption" color={color} sx={{ display: "inline-block", minWidth: "16px", px: "4px" }}>
                {count}
            </Typography>
        </Tooltip>
    );
};

export const severityColor = (colors: LogsColors, sev: string): string => {
    switch (sev) {
        case "Fatal":
            return colors.Fatal;
        case "Error":
            return colors.Error;
        case "Warning":
            return colors.Warning;
        case "Log":
        case "Display":
            return colors.Info;
        case "Verbose":
            return colors.Verbose;
        case "VeryVerbose":
            return colors.VeryVerbose;
        default:
            return colors.Contents;
    }
};
