package memory_test

import (
	"testing"

	"github.com/r23vme/eventsourcing/core"
	"github.com/r23vme/eventsourcing/core/testsuite"
	"github.com/r23vme/eventsourcing/eventstore/memory"
)

func TestEventStore(t *testing.T) {
	f := func() (core.EventStore, func(), error) {
		es := memory.Create()
		return es, func() { es.Close() }, nil
	}
	testsuite.Test(t, f)
}

func TestFetcherAll(t *testing.T) {
	es := memory.Create()
	defer es.Close()
	testsuite.TestFetcher(t, es, es.All(0, 10))
}
