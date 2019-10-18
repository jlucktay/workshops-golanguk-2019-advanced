/*
Package prof contains the following:
- Write a benchmark for the strings.ToUpper function.
- Uppercase the following string in the benchmark: Education is what remains after one has forgotten what one has learned in school.
- Using the pprof tool, generate a cpu profile
- Use the following commands: top, list, web
- Extra Credit - Name who the quote is from.
*/
package prof

import (
	"strings"
	"testing"
)

const benchTarget = "Education is what remains after one has forgotten what one has learned in school."

var sink string

func BenchmarkUp(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = strings.ToUpper(benchTarget)
	}
	sink = s
}
