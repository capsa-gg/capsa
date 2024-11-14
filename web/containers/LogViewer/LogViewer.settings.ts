import * as Monaco from "monaco-editor";

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
                [/\[\d{4}\.\d{2}\.\d{2}-\d{2}\.\d{2}\.\d{2}:\d{3}]/, "timestamp"],
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
            { token: "", foreground: "#cccccc" },
            { token: "timestamp", foreground: "#666666" },
            { token: "loglevel.fatal", foreground: "#AE81FF" },
            { token: "loglevel.error", foreground: "#f92672" },
            { token: "loglevel.warning", foreground: "#e2e22e" },
            { token: "loglevel.info", foreground: "#23aa59" },
            { token: "loglevel.verbose", foreground: "#A6E22E" },
            { token: "loglevel.veryverbose", foreground: "#66D9EF" },
            { token: "category", foreground: "#E6DB74", fontStyle: "italic" },
        ],
    });
    monaco.editor.setTheme("capsalogtheme");
};
