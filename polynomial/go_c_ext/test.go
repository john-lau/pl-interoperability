//usr/bin/env go run "$0" "$@"; exit

package main

/*
#include "lib/lib.h"
#include <stdlib.h>
#cgo LDFLAGS: -L. -llib */
import "C"

import (
	"flag"
	"fmt"
	// "github.com/emilymaier/cmemory"
	"strconv"
	// "time"
  // "runtime"
)

func main() {
	// start := time.Now()
	// cmemory.StartInstrumentation()
	flag.Parse()
	n, _ := strconv.Atoi(flag.Arg(0))
	x, _ := strconv.ParseFloat(flag.Arg(1), 64)
	pu := 0.0
	x_c := C.float(x)
	for i := 0; i < n; i++ {
		pu += float64(C.poly(x_c))
	}
	// fmt.Println(float32(time.Since(start)) / 1000000000.0)
	fmt.Printf("%f\n", pu)
	// stats := cmemory.MemoryAnalysis()
	// fmt.Println(stats.TotalBytesAllocated)
  // var m runtime.MemStats
  // runtime.ReadMemStats(&m)
  // fmt.Println((m.Sys + m.TotalAlloc + 16000)/1024.0)
}
