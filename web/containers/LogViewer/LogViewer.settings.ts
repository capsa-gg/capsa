import * as Monaco from "monaco-editor";
import monokai from "monaco-themes/themes/Monokai.json";

export const monacoOptions: Monaco.editor.IStandaloneEditorConstructionOptions = {
    readOnly: true,
    smoothScrolling: true,
    minimap: { enabled: false },
    scrollBeyondLastLine: false,
};

export const initializeMonaco = (monaco: typeof Monaco): void => {
    // Set theme
    //@ts-ignore
    monaco.editor.defineTheme("monokai", monokai);
    monaco.editor.setTheme("monokai");
};
