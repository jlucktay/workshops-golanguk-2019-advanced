package packit

import (
	"testing"

	"github.com/kr/pretty"
)

func TestPack(t *testing.T) {
	p, err := Pack("one", 1, 1.1, []byte("bytes"))
	if err != nil {
		t.Fatal(err)
	}
	values, err := Unpack(p)
	if err != nil {
		t.Fatal(err)
	}
	if len(values) != 4 {
		t.Log(pretty.Sprint(values))
		t.Fatalf("unexpected value length.  got: %d, exp:% d", len(values), 5)
	}
	s, ok := values[0].(string)
	if !ok {
		t.Errorf("expected string, got %T", values[0])
	}
	if s != "one" {
		t.Errorf("unexpected string value: got: %s, exp: %s", s, "one")
	}

	i, ok := values[1].(int)
	if !ok {
		t.Errorf("expected int, got %T", values[1])
	}
	if i != 1 {
		t.Errorf("unexpected int value: got: %d, exp: %d", i, 1)
	}
	f, ok := values[2].(float64)
	if !ok {
		t.Errorf("expected float64, got %T", values[2])
	}
	if f != 1.1 {
		t.Errorf("unexpected float64 value: got: %.2f, exp: %.2f", f, 1.1)
	}
	b, ok := values[3].([]byte)
	if !ok {
		t.Errorf("expected []byte, got %T", values[3])
	}
	if string(b) != "bytes" {
		t.Errorf("unexpected byte value: got: %s, exp: %s", string(b), "bytes")
	}
}

var sink error

// section: bench-strings
func BenchmarkPackStrings(b *testing.B) {
	var err error
	for n := 0; n < b.N; n++ {
		_, err = Pack("one", "two", "three", "four", "five")
	}
	sink = err
}

// section: bench-strings

// section: bench-everything
func BenchmarkPackEverything(b *testing.B) {
	var err error
	for n := 0; n < b.N; n++ {
		_, err = Pack(
			"one",
			1,
			1.1,
			[]byte{1},
			"two",
			2,
			2.2,
			[]byte{2},
			"three",
			3,
			3.3,
			[]byte{3},
			"four",
			4,
			4.4,
			[]byte{4},
			"five",
			5,
			5.5,
			[]byte{5},
		)
	}
	sink = err
}

// section: bench-everything
