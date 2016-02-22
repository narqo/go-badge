package badge

type Color string

const (
	ColorBrightgreen = Color("brightgreen")
	ColorGreen       = Color("green")
	ColorYellow      = Color("yellow")
	ColorYellowgreen = Color("yellowgreen")
	ColorOrange      = Color("orange")
	ColorRed         = Color("red")
	ColorBlue        = Color("blue")
	ColorGrey        = Color("grey")
	ColorGray        = Color("gray")
	ColorLightgrey   = Color("lightgrey")
	ColorLightgray   = Color("lightgray")
)

var ColorScheme = map[string]string{
	"brightgreen": "#4c1",
	"green":       "#97ca00",
	"yellow":      "#dfb317",
	"yellowgreen": "#a4a61d",
	"orange":      "#fe7d37",
	"red":         "#e05d44",
	"blue":        "#007ec6",
	"grey":        "#555",
	"gray":        "#555",
	"lightgrey":   "#9f9f9f",
	"lightgray":   "#9f9f9f",
}

func (c Color) String() string {
	color, ok := ColorScheme[string(c)]
	if ok {
		return color
	}
	return string(c)
}
