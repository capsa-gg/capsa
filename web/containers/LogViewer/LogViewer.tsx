import * as React from "react";
import { Editor, useMonaco } from "@monaco-editor/react";
import { initializeMonaco, monacoOptions } from "./LogViewer.settings";
import { useEffect } from "react";

const LogViewer: React.FC<LogViewerProps> = ({ data }) => {
    const monaco = useMonaco();

    useEffect(() => {
        if (monaco) {
            initializeMonaco(monaco);
        }
    }, [monaco]);

    return <Editor value={data} options={monacoOptions} />;
};

export default LogViewer;

export interface LogViewerProps {
    data: string;
}
