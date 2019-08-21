// +build mage

package main

import (
	"github.com/magefile/mage/sh"
)

const srcDir = "./src/"

// Default target to run when none is specified.
// If not set, running mage will list available targets.
// var Default = Build

// // A build step that requires additional params, or platform specific steps for example
// func Build() error {
// 	// mg.Deps(InstallDeps)
// 	fmt.Println("Building...")
// 	cmd := exec.Command("go", "build", "-o", "MyApp", ".")
// 	return cmd.Run()
// }

// // Clean up after yourself
// func Clean() {
// 	fmt.Println("Cleaning...")
// 	os.RemoveAll("MyApp")
// }

// go test -benchmem -bench=.
func Benchmark() error {
	return sh.RunV("go", "test", srcDir, "-benchmem", "-bench=.")
}

// go test -benchmem -bench=. -count=10 > old.bench
func Benchstat() error {
	_, err := sh.Output("go", "test", srcDir, "-benchmem", "-bench=.", "-count=10")
	if err != nil {
		return err
	}

	// sh.

	return nil
}
