package params

// A key-value pair stored in an array [k1, v1, k2, v2, ..., kn, vn]
// More memory efficient and faster for a small number of keys (~10)
// Also, can be pooled and reused
type ArrayParams struct {
	pool   *Pool
	length int
	lookup []struct{ key, value string }
}

// Create a new ArrayParam that can hold up to length pairs
func New(length int) Params {
	return pooled(nil, length)
}

func pooled(pool *Pool, length int) Params {
	return &ArrayParams{
		pool:   pool,
		lookup: make([]struct{ key, value string }, length),
	}
}

// Get a value by key
func (p *ArrayParams) Get(key string) (string, bool) {
	position, exists := p.indexOf(key)
	if exists == false {
		return "", false
	}
	return p.lookup[position].value, true
}

// Set the value to the specified key
// Set can return a new Params object better equipped to handle the large size
// (always assign the result of Set back to the params variable, like you do with append)
func (p *ArrayParams) Set(key, value string) Params {
	position, exists := p.indexOf(key)
	if exists {
		p.lookup[position].value = value
		return p
	}

	if p.length == len(p.lookup) {
		m := p.toMap(key, value)
		p.Release()
		return m
	}

	pair := struct{ key, value string }{key, value}
	if position != p.length {
		copy(p.lookup[position+1:], p.lookup[position:])
		p.lookup[position] = pair
	} else {
		p.lookup[position] = pair
	}
	p.length += 1
	return p
}

// Iterate over each key value pair
func (p *ArrayParams) Each(f func(string, value string)) {
	for i, l := 0, p.length; i < l; i++ {
		pair := p.lookup[i]
		f(pair.key, pair.value)
	}
}

// Get the number of pairs
func (p *ArrayParams) Len() int {
	return p.length
}

// Clears the param
func (p *ArrayParams) Clear() Params {
	p.length = 0
	return p
}

// Delete a value by key
func (p *ArrayParams) Delete(key string) (string, bool) {
	position, exists := p.indexOf(key)
	if exists == false {
		return "", false
	}
	value := p.lookup[position].value
	for i := position + 1; i < p.length; i++ {
		p.lookup[i-1] = p.lookup[i]
	}
	p.length--
	return value, true
}

func (p *ArrayParams) toMap(key, value string) Params {
	m := make(MapParams, p.length+1)
	p.Each(func(key, value string) { m[key] = value })
	m[key] = value
	return m
}

func (p *ArrayParams) indexOf(key string) (int, bool) {
	for i := 0; i < p.length; i++ {
		k := p.lookup[i].key
		if k == key {
			return i, true
		}
		if k > key {
			return i, false
		}
	}
	return p.length, false
}

// Return the params to the pool
// Safe to call on non-pooled params
func (p *ArrayParams) Release() {
	if p.pool != nil {
		p.length = 0
		p.pool.list <- p
	}
}
