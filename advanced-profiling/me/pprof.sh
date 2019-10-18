#!/usr/bin/env bash

# Profile CPU
go test . -bench=. -v -cpuprofile=cpu.pprof
go tool pprof cpu.pprof

# Profile memory
go test . -bench=. -v -memprofile=mem.pprof
go tool pprof --alloc_objects ./src/mem.pprof
go tool pprof --inuse_objects ./src/mem.pprof
