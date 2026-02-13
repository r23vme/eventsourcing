package bbolt_test

import (
	"os"
	"testing"

	"github.com/r23vme/eventsourcing/core"
	"github.com/r23vme/eventsourcing/core/testsuite"
	"github.com/r23vme/eventsourcing/eventstore/bbolt"
)

func TestEventStoreSuite(t *testing.T) {
	f := func() (core.EventStore, func(), error) {
		dbFile := "bolt.db"
		es, err := bbolt.New(dbFile)
		if err != nil {
			return nil, nil, err
		}
		return es, func() {
			es.Close()
			os.Remove(dbFile)
		}, nil
	}
	testsuite.Test(t, f)
}

func TestFetchFuncAll(t *testing.T) {
	dbFile := "bolt.db"
	es, err := bbolt.New(dbFile)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		es.Close()
		os.Remove(dbFile)
	}()

	testsuite.TestFetcher(t, es, es.All(0))
}
