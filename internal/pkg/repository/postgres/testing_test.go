package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	repoPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository"
)

type bookRepoFixtures struct {
	bookRepo repoPkg.Interface
	db       *sqlx.DB
	dbMock   sqlmock.Sqlmock
}

// func setUp(t *testing.T) bookRepoFixtures {
// 	var fixture bookRepoFixtures
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("some error")unc setUp(t *testing.T) bookRepoFixtures {
// 	var fixture bookRepoFixtures
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("some error")
// 	}
// 	fixture.db = sqlx.NewDb(db, "sqlmock")
// 	fixture.dbMock = mock
// 	fixture.bookRepo = *NewRepository(fixture.db)
// 	return fixture
// }
// 	}
// 	fixture.db = sqlx.NewDb(db, "sqlmock")
// 	fixture.dbMock = mock
// 	fixture.bookRepo = *NewRepository(fixture.db)
// 	return fixture
// }
