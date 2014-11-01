// a key-value (string, string) wrapper
package params

// An interface to a key-value lookup
type Params interface {
	Get(key string) string
	Set(key, value string) Params
	Release()
	Each(func(key, value string))
	Len() int
}
