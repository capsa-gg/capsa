import { LogsColorsDarkMode, LogsColorsLightMode } from "@/styles/LogsColors";
import type * as Monaco from "monaco-editor";

export const monacoOptions: Monaco.editor.IStandaloneEditorConstructionOptions = {
    readOnly: true,
    smoothScrolling: true,
    minimap: { enabled: false },
    scrollBeyondLastLine: false,
};

export const themes = {
    lightMode: "capsalogtheme-light",
    darkMode: "capsalogtheme-dark",
};

export const initializeMonaco =
    (initialTheme: string) =>
    (editor: Monaco.editor.IStandaloneCodeEditor, monaco: typeof Monaco): void => {
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

        monaco.editor.defineTheme(themes.lightMode, {
            // @ts-ignore
            base: "vs",
            inherit: true,
            colors: {
                "editor.background": "#ffffff",
                "editor.foreground": "#eeeeee",
                "selection.background": "#878b9180",
                "editor.selectionHighlightBackground": "#575b6180",
                "editor.selectionBackground": "#878b9180",
            },
            rules: [
                { token: "", foreground: LogsColorsLightMode.Contents },
                { token: "timestamp", foreground: LogsColorsLightMode.Timestamp },
                { token: "loglevel.fatal", foreground: LogsColorsLightMode.Fatal },
                { token: "loglevel.error", foreground: LogsColorsLightMode.Error },
                { token: "loglevel.warning", foreground: LogsColorsLightMode.Warning },
                { token: "loglevel.info", foreground: LogsColorsLightMode.Info },
                { token: "loglevel.verbose", foreground: LogsColorsLightMode.Verbose },
                { token: "loglevel.veryverbose", foreground: LogsColorsLightMode.VeryVerbose },
                { token: "category", foreground: LogsColorsLightMode.Category, fontStyle: "italic" },
            ],
        });

        monaco.editor.defineTheme(themes.darkMode, {
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
                { token: "", foreground: LogsColorsDarkMode.Contents },
                { token: "timestamp", foreground: LogsColorsDarkMode.Timestamp },
                { token: "loglevel.fatal", foreground: LogsColorsDarkMode.Fatal },
                { token: "loglevel.error", foreground: LogsColorsDarkMode.Error },
                { token: "loglevel.warning", foreground: LogsColorsDarkMode.Warning },
                { token: "loglevel.info", foreground: LogsColorsDarkMode.Info },
                { token: "loglevel.verbose", foreground: LogsColorsDarkMode.Verbose },
                { token: "loglevel.veryverbose", foreground: LogsColorsDarkMode.VeryVerbose },
                { token: "category", foreground: LogsColorsDarkMode.Category, fontStyle: "italic" },
            ],
        });

        monaco.editor.setTheme(initialTheme);
    };
