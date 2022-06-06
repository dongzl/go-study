package uber_go_guide

import "sync"

type Stats struct {
	mu       sync.Mutex
	counters map[string]int
}

func (s *Stats) SnapshotGood() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make(map[string]int, len(s.counters))
	for k, v := range s.counters {
		result[k] = v
	}
	return result
}

// Snapshot returns the current stats.
func (s *Stats) SnapshotBad() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.counters
}
