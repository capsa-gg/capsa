import useUser from "@/context/UserContext";
import { Editor } from "@monaco-editor/react";
import type * as React from "react";
import { initializeMonaco, monacoOptions, themes } from "./LogViewer.settings";

const LogViewer: React.FC<LogViewerProps> = ({ data, absoluteLineNumbers }) => {
    const {
        preferences: {
            preferences: { darkMode },
        },
    } = useUser();

    const theme = darkMode ? themes.darkMode : themes.lightMode;

    return (
        <Editor
            value={data}
            language="capsalog"
            theme={theme}
            onMount={initializeMonaco(theme)}
            options={{
                ...monacoOptions,
                lineNumbers: absoluteLineNumbers ? (i: number) => `${absoluteLineNumbers[i - 1]}` : "on",
                lineNumbersMinChars: 9,
            }}
        />
    );
};

export default LogViewer;

export interface LogViewerProps {
    data: string;
    absoluteLineNumbers: string[] | null;
}
