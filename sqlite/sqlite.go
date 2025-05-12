package sqlite

import (
	"binanceTemporal/entities"
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

type sqlite struct {
	DB *bun.DB
}

func New(path string) (*sqlite, error) {
	DB, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	bunDB := bun.NewDB(DB, sqlitedialect.New())
	_, err = bunDB.NewCreateTable().
		Model(new(entities.KlineEvent)).
		IfNotExists().
		Exec(context.TODO())

	if err != nil {
		return nil, err
	}
	sqlite := &sqlite{
		DB: bunDB,
	}
	return sqlite, nil
}

func (s *sqlite) AddKlineEvent(ctx context.Context, klineEvent *entities.KlineEvent) error {
	_, err := s.DB.NewInsert().
		Model(klineEvent).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
