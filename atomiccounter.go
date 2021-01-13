// Edge Network
// (c) 2020 Edge Network technologies Ltd.

// Edge Network
// (c) 2020 Edge Network technologies Ltd.

package atomiccounter

import (
	"sync/atomic"
)

// Counter is an integer.
type Counter uint64

// Add adds an input value to the counter.
func (c *Counter) Add(val uint64) uint64 {
	return atomic.AddUint64((*uint64)(c), val)
}

// Sub reduces the value of the counter by an input value.
func (c *Counter) Sub(val uint64) uint64 {
	for v := c.Get(); (v - val) >= 0; v = c.Get() {
		if atomic.CompareAndSwapUint64((*uint64)(c), v, v-val) {
			return v - val
		}
	}
	return 0
}

// Inc increments the counter by 1.
func (c *Counter) Inc() uint64 {
	return atomic.AddUint64((*uint64)(c), 1)
}

// Dec reduces the counter by 1.
func (c *Counter) Dec() uint64 {
	for v := c.Get(); v > 0; v = c.Get() {
		if atomic.CompareAndSwapUint64((*uint64)(c), v, v-1) {
			return v - 1
		}
	}
	return 0
}

// Set sets the value of the counter.
func (c *Counter) Set(v uint64) {
	atomic.StoreUint64((*uint64)(c), v)
}

// Get gets the value of the counter.
func (c *Counter) Get() uint64 {
	return atomic.LoadUint64((*uint64)(c))
}

// New creates a new counter.
func New() *Counter {
	return new(Counter)
}
