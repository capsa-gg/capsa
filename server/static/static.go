package static

import "embed"

// FS contains the embedded filesystem with static server files.
// Note: because `go:embed *` will include .go files, this file can be accessed through `[server]/v1/static/static.go`.
//
//go:embed *
var FS embed.FS
