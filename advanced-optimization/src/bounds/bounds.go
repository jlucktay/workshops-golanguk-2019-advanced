package bounds

// section: forward
func forward(ints []int) int {
	s := ints[0]
	s = ints[1]
	s = ints[2]
	s = ints[3]
	s = ints[4]

	return s
}

// section: forward

// section: backwards
func backwards(ints []int) int {
	s := ints[4]
	s = ints[3]
	s = ints[2]
	s = ints[1]
	s = ints[0]

	return s
}

// section: backwards

// section: nocheck
func nocheck(ints []int, index int) int {
	if index >= 0 && index < len(ints) {
		return ints[index]
	}
	return 0
}

// section: nocheck

// section: range
func rangeBounds(ints []int) int {
	var s int
	for i := range ints {
		s += ints[i]
	}
	return s

}

// section: range

/*

// section: command
go build -gcflags="-d=ssa/check_bce/debug=1" ./bounds.go
// section: command

// section: command-prove
go build -gcflags="-d=ssa/prove/debug=1" ./bounds.go
// section: command-prove


// section: command-prove-2
go build -gcflags="-d=ssa/prove/debug=2" ./bounds.go
// section: command-prove-2

// section: command-ssafunc
GOSSAFUNC=forward go build
// section: command-ssafunc


// section: output
./bounds.go:5:11: Found IsInBounds
./bounds.go:6:10: Found IsInBounds
./bounds.go:7:10: Found IsInBounds
./bounds.go:8:10: Found IsInBounds
./bounds.go:9:10: Found IsInBounds
./bounds.go:18:11: Found IsInBounds
// section: output

// section: output-prove
./bounds.go:19:11: Proved IsInBounds
./bounds.go:20:11: Proved IsInBounds
./bounds.go:21:11: Proved IsInBounds
./bounds.go:22:11: Proved IsInBounds
./bounds.go:32:14: Proved IsInBounds
./bounds.go:42:11: Induction variable: limits [0,?), increment 1
./bounds.go:43:12: Proved IsInBounds
// section: output-prove


// section: output-prove-2
./bounds.go:19:11: Proved IsInBounds (v20)
./bounds.go:20:11: Proved IsInBounds (v30)
./bounds.go:21:11: Proved IsInBounds (v40)
./bounds.go:22:11: Proved IsInBounds (v49)
./bounds.go:32:14: Proved IsInBounds (v20)
./bounds.go:42:11: Induction variable: limits [0,v8), increment 1 (v9)
./bounds.go:43:12: Proved IsInBounds (v16)
// section: output-prove-2

// section: output-ssafunc
# github.com/gopherguides/learn/_training/advanced/optimization/src/bounds
dumped SSA to ./ssa.html
// section: output-ssafunc


// section: annotated-forward
5) s := ints[0] // Bounds Checked
6) s = ints[1]  // Bounds Checked
7) s = ints[2]  // Bounds Checked
8) s = ints[3]  // Bounds Checked
9) s = ints[4]  // Bounds Checked
// section: annotated-forward
// section: annotated-backwards
18) s := ints[4] // Bounds Checked
19) s = ints[3]  // Bounds Check Eliminated!
20) s = ints[2]  // Bounds Check Eliminated!
21) s = ints[1]  // Bounds Check Eliminated!
22) s = ints[0]  // Bounds Check Eliminated!
// section: annotated-backwards
*/
