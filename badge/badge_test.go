package badge

import (
	"io/ioutil"
	"testing"
)

func BenchmarkRender(b *testing.B) {
	// warm up
	Render("XXX", "YYY", ColorBlue, ioutil.Discard)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := Render("XXX", "YYY", ColorBlue, ioutil.Discard)
		if err != nil {
			b.Fatal(err)
		}
	}
}
