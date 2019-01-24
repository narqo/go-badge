package badge

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"strings"
	"sync"
	"testing"
)

func TestBadgeDrawerRender(t *testing.T) {
	mockTemplate := strings.TrimSpace(`
	{{.Subject}},{{.Status}},{{.Color}},{{with .Bounds}}{{.SubjectX}},{{.SubjectDx}},{{.StatusX}},{{.StatusDx}},{{.Dx}}{{end}}
	`)
	mockFontSize := 11.0
	mockDPI := 72.0

	d := &badgeDrawer{
		fd:    mustNewFontDrawer(mockFontSize, mockDPI),
		tmpl:  template.Must(template.New("mock-template").Parse(mockTemplate)),
		mutex: &sync.Mutex{},
	}

	output := "XXX,YYY,#c0c0c0,18,34,50,34,68"

	var buf bytes.Buffer
	err := d.Render("XXX", "YYY", "#c0c0c0", &buf)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	result := buf.String()
	if result != output {
		t.Errorf("expect %q got %q", output, result)
	}
}

func TestBadgeDrawerRenderBytes(t *testing.T) {
	mockTemplate := strings.TrimSpace(`
	{{.Subject}},{{.Status}},{{.Color}},{{with .Bounds}}{{.SubjectX}},{{.SubjectDx}},{{.StatusX}},{{.StatusDx}},{{.Dx}}{{end}}
	`)
	mockFontSize := 11.0
	mockDPI := 72.0

	d := &badgeDrawer{
		fd:    mustNewFontDrawer(mockFontSize, mockDPI),
		tmpl:  template.Must(template.New("mock-template").Parse(mockTemplate)),
		mutex: &sync.Mutex{},
	}

	output := "XXX,YYY,#c0c0c0,18,34,50,34,68"

	bytes, err := d.RenderBytes("XXX", "YYY", "#c0c0c0")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if string(bytes) != output {
		t.Errorf("expect %q got %q", output, string(bytes))
	}
}

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
