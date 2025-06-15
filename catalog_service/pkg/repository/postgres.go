package repository

import (
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPostgresdb(DBUrl string) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", os.Getenv("POSTGRES_URL"))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(migrationsPath string, DBUrl string) error {
	m, err := migrate.New("file://"+migrationsPath, DBUrl)
	if err != nil {
		panic(err)
	}
	defer m.Close()

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		m.Down()
		return err
	}
	return nil
}
