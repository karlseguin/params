package params

var Empty EmptyParams

type EmptyParams struct {}

func (p EmptyParams) Get(key string) string {
	return ""
}

func (p EmptyParams) Set(key, value string) Params {
	return p
}


func (p EmptyParams) Each(f func(string, value string)) {
}

func (p EmptyParams) Release() {
}
