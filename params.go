// a key-value (string, string) wrapper
package params

// An interface to a key-value lookup
type Params interface {
	Get(key string) (string, bool)
	Set(key, value string) Params
	Delete(key string) (string, bool)
	Release()
	Each(func(key, value string))
	Len() int
	Clear() Params
}
