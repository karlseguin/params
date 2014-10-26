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

func (a *ArrayParamsTests) GetsAnValue() {
	p := New(10)
	p = p.Set("leto", "ghanima")
	p = p.Set("paul", "alia")
	Expect(p.Get("leto")).To.Equal("ghanima")
	Expect(p.Get("paul")).To.Equal("alia")
	Expect(p.Get("vladimir")).To.Equal("")
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
