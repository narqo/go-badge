package badge

import (
	"html/template"
	"path/filepath"
	"sync"
	"io"
	"io/ioutil"

	"golang.org/x/image/font"
	"github.com/golang/freetype/truetype"
)

type drawer struct {
	once     sync.Once
	fd       *font.Drawer
	tmplName string
	tmpl     *template.Template
}

func (d *drawer) render(subject, status string, color Color, w io.Writer) error {
	d.once.Do(func() {
		d.tmpl = template.Must(template.ParseFiles(filepath.Join("../templates", d.tmplName)))
	})

	subjectDx := d.measureString(subject)
	statusDx := d.measureString(status)

	data := map[string]interface{}{
		"Subject": subject,
		"Status": status,
		"Color": color,
		"Bounds": map[string]float64{
			"Dx": subjectDx + statusDx,
			"SubjectDx": subjectDx,
			"SubjectX": subjectDx / 2.0,
			"StatusDx": statusDx,
			"StatusX": subjectDx + statusDx / 2.0 - 1,
		},
	}
	return d.tmpl.Execute(w, data)
}

// shild.io uses Verdana.ttf to measure text width and an extra of 10px.
// As we use DejaVuSans.ttf, we have to tune this value a little.
const extraDx = 13

func (d *drawer) measureString(s string) float64 {
	sm := d.fd.MeasureString(s)
	return float64(sm) / 64 + extraDx
}

var defaultDrawer *drawer

func Render(subject, status string, color Color, w io.Writer) error {
	return defaultDrawer.render(subject, status, color, w)
}

const (
	dpi = 72
	fontsize = 11
	fontfile = "../res/dejavu-sans/ttf/DejaVuSans.ttf"
)

func init() {
	fontBytes, err := ioutil.ReadFile(fontfile);
	if err != nil {
		panic(err)
	}
	ttf, err := truetype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}
	fd := &font.Drawer{
		Face: truetype.NewFace(ttf, &truetype.Options{
			Size: fontsize,
			DPI: dpi,
			Hinting: font.HintingFull,
		}),
	}
	defaultDrawer = &drawer{
		fd: fd,
		tmplName: "flat-template.svg",
	}
}
