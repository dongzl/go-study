package uber_go_guide

import "sync"

// https://github.com/uber-go/guide/blob/master/style.md#zero-value-mutexes-are-valid
type SMap struct {
	mu   sync.Mutex
	data map[string]string
}

func NewSMap() *SMap {
	return &SMap{data: make(map[string]string)}
}
func (m *SMap) Get(k string) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.data[k]
}
