package range_test

// Range represents a value range.
type Range interface {
	// HasNext returns true if Range is not EOF.
	HasNext() bool
	// Next returns the next value.
	Next() interface{}
}
