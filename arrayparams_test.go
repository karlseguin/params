package params

import (
	. "github.com/karlseguin/expect"
	"testing"
)

type ArrayParamsTests struct{}

func Test_ArrayParams(t *testing.T) {
	Expectify(new(ArrayParamsTests), t)
}

func (_ ArrayParamsTests) GetsFromEmpty() {
	Expect(New(1).Get("baron:friends")).To.Equal("", false)
}

func (_ ArrayParamsTests) GetsAValue() {
	p := New(10)
	p = p.Set("leto", "ghanima")
	p = p.Set("paul", "alia")
	p = p.Set("duncan", "")
	Expect(p.Get("leto")).To.Equal("ghanima", true)
	Expect(p.Get("paul")).To.Equal("alia", true)
	Expect(p.Get("duncan")).To.Equal("", true)
	Expect(p.Get("vladimir")).To.Equal("", false)
}

func (_ ArrayParamsTests) OverwritesAnExistingValue() {
	p := New(10)
	p = p.Set("leto", "ghaima")
	p = p.Set("leto", "ghanima")
	Expect(p.Get("leto")).To.Equal("ghanima", true)
}

func (_ ArrayParamsTests) ExpandsBeyondTheSpecifiedSize() {
	p := New(1)
	p = p.Set("leto", "ghanima")
	p = p.Set("paul", "alia")
	Expect(p.Get("leto")).To.Equal("ghanima", true)
	Expect(p.Get("paul")).To.Equal("alia", true)
	Expect(p.Get("vladimir")).To.Equal("", false)
	_, ok := p.(MapParams)
	Expect(ok).ToEqual(true)
}

func (_ ArrayParamsTests) ExpansionReleasesTheParam() {
	pool := NewPool(1, 1)
	p := pool.Checkout()
	Expect(len(pool.list)).To.Equal(0)
	p = p.Set("leto", "ghanima")
	p = p.Set("paul", "alia")
	Expect(len(pool.list)).To.Equal(1)
	Expect(pool.Checkout().(*ArrayParams).length).ToEqual(0)
}

func (_ ArrayParamsTests) Iterates() {
	p := New(10)
	p = p.Set("leto", "ghanima")
	p = p.Set("paul", "alia")
	saw := make(map[string]string, 2)
	p.Each(func(key, value string) {
		saw[key] = value
	})
	Expect(saw["leto"]).To.Equal("ghanima")
	Expect(saw["paul"]).To.Equal("alia")
	Expect(len(saw)).To.Equal(2)
}

func (_ ArrayParamsTests) ClearsTheParam() {
	p := New(10)
	p = p.Set("leto", "ghanima")
	p = p.Set("paul", "alia")
	p.Clear()
	Expect(p.Len()).To.Equal(0)
	Expect(p.Get("leto")).To.Equal("", false)
}

func (_ ArrayParamsTests) DeletesFromNothing() {
	p := New(10)
	Expect(p.Delete("x")).To.Equal("", false)
}

func (_ ArrayParamsTests) DeletesNonExistingKey() {
	p := New(10)
	p.Set("leto", "ghanima")
	Expect(p.Delete("y")).To.Equal("", false)
}

func (_ ArrayParamsTests) DeletesOneItem() {
	p := New(10)
	p.Set("leto", "ghanima")
	Expect(p.Delete("leto")).To.Equal("ghanima", true)
	Expect(p.Len()).To.Equal(0)
}

func (_ ArrayParamsTests) DeletesAnItem() {
	p := New(10)
	p = p.Set("leto", "ghanima")
	p = p.Set("paul", "alia")
	p = p.Set("duncan", "")
	Expect(p.Delete("leto")).To.Equal("ghanima", true)
	Expect(p.Len()).To.Equal(2)
	Expect(p.Get("duncan")).To.Equal("", true)
	Expect(p.Get("paul")).To.Equal("alia", true)
	Expect(p.Get("leto")).To.Equal("", false)
}
