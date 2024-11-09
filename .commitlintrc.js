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
            ['protocol', 'server', 'web', 'webapp' /* NOTE: prefer web over webapp */, 'deployment', 'tools']
        ],
        'body-max-line-length': [2, 'always', 200]
    },
    ignores: [(commit) => commit.includes("Merge")],
};
