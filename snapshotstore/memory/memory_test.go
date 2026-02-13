package memory_test

import (
	"testing"

	"github.com/r23vme/eventsourcing/core"
	"github.com/r23vme/eventsourcing/core/testsuite"
	"github.com/r23vme/eventsourcing/snapshotstore/memory"
)

func TestSuite(t *testing.T) {
	f := func() (core.SnapshotStore, func(), error) {
		ss := memory.Create()
		return ss, func() { ss.Close() }, nil
	}
	testsuite.TestSnapshotStore(t, f)
}
