package postgres

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	repoPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository"
)

type bookRepoFixtures struct {
	repo   repoPkg.Interface
	db     *sqlx.DB
	dbMock sqlmock.Sqlmock
}

func setUp(t *testing.T) bookRepoFixtures {
	var fixture bookRepoFixtures
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("some error")
	}
	fixture.db = sqlx.NewDb(db, "sqlmock")
	fixture.dbMock = mock
	fixture.repo = NewRepository(fixture.db)
	return fixture
}

func (f *bookRepoFixtures) tearDown() {
	f.db.Close()
}
