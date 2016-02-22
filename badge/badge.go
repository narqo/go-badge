package badge

import (
	"html/template"
	"io"
	"io/ioutil"
	"path/filepath"
	"sync"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type badgeDrawer struct {
	once     sync.Once
	fd       *font.Drawer
	tmplName string
	tmpl     *template.Template
}

func (d *badgeDrawer) Render(subject, status string, color Color, w io.Writer) error {
	d.once.Do(func() {
		d.tmpl = template.Must(template.ParseFiles(filepath.Join("../templates", d.tmplName)))
	})

	subjectDx := d.measureString(subject)
	statusDx := d.measureString(status)

	data := map[string]interface{}{
		"Subject": subject,
		"Status":  status,
		"Color":   color,
		"Bounds":  map[string]float64{
			"Dx":        subjectDx + statusDx,
			"SubjectDx": subjectDx,
			"SubjectX":  subjectDx/2.0 + 1,
			"StatusDx":  statusDx,
			"StatusX":   subjectDx + statusDx/2.0 - 1,
		},
	}
	return d.tmpl.Execute(w, data)
}

// shield.io uses Verdana.ttf to measure text width with an extra 10px.
// As we use Vera.ttf, we have to tune this value a little.
const extraDx = 13

func (d *badgeDrawer) measureString(s string) float64 {
	sm := d.fd.MeasureString(s)
	return float64(sm)/64 + extraDx
}

var drawer *badgeDrawer

// Render renders a badge of the given color, with given subject and status to w.
func Render(subject, status string, color Color, w io.Writer) error {
	return drawer.Render(subject, status, color, w)
}

const (
	dpi      = 72
	fontsize = 11
	fontfile = "../res/bitstream-vera-sans/Vera.ttf"
)

func init() {
	setDrawer(&badgeDrawer{tmplName: "flat-template.svg"})
}

func setDrawer(d *badgeDrawer) {
	d.fd = mustNewFontDrawer(fontfile, fontsize, dpi)
	drawer = d
}

func mustNewFontDrawer(fontfile string, size, dpi float64) *font.Drawer {
	fontBytes, err := ioutil.ReadFile(fontfile)
	if err != nil {
		panic(err)
	}
	ttf, err := truetype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}
	return &font.Drawer{
		Face: truetype.NewFace(ttf, &truetype.Options{
			Size:    size,
			DPI:     dpi,
			Hinting: font.HintingFull,
		}),
	}
}
