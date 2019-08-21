package main

import "fmt"

func main() {
	// section: main
	s := Cat("hello ", "world")
	fmt.Println(s)

	b := BCat([]byte("hello "), []byte("world"))
	fmt.Println(string(b))
	// section: main
}

// section: stringcat
func Cat(s1, s2 string) string {
	return s1 + s2
}

// section: stringcat

// section: bytecat
func BCat(b1, b2 []byte) []byte {
	return append(b1, b2...)
}

// section: bytecat

/*
// section: output
./immutable.go:16:6: can inline Cat
./immutable.go:23:6: can inline BCat
./immutable.go:7:10: inlining call to Cat
./immutable.go:8:13: inlining call to fmt.Println
./immutable.go:10:11: inlining call to BCat
./immutable.go:11:13: inlining call to fmt.Println
/var/folders/l7/3s7z7s1s4n72lvj4w6g_fdmm0000gn/T/go-build200666631/b001/_gomod_.go:6:6: can inline init.0
./immutable.go:8:13: s escapes to heap
./immutable.go:7:10: s1 + s2 escapes to heap
./immutable.go:8:13: io.Writer(os.Stdout) escapes to heap
./immutable.go:11:13: io.Writer(os.Stdout) escapes to heap
./immutable.go:11:20: string(b) escapes to heap
./immutable.go:11:20: string(b) escapes to heap
./immutable.go:8:13: main []interface {} literal does not escape
./immutable.go:10:18: main ([]byte)("hello ") does not escape
./immutable.go:10:36: main ([]byte)("world") does not escape
./immutable.go:11:13: main []interface {} literal does not escape
./immutable.go:17:12: s1 + s2 escapes to heap
./immutable.go:16:10: Cat s1 does not escape
./immutable.go:16:14: Cat s2 does not escape
./immutable.go:23:11: leaking param: b1 to result ~r2 level=0
./immutable.go:23:15: BCat b2 does not escape
// section: output


*/
