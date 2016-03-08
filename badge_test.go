package badge

import (
	"io/ioutil"
	"testing"
)

func BenchmarkRenderSeq(b *testing.B) {
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

func BenchmarkRenderParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := Render("XXX", "YYY", ColorBlue, ioutil.Discard)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
