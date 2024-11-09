/**
 * @type {import('semantic-release').GlobalConfig}
 */
export default {
    branches: ["main"],
    plugins: [
        [
            // Analyse commits to determine next version
            "@semantic-release/commit-analyzer",
            {
                preset: "conventionalcommits",
                releaseRules: [
                    { breaking: true, release: "patch" }, // after 1.0: major
                    { revert: true, release: "patch" },
                    { type: "feat", release: "patch" }, // after 1.0: minor
                    { type: "fix", release: "patch" },
                    { type: "perf", release: "patch" },
                    { type: "build", release: "patch" },
                    { type: "ci", release: "patch" },
                    { type: "docs", release: "patch" },
                ],
            },
        ],

        // Generate release notes
        "@semantic-release/release-notes-generator",

        // Generate changelog
        "@semantic-release/changelog",

        // Run release preparation script to:
        // - bump version numbers in project
        // - build and zip server binaries
        // Publish script will
        // - build and push Docker images to registry
        [
            "@semantic-release/exec",
            {
                prepareCmd:
                    "./bin/prepare-release ${lastRelease.version} ${nextRelease.version}",
                publishCmd: "./bin/publish-docker ${nextRelease.version}"
            },
        ],

        // Commit changes made by release prep script back to repository
        [
            "@semantic-release/git",
            {
                message:
                    "release: ${nextRelease.version} [skip ci]\n\n${nextRelease.notes}",
                assets: [
                    "CHANGELOG.md",
                    "package.json",
                    "package-lock.json",
                    "web/package.json",
                    "web/package-lock.json",
                    "web/version.ts",
                    "server/constants/version.go",
                ],
            },
        ],

        // Create a release with release assets
        [
            "@semantic-release/github",
            {
                assets: [
                    {
                        path: "dist/*.zip",
                        label: "Capsa ${nextRelease.version}",
                    },
                ],
            },
        ],
    ],
};
