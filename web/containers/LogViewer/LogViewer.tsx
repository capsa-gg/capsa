import * as React from "react";
import { Editor } from "@monaco-editor/react";
import { initializeMonaco, monacoOptions } from "./LogViewer.settings";

const LogViewer: React.FC<LogViewerProps> = ({ data, absoluteLineNumbers }) => {
    return (
        <Editor
            value={data}
            language="capsalog"
            onMount={initializeMonaco}
            options={{
                ...monacoOptions,
                lineNumbers: absoluteLineNumbers ? (i: number) => `${absoluteLineNumbers[i - 1]}` : "on",
            }}
        />
    );
};

export default LogViewer;

export interface LogViewerProps {
    data: string;
    absoluteLineNumbers: number[] | null;
}
