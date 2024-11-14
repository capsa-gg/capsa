import * as React from "react";
import { Editor } from "@monaco-editor/react";
import { initializeMonaco, monacoOptions } from "./LogViewer.settings";

const LogViewer: React.FC<LogViewerProps> = ({ data }) => {
    return <Editor value={data} language="capsalog" onMount={initializeMonaco} options={monacoOptions} />;
};

export default LogViewer;

export interface LogViewerProps {
    data: string;
}
