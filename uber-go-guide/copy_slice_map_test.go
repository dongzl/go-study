package uber_go_guide

import (
	"testing"
)

func TestStats_SnapshotBad(t *testing.T) {
	// snapshot is no longer protected by the mutex, so any
	// access to the snapshot is subject to data races.
	stats := Stats{}
	stats.SnapshotBad()
}

func TestStats_SnapshotGood(t *testing.T) {
	// Snapshot is now a copy.
	stats := Stats{}
	stats.SnapshotGood()
}
