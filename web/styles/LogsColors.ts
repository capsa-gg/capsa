export enum LogsColorsLightMode {
    Timestamp = "#888888",
    Category = "#555555",
    Contents = "#333333",

    Fatal = "#a80000",
    Error = "#f13434",
    Warning = "#ef7300",
    Info = "#e5be10",
    Verbose = "#527311",
    VeryVerbose = "#1e3877",
}

export enum LogsColorsDarkMode {
    Timestamp = "#555555",
    Category = "#888888",
    Contents = "#cccccc",

    Fatal = "#F44747",
    Error = "#F08080",
    Warning = "#F5A623",
    Info = "#F4E288",
    Verbose = "#6B8E23",
    VeryVerbose = "#3E4E74",
}

export type LogsColors = typeof LogsColorsLightMode | typeof LogsColorsDarkMode;

export const getModeLogColors = (darkMode: boolean): LogsColors =>
    darkMode ? LogsColorsDarkMode : LogsColorsLightMode;
