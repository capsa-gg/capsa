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
        ]
    },
    ignores: [(commit) => commit.includes("Merge")],
}
