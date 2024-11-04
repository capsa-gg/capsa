export default {
    extends: ["@commitlint/config-conventional"],
    rules: {
        'body-max-line-length': 200,
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
        ]
    },
    ignores: [(commit) => commit.includes("Merge")],
}
