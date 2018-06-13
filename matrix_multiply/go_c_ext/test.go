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
	// "runtime"
	"strconv"
	// "time"
	"unsafe"
)

func matgen(n int) [][]float64 {
	a := make([][]float64, n)
	tmp := float64(1.0) / float64(n) / float64(n)
	for i := 0; i < n; i++ {
		a[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			a[i][j] = tmp * float64(i-j) * float64(i+j)
		}
	}
	return a
}

func main() {
	// cmemory.StartInstrumentation()
	// start := time.Now()

	// CPU profiling by default
	// defer profile.Start(profile.MemProfile).Stop()

	n := int(100)
	flag.Parse()
	if flag.NArg() > 0 {
		n, _ = strconv.Atoi(flag.Arg(0))
	}

	N := C.int(n)

	a := matgen(n)
	b := matgen(n)

	//Create an array of pointers of type C float, of size n
	a_matrix := make([]*C.float, n)
	b_matrix := make([]*C.float, n)

	//for each "row" in matrix a
	for i, row := range a {
		//create c array of length n
		a_row := (*C.float)(C.malloc(C.size_t(N) * C.size_t(unsafe.Sizeof(C.float(0)))))
		a_matrix[i] = a_row
		//convert c array into golang array so we can index into it
		a_item := (*[1<<30 - 1]C.float)(unsafe.Pointer(a_row))
		for j, v := range row {
			(*a_item)[j] = C.float(v)
		}
	}

	for i, row := range b {
		//create c array of length n
		b_row := (*C.float)(C.malloc(C.size_t(N) * C.size_t(unsafe.Sizeof(C.float(0)))))
		b_matrix[i] = b_row
		//convert c array into golang array so we can index into it
		b_item := (*[1<<30 - 1]C.float)(unsafe.Pointer(b_row))
		for j, v := range row {
			(*b_item)[j] = C.float(v)
		}
	}

	result := C.matmul(N, (&a_matrix[0]), (&b_matrix[0]))

	fmt.Printf("%f\n", result)

	// fmt.Println(time.Since(start))

	// stats := cmemory.MemoryAnalysis()
	// fmt.Println((stats.TotalBytesAllocated / 1024.0))

	// var m runtime.MemStats
	// runtime.ReadMemStats(&m)
	// fmt.Println((m.Sys + m.TotalAlloc) / 1024.0)

}
