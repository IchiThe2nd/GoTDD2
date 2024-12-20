package clockface_test

import (
	"strings"
	"testing"
	"time"

	"github.com/IchiThe2nd/GoTDD2/clockface"
	//"github.com/IchiThe2nd/GoTDD2/clockface" saved because I keep fucking up imports
)

func TestSVGWriterAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	var b strings.Builder
	clockface.SVGWriter(&b, tm)
	got := b.String()

	want := `<line x1="150" y1="150" x2="150.000" y2="240.000"`

	if !strings.Contains(got, want) {
		t.Errorf("Expected to find the second hand %v, in the SVG output %v", want, got)
	}
}
