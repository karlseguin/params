package params

type MapParams map[string]string

func (p MapParams) GetIf(key string) (string, bool) {
	v, k := p[key]
	return v, k
}

func (p MapParams) Set(key, value string) bool {
	p[key] = value
	return true
}

func (p MapParams) Delete(key string) (string, bool) {
	value, exists := p.GetIf(key)
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

func (p MapParams) Len() int {
	return len(p)
}

func (p MapParams) Clear() {
}

func (p MapParams) ToMap(key, value string) params {
	return p
}
