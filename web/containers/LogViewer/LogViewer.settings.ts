import * as Monaco from "monaco-editor";
import LogsColors from "@/styles/LogsColors";

export const monacoOptions: Monaco.editor.IStandaloneEditorConstructionOptions = {
    readOnly: true,
    smoothScrolling: true,
    minimap: { enabled: false },
    scrollBeyondLastLine: false,
};

export const initializeMonaco = (editor: Monaco.editor.IStandaloneCodeEditor, monaco: typeof Monaco): void => {
    monaco.languages.register({ id: "capsalog" });
    monaco.languages.setMonarchTokensProvider("capsalog", {
        tokenizer: {
            root: [
                [/\[\d{4}\.\d{2}\.\d{2}-\d{2}\.\d{2}\.\d{2}.\d{3}]/, "timestamp"],
                [/\[Fatal]/, "loglevel.fatal"],
                [/\[Error]/, "loglevel.error"],
                [/\[Warning]/, "loglevel.warning"],
                [/\[Display]/, "loglevel.info"],
                [/\[Log]/, "loglevel.info"],
                [/\[Verbose]/, "loglevel.verbose"],
                [/\[VeryVerbose]/, "loglevel.veryverbose"],
                [/\[[\w\\.]+]:/, "category"],
            ],
        },
    });
    monaco.editor.defineTheme("capsalogtheme", {
        // @ts-ignore
        base: "vs-dark",
        inherit: true,
        colors: {
            "editor.background": "#272822",
            "editor.foreground": "#f8f8f2",
            "selection.background": "#878b9180",
            "editor.selectionHighlightBackground": "#575b6180",
            "editor.selectionBackground": "#878b9180",
        },
        rules: [
            { token: "", foreground: LogsColors.Contents },
            { token: "timestamp", foreground: LogsColors.Timestamp },
            { token: "loglevel.fatal", foreground: LogsColors.Fatal },
            { token: "loglevel.error", foreground: LogsColors.Error },
            { token: "loglevel.warning", foreground: LogsColors.Warning },
            { token: "loglevel.info", foreground: LogsColors.Info },
            { token: "loglevel.verbose", foreground: LogsColors.Verbose },
            { token: "loglevel.veryverbose", foreground: LogsColors.VeryVerbose },
            { token: "category", foreground: LogsColors.Category, fontStyle: "italic" },
        ],
    });
    monaco.editor.setTheme("capsalogtheme");
};
