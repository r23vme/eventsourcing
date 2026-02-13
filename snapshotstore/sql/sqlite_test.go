package sql_test

import (
	sqldriver "database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/r23vme/eventsourcing/core"
	"github.com/r23vme/eventsourcing/core/testsuite"
	"github.com/r23vme/eventsourcing/snapshotstore/sql"
)

func TestSuite(t *testing.T) {
	f := func() (core.SnapshotStore, func(), error) {
		return snapshotstore()
	}
	testsuite.TestSnapshotStore(t, f)
}

func snapshotstore() (*sql.SQLite, func(), error) {
	db, err := sqldriver.Open("sqlite3", "file::memory:?locked.sqlite?cache=shared")
	if err != nil {
		return nil, nil, err
	}

	db.SetMaxOpenConns(1)
	err = db.Ping()
	if err != nil {
		return nil, nil, err
	}

	store, err := sql.NewSQLite(db)
	if err != nil {
		return nil, nil, err
	}

	return store, func() {
		store.Close()
	}, nil
}
