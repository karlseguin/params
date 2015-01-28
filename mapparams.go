package params

type MapParams map[string]string

func (p MapParams) Get(key string) (string, bool) {
	v, k := p[key]
	return v, k
}

func (p MapParams) Set(key, value string) Params {
	p[key] = value
	return p
}

func (p MapParams) Delete(key string) (string, bool) {
	value, exists := p.Get(key)
	if exists {
		delete(p, key)
	}
	return value, exists
}

func (p MapParams) Each(f func(string, value string)) {
	for k, v := range p {
		f(k, v)
	}
}

func (p MapParams) Release() {

}

func (p MapParams) Len() int {
	return len(p)
}

func (p MapParams) Clear() Params {
	return make(MapParams)
}
