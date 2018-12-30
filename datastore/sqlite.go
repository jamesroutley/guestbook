package datastore

import (
	"database/sql"

	"github.com/jamesroutley/guestbook/domain"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteDatastore struct {
	db *sql.DB
}

func NewSQLiteDatastore(path string) (Datastore, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return &sqliteDatastore{
		db: db,
	}, nil
}

func (d *sqliteDatastore) Store(visit domain.Visit) error {
	stmt, err := d.db.Prepare(
		"INSERT INTO visits(url, referrer, ip, created) values (?,?,?,?)",
	)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(visit.URL, visit.Referrer, visit.IP, visit.Created)
	if err != nil {
		return err
	}
	return nil
}
