package badge

import (
	"io/ioutil"
	"testing"
)

func BenchmarkRender(b *testing.B) {
	Render("XXX", "YYY", ColorBlue, ioutil.Discard)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Render("XXX", "YYY", ColorBlue, ioutil.Discard)
	}
}
