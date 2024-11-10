// This config defines the rules for commit messages, the form is `type(scope): message`.
// The scope is optional, the message should always begin with a lowercase letter.
export default {
    extends: ["@commitlint/config-conventional"],
    rules: {
        'type-enum': [
            2, // Severity level 2: error
            'always',
            [
                'build',    // Changes related to the build system or external dependencies
                'ci',       // Changes related to Continuous Integration configuration
                'chore',    // Routine tasks and maintenance tasks
                'deps',     // Dependency updates
                'docs',     // Documentation changes
                'feat',     // New features
                'fix',      // Bug fixes
                'perf',     // Performance improvements
                'refactor', // Code refactoring (changes that neither fix a bug nor add a feature)
                'release',  // Release-related changes
                'revert',   // Reverting a previous commit
                'style',    // Code style changes (e.g., formatting, white spaces)
                'test',     // Adding or updating tests
            ]
        ],
        'scope-enum': [
            2, // Severity level 2: error
            'always',
            ['protocol', 'server', 'web', 'webapp' /* NOTE: prefer web over webapp */, 'deployment', 'tools']
        ],
        'body-max-line-length': [2, 'always', 200]
    },

    // Ignores commits that include the word "Merge", to allow for default merge commit messages.
    ignores: [(commit) => commit.includes("Merge")],
};
