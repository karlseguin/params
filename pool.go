package params

import (
	"sync/atomic"
)

// A pool of Params
type Pool struct {
	misses int64
	size   int
	list   chan *Params
}

// Create a new pool of count params.
// Each ArrayParam has a maximum size of size.
// If the size exceeds this value, the ArrayParam is converted to a MapParams
func NewPool(size, count int) *Pool {
	pool := &Pool{
		size: size,
		list: make(chan *Params, count),
	}
	for i := 0; i < count; i++ {
		pool.list <- pooled(pool, size)
	}
	return pool
}

// The number of times we tried to checkout but had no available
// params in our pool
func (p *Pool) Misses() int64 {
	return atomic.SwapInt64(&p.misses, 0)
}

// Get a param from the pool or create a new one if the pool is empty
func (p *Pool) Checkout() *Params {
	select {
	case item := <-p.list:
		return item
	default:
		atomic.AddInt64(&p.misses, 1)
		return New(p.size)
	}
}
