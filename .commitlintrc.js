export default {
    extends: ["@commitlint/config-conventional"],
    rules: {
        'type-enum': [
            2,
            'always',
            [
                'build',
                'ci',
                'chore',
                'deps',
                'docs',
                'feat',
                'fix',
                'perf',
                'refactor',
                'release',
                'revert',
                'style',
                'test',
            ]
        ],
        'scope-enum': [
            2,
            'always',
            ['protocol', 'server', 'webapp', 'deployment', 'tools']
        ],
        'max-body-length': [2, 'always', 200]
    },
    ignores: [(commit) => commit.includes("Merge")],
};
