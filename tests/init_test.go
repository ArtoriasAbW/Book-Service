//go:build integration
// +build integration

package tests

import (
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/tests/postgres"
)

var (
	Db *postgres.TDB
)

func init() {
	Db = postgres.NewFromEnv()
}
