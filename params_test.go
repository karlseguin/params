package params

import (
	"testing"

	. "github.com/karlseguin/expect"
)

type ArrayParamsTests struct{}

func Test_ArrayParams(t *testing.T) {
	Expectify(new(ArrayParamsTests), t)
}

func (_ ArrayParamsTests) GetsFromEmpty() {
	Expect(New(1).GetIf("baron:friends")).To.Equal("", false)
}

func (_ ArrayParamsTests) GetsAValue() {
	p := New(10)
	p.Set("leto", "ghanima")
	p.Set("paul", "alia")
	p.Set("duncan", "")
	Expect(p.GetIf("leto")).To.Equal("ghanima", true)
	Expect(p.GetIf("paul")).To.Equal("alia", true)
	Expect(p.GetIf("duncan")).To.Equal("", true)
	Expect(p.GetIf("vladimir")).To.Equal("", false)

	Expect(p.Get("paul")).To.Equal("alia")
	Expect(p.Get("vladimir")).To.Equal("")
}

func (_ ArrayParamsTests) OverwritesAnExistingValue() {
	p := New(10)
	p.Set("leto", "ghaima")
	p.Set("leto", "ghanima")
	Expect(p.GetIf("leto")).To.Equal("ghanima", true)
}

func (_ ArrayParamsTests) ExpandsBeyondTheSpecifiedSize() {
	p := New(1)
	p.Set("leto", "ghanima")
	p.Set("paul", "alia")
	Expect(p.GetIf("leto")).To.Equal("ghanima", true)
	Expect(p.GetIf("paul")).To.Equal("alia", true)
	Expect(p.GetIf("vladimir")).To.Equal("", false)
	_, ok := p.current.(MapParams)
	Expect(ok).ToEqual(true)
}

func (_ ArrayParamsTests) Iterates() {
	p := New(10)
	p.Set("leto", "ghanima")
	p.Set("paul", "alia")
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
	p.Set("leto", "ghanima")
	p.Set("paul", "alia")
	p.Clear()
	Expect(p.Len()).To.Equal(0)
	Expect(p.GetIf("leto")).To.Equal("", false)
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
	p.Set("leto", "ghanima")
	p.Set("paul", "alia")
	p.Set("duncan", "")
	Expect(p.Delete("leto")).To.Equal("ghanima", true)
	Expect(p.Len()).To.Equal(2)
	Expect(p.GetIf("duncan")).To.Equal("", true)
	Expect(p.GetIf("paul")).To.Equal("alia", true)
	Expect(p.GetIf("leto")).To.Equal("", false)
}
