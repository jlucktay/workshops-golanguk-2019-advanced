// +build mage

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
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
	cleanPatterns := []string{"bench.*.log", "cpu.*.pprof", "mem.*.pprof"}
	needsCleaning := make([]string, 0)
	for _, pattern := range cleanPatterns {
		matches, errGlob := filepath.Glob(pattern)
		if errGlob != nil {
			return errGlob
		}
		needsCleaning = append(needsCleaning, matches...)
	}

	for _, clean := range needsCleaning {
		if errRemove := os.Remove(clean); errRemove != nil {
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

// timestamp returns the current time in 'YYYYMMDD.hhmmss.SSSSSS±zzzz' format, where:
// Y = year
// M = month
// D = date
// h = hours (24, not 12)
// m = minutes
// s = seconds
// S = microseconds
// z = UTC offset
//
// All fields are zero-padded where appropriate.
func timestamp() string {
	return time.Now().Format("20060102.150405.000000-0700")
}

// Runs all benchmarks with memory allocation statistics 10 times, and saves the output.
// go test -benchmem -bench=. -count=10 > old.bench
func SaveBench() error {
	out, err := sh.Output("go", "test", srcDir, "-benchmem", "-bench=.", "-count=10")
	if err != nil {
		return err
	}

	benchFile := fmt.Sprintf("bench.%s.log", timestamp())
	return ioutil.WriteFile(benchFile, []byte(out), 0644)
}

// Runs 'benchstat' to compare the two most recent benchmark logs.
// benchstat old.bench new.bench
func CompareBench() error {
	matches, errGlob := filepath.Glob("bench.*.log")
	if errGlob != nil {
		return errGlob
	}

	if len(matches) < 2 {
		return errors.New("there must be at least 2 'bench.*.log' files in this directory to compare to one another")
	}

	type fileMod struct {
		name    string
		modTime time.Time
	}

	fileMods := make([]fileMod, len(matches))

	for i, match := range matches {
		fi, errStat := os.Stat(match)
		if errStat != nil {
			return errStat
		}

		fileMods[i] = fileMod{
			name:    match,
			modTime: fi.ModTime(),
		}
	}

	sort.Slice(fileMods, func(i, j int) bool {
		return fileMods[i].modTime.Before(fileMods[j].modTime)
	})

	return sh.RunV("benchstat", fileMods[len(fileMods)-2].name, fileMods[len(fileMods)-1].name)
}

// Runs the benchmarks, generates a CPU profile, and opens it in the interactive profiler.
// go test -bench=. -cpuprofile=cpu.pprof ; pprof -http=:8181 cpu.pprof
func CpuProfile() error {
	proFile := fmt.Sprintf("cpu.%s.pprof", timestamp())
	if errTest := sh.Run("go", "test", srcDir, "-bench=.", "-cpuprofile="+proFile); errTest != nil {
		return errTest
	}

	return sh.RunV("pprof", "-http=:8181", proFile)
}

// Runs the benchmarks, generates a memory profile, and opens it in the interactive profiler, indexing allocated
// objects.
// go test -bench=. -memprofile=mem.pprof ; pprof -sample_index=alloc_objects -http=:8181 mem.pprof
func MemProfileAlloc() error {
	proFile := fmt.Sprintf("mem.%s.pprof", timestamp())
	if errTest := sh.Run("go", "test", srcDir, "-bench=.", "-memprofile="+proFile); errTest != nil {
		return errTest
	}

	return sh.RunV("pprof", "-sample_index=alloc_objects", "-http=:8181", proFile)
}

// Runs the benchmarks, generates a memory profile, and opens it in the interactive profiler, indexing in-use objects.
// go test -bench=. -memprofile=mem.pprof ; pprof -sample_index=inuse_objects -http=:8181 mem.pprof
func MemProfileInUse() error {
	proFile := fmt.Sprintf("mem.%s.pprof", timestamp())
	if errTest := sh.Run("go", "test", srcDir, "-bench=.", "-memprofile="+proFile); errTest != nil {
		return errTest
	}

	return sh.RunV("pprof", "-sample_index=inuse_objects", "-http=:8181", proFile)
}
