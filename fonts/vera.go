package fonts

import (
	_ "embed"
)

// VeraSans is vera.ttf font inlined to the bytes slice.
//
//go:embed verdana.ttf
var VeraSans []byte
