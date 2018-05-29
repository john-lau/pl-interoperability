//usr/bin/env go run "$0" "$@"; exit

package main

/*
#include "lib/lib.h"
#include <stdlib.h>
#cgo LDFLAGS: -L. -llib
*/
import "C"

import (
	"flag"
	"fmt"
	// "github.com/emilymaier/cmemory"
	// "time"
	// "github.com/pkg/profile"
	// "math"
	// "os"
	"runtime"
	"strconv"
	"unsafe"
)

func main() {
	// CPU profiling by default
	// defer profile.Start(profile.MemProfile).Stop()
	// cmemory.StartInstrumentation()
	flag.Parse()
	n := 10
	if flag.NArg() > 0 {
		n, _ = strconv.Atoi(flag.Arg(0))
	}

	N := C.int(n)

	u := (C.malloc(C.size_t(N) * C.size_t(unsafe.Sizeof(C.double(0)))))
	u_slice := (*[1<<30 - 1]C.double)(u)[:n:n]

	for i := 0; i < n; i++ {
		u_slice[i] = C.double(1.0)
	}

	v := (C.malloc(C.size_t(N) * C.size_t(unsafe.Sizeof(C.double(0)))))
	v_slice := (*[1<<30 - 1]C.double)(v)[:n:n]

	for i := 0; i < 10; i++ {
		C.eval_AtA_times_u(N, (*C.double)(u), (*C.double)(v))
		C.eval_AtA_times_u(N, (*C.double)(v), (*C.double)(u))
	}
	var vBv, vv float64
	for i := 0; i < n; i++ {
		vBv += float64(u_slice[i]) * float64(v_slice[i])
		vv += float64(v_slice[i]) * float64(v_slice[i])
	}
	// fmt.Printf("%0.9f\n", math.Sqrt(vBv/vv))
	// stats := cmemory.MemoryAnalysis()
	// fmt.Println(stats.TotalBytesAllocated)
	 var m runtime.MemStats
	 runtime.ReadMemStats(&m)
  fmt.Println((m.Sys + m.TotalAlloc + 16000)/1024.0)
}
