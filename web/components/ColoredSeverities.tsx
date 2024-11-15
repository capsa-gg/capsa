import * as React from "react";
import { Tooltip, Typography } from "@mui/material";
import LogsColors from "@/styles/LogsColors";

const ColoredSeverities: React.FC<{ severities: Record<string, number> }> = ({ severities }) => {
    const orderedSeverities = ["Fatal", "Error", "Warning", "Log", "Display", "Verbose", "VeryVerbose"];

    const severityItems = orderedSeverities.map(s => <Severity key={s} severity={s} count={severities[s] ?? 0} />);

    // eslint-disable-next-line react/jsx-no-useless-fragment
    return <>{severityItems}</>;
};

export default ColoredSeverities;

const Severity: React.FC<{ severity: string; count: number }> = ({ severity, count }) => {
    const color = severityColor(severity);

    return (
        <Tooltip title={severity}>
            <Typography variant="caption" color={color} sx={{ display: "inline-block", minWidth: "16px", px: "4px" }}>
                {count}
            </Typography>
        </Tooltip>
    );
};

export const severityColor = (sev: string): string => {
    switch (sev) {
        case "Fatal":
            return "#000";
        case "Error":
            return LogsColors.Error;
        case "Warning":
            return LogsColors.Warning;
        case "Log":
        case "Display":
            return "#4b4";
        case "Verbose":
            return LogsColors.Verbose;
        case "VeryVerbose":
            return LogsColors.VeryVerbose;
        default:
            return LogsColors.Contents;
    }
};
