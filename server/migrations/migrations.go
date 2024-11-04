package migrations

import "embed"

// Migrations contains the embedded file system with migration files
//
//go:embed *.sql
var Migrations embed.FS
