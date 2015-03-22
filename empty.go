package params

var Empty EmptyParams

type EmptyParams struct{}

func (p EmptyParams) Get(key string) (string, bool) {
	return "", false
}

func (p EmptyParams) Set(key, value string) {
}

func (p EmptyParams) Each(f func(string, value string)) {
}

func (p EmptyParams) Release() {
}

func (p EmptyParams) Len() int {
	return 0
}

func (p EmptyParams) Delete(key string) (string, bool) {
	return "", false
}

func (p EmptyParams) Clear() {
}
