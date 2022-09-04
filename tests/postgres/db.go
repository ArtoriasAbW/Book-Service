//go:build integration
// +build integration

package postgres

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/tests/config"
)

type TDB struct {
	DB *sqlx.DB
	sync.Mutex
}

func NewFromEnv() *TDB {
	cfg, err := config.FromEnv()
	if err != nil {
		log.Fatal(err.Error())
	}
	if err != nil {
		panic(err.Error())
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)
	fmt.Println(psqlInfo)
	db := sqlx.MustOpen("postgres", psqlInfo)
	return &TDB{
		DB: db,
	}

}

func (d *TDB) Setup(t *testing.T) {
	t.Helper()
	ctx := context.Background()
	d.Lock()
	d.Truncate(ctx)
}

func (d *TDB) TearDown() {
	defer d.Unlock()
	d.Truncate(context.Background())
}

func (db *TDB) Truncate(ctx context.Context) {
	var tables []string
	err := db.DB.Select(&tables, "SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE'"+
		"and table_name != 'goose_db_version'")

	if err != nil {
		log.Fatal(err.Error())
	}
	if len(tables) == 0 {
		log.Fatal("run migrations please")
	}
	q := fmt.Sprintf("Truncate table %s", strings.Join(tables, ","))
	if _, err := db.DB.Exec(q); err != nil {
		log.Fatal(err.Error())
	}
}
