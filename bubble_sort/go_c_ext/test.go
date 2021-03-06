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
	"math/rand"
	// "runtime"
	"strconv"
	// "time"
	"unsafe"
)

func swap(arrayzor []int, i, j int) {
	tmp := arrayzor[j]
	arrayzor[j] = arrayzor[i]
	arrayzor[i] = tmp
}

func bubbleSort(arrayzor []int) {

	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arrayzor)-1; i++ {
			if arrayzor[i+1] < arrayzor[i] {
				swap(arrayzor, i, i+1)
				swapped = true
			}
		}
	}
}

func rangeInt(min int, max int, n int) []int {
	arr := make([]int, n)
	var r int
	for r = 0; r <= n-1; r++ {
		arr[r] = rand.Intn(max) + min
	}
	return arr
}

func main() {
	// cmemory.StartInstrumentation()
	// start := time.Now()

	//generate random array
	n := int(100)
	flag.Parse()
	if flag.NArg() > 0 {
		n, _ = strconv.Atoi(flag.Arg(0))
	}

	N := C.int(n)

	goArray := rangeInt(0, 1000, n)

	fmt.Println("Unsorted array: ", goArray)

	// Create C array from golang array
	cArray := (C.malloc(C.size_t(N) * C.size_t(unsafe.Sizeof(C.int(0)))))

	arraySlice := (*[1<<30 - 1]C.int)(cArray)[:n:n]

	for i := 0; i < n; i++ {
		arraySlice[i] = C.int(goArray[i])
	}

	// Call C function now.
	C.bubble_sort((*C.int)(cArray), N)

	fmt.Println("Sorted array: ", arraySlice)

	// fmt.Println(time.Since(start))

	// stats := cmemory.MemoryAnalysis()
	// fmt.Println((stats.TotalBytesAllocated / 1024.0))

	// var m runtime.MemStats
	// runtime.ReadMemStats(&m)
	// fmt.Println((m.Sys + m.TotalAlloc) / 1024.0)
}
