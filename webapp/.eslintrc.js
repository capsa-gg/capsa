module.exports = {
    extends: [
        "plugin:react/recommended",
        "plugin:@next/next/recommended",
        "airbnb",
        "prettier",
        "next/core-web-vitals",
        // "next/typescript",
    ],
    plugins: ["react", "prettier", "no-relative-import-paths"],
    globals: { React: false },
    rules: {
        // In alphabetical order, all rules require a comment to explain the change and necessity
        "import/extensions": 0, // The Typescript compiler will check imports extensions
        "import/no-unresolved": 0, // Typescript compiler checks for import paths, configuring the @/ import syntax is not worth it
        "import/prefer-default-export": 0, // Allow 'export const' exports
        "no-relative-import-paths/no-relative-import-paths": [2, { allowSameFolder: true }], // Require absolute import paths, except for current directory
        "no-use-before-define": 0, // Allow sane file layouts
        "object-curly-spacing": ["error", "always"], // Force consistency for { item } instead of {item}
        "react/function-component-definition": [
            // Force arrow function
            1,
            {
                namedComponents: "arrow-function",
                unnamedComponents: "arrow-function",
            },
        ],
        "react/jsx-filename-extension": [1, { extensions: [".tsx", ".jsx"] }], // Force files to have .jsx or .tsx
        "react/jsx-props-no-spreading": 0, // Spreading necessary for form registration, just use with caution
        "react/prop-types": 0, // Done by using React.FC<PropsInterface>
        "react/require-default-props": 0, // Otherwise you get false positives, as regular React ESLint rules doesn't fully get Typescript React code
        "@typescript-eslint/no-use-before-define": 0, // Allow sane file layouts
    },
    overrides: [
        {
            files: ["*.test.js", "*.test.jsx", "*.test.ts", "*.test.tsx"], // All test files
            rules: {
                "import/no-extraneous-dependencies": ["error", { devDependencies: true }], // Test files import devDependencies
                "@typescript-eslint/no-empty-function": 0, // Allow `() => {}` in test files
            },
        },
    ],
};
