package params

import (
	. "github.com/karlseguin/expect"
	"testing"
)

type ArrayParamsTests struct{}

func Test_ArrayParams(t *testing.T) {
	Expectify(new(ArrayParamsTests), t)
}

func (a *ArrayParamsTests) GetsFromEmpty() {
	Expect(New(1).Get("baron:friends")).To.Equal("")
}

func (a *ArrayParamsTests) GetsAValue() {
	p := New(10)
	p = p.Set("leto", "ghanima")
	p = p.Set("paul", "alia")
	Expect(p.Get("leto")).To.Equal("ghanima")
	Expect(p.Get("paul")).To.Equal("alia")
	Expect(p.Get("vladimir")).To.Equal("")
}

func (a *ArrayParamsTests) OverwritesAnExistingValue() {
	p := New(10)
	p = p.Set("leto", "ghaima")
	p = p.Set("leto", "ghanima")
	Expect(p.Get("leto")).To.Equal("ghanima")
}

func (a *ArrayParamsTests) ExpandsBeyondTheSpecifiedSize() {
	p := New(1)
	p = p.Set("leto", "ghanima")
	p = p.Set("paul", "alia")
	Expect(p.Get("leto")).To.Equal("ghanima")
	Expect(p.Get("paul")).To.Equal("alia")
	Expect(p.Get("vladimir")).To.Equal("")
	_, ok := p.(MapParams)
	Expect(ok).ToEqual(true)
}

func (a *ArrayParamsTests) ExpansionReleasesTheParam() {
	pool := NewPool(1, 1)
	p := pool.Checkout()
	Expect(len(pool.list)).To.Equal(0)
	p = p.Set("leto", "ghanima")
	p = p.Set("paul", "alia")
	Expect(len(pool.list)).To.Equal(1)
	Expect(pool.Checkout().(*ArrayParams).length).ToEqual(0)
}

func (a *ArrayParamsTests) Iterates() {
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
