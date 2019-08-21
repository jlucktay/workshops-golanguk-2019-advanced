// +build mage

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

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

// Clean up various generated output files.
func Clean() error {
	matches, errGlob := filepath.Glob("bench.*.log")
	if errGlob != nil {
		return errGlob
	}

	for _, match := range matches {
		if errRemove := os.Remove(match); errRemove != nil {
			return errRemove
		}
	}

	return nil
}

// Runs all benchmarks with memory allocation statistics.
// go test -benchmem -bench=.
func Benchmark() error {
	return sh.RunV("go", "test", srcDir, "-benchmem", "-bench=.")
}

// Runs all benchmarks with memory allocation statistics 10 times, and saves the output.
// go test -benchmem -bench=. -count=10 > old.bench
func SaveBench() error {
	out, err := sh.Output("go", "test", srcDir, "-benchmem", "-bench=.", "-count=10")
	if err != nil {
		return err
	}

	now := time.Now().Format("20060102.150405.000000-0700")
	benchFile := fmt.Sprintf("bench.%s.log", now)
	return ioutil.WriteFile(benchFile, []byte(out), 0644)
}
