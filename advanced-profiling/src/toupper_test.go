package profiling

import (
	"strings"
	"testing"
)

var Result string

func BenchmarkToUpper(b *testing.B) {
	var u string
	for i := 0; i < b.N; i++ {
		u = strings.ToUpper("Education is what remains after one has forgotten what one has learned in school.")
	}
	Result = u
}
