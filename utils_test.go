package badge

import (
	"testing"
)

// Keep newlines, but clean up whitespace
func TestStripXmlWhitespace(t *testing.T) {
	const mock = `<xml>
	<prop></prop>
<prop></prop><prop>
aaa
</prop>  
	<prop>aaaa
	</prop>
<xml>  `
	const expected = `<xml><prop></prop><prop></prop><prop>aaa
</prop><prop>aaaa
	</prop><xml>`

	if stripXmlWhitespace(mock) != expected {
		t.Errorf("stripXmlWhitespace failed")
	}
}
