import react from "eslint-plugin-react";
import prettier from "eslint-plugin-prettier";
import noRelativeImportPaths from "eslint-plugin-no-relative-import-paths";
import path from "node:path";
import { fileURLToPath } from "node:url";
import js from "@eslint/js";
import { FlatCompat } from "@eslint/eslintrc";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
const compat = new FlatCompat({
    baseDirectory: __dirname,
    recommendedConfig: js.configs.recommended,
    allConfig: js.configs.all,
});

export default [
    ...compat.extends(
        "plugin:react/recommended",
        "plugin:@next/next/recommended",
        "airbnb",
        "prettier",
        "next/core-web-vitals",
        "next/typescript",
    ),
    {
        plugins: {
            react,
            prettier,
            "no-relative-import-paths": noRelativeImportPaths,
        },

        languageOptions: {
            globals: {
                React: false,
            },
        },

        // In alphabetical order, all rules require a comment to explain the change and necessity
        rules: {
            "import/extensions": 0, // The Typescript compiler will check imports extensions
            "import/no-named-as-default": 0, // Gives false positives
            "import/no-unresolved": 0, // Typescript compiler checks for import paths, configuring the @/ import syntax is not worth it
            "import/prefer-default-export": 0, // Allow 'export const' exports

            // Require absolute import paths, except for current directory
            "no-relative-import-paths/no-relative-import-paths": [
                2,
                {
                    allowSameFolder: true,
                },
            ],

            "no-use-before-define": 0, // Allow sane file layouts
            "object-curly-spacing": ["error", "always"], // Force consistency for { item } instead of {item}

            // Force arrow function
            "react/function-component-definition": [
                1,
                {
                    namedComponents: "arrow-function",
                    unnamedComponents: "arrow-function",
                },
            ],

            // Force files to have .jsx or .tsx
            "react/jsx-filename-extension": [
                1,
                {
                    extensions: [".tsx", ".jsx"],
                },
            ],

            "react/jsx-props-no-spreading": 0, // Spreading necessary for form registration, just use with caution
            "react/prop-types": 0, // Done by using React.FC<PropsInterface>
            "react/require-default-props": 0, // Otherwise you get false positives, as regular React ESLint rules doesn't fully get Typescript React code
            "@typescript-eslint/no-use-before-define": 0, // Allow sane file layouts
        },
    },
    {
        files: ["**/*.test.js", "**/*.test.jsx", "**/*.test.ts", "**/*.test.tsx"], // All test files
        rules: {
            // Test files import devDependencies
            "import/no-extraneous-dependencies": [
                "error",
                {
                    devDependencies: true,
                },
            ],
            "@typescript-eslint/no-empty-function": 0, // Allow `() => {}` in test files
        },
    },
];
