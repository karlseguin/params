package params

type MapParams map[string]string

func (p MapParams) Get(key string) string {
	return p[key]
}

func (p MapParams) Set(key, value string) Params {
	p[key] = value
	return p
}

func (p MapParams) Release() {

}
