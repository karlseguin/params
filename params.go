// a key-value (string, string) wrapper
package params

// An interface to a key-value lookup
type params interface {
	GetIf(key string) (string, bool)
	Set(key, value string) bool
	ToMap(key, value string) params
	Delete(key string) (string, bool)
	Each(func(key, value string))
	Len() int
	Clear()
}

type Params struct {
	pool    *Pool
	fast    params
	current params
}

// Creates a new params object
func New(length int) *Params {
	return pooled(nil, length)
}

func pooled(pool *Pool, length int) *Params {
	ap := &ArrayParams{
		lookup: make([]struct{ key, value string }, length),
	}
	return &Params{
		pool:    pool,
		fast:    ap,
		current: ap,
	}
}

func (p *Params) Get(key string) string {
	value, _ := p.GetIf(key)
	return value
}

func (p *Params) GetIf(key string) (string, bool) {
	return p.current.GetIf(key)
}

func (p *Params) Set(key, value string) {
	if p.current.Set(key, value) == false {
		p.current = p.current.ToMap(key, value)
	}
}
func (p *Params) Delete(key string) (string, bool) {
	return p.current.Delete(key)
}

func (p *Params) Each(f func(key, value string)) {
	p.current.Each(f)
}

func (p *Params) Len() int {
	return p.current.Len()
}

func (p *Params) Clear() {
	p.current = p.fast
	p.current.Clear()
}

func (p *Params) Release() {
	if p.pool != nil {
		p.Clear()
	}
}
