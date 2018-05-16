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
	"strconv"
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

func matmul(a [][]float64, b [][]float64) [][]float64 {
	m := len(a)
	n := len(a[0])
	p := len(b[0])
	x := make([][]float64, m)
	c := make([][]float64, p)
	for i := 0; i < p; i++ { //transpose b
		c[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			c[i][j] = b[j][i]
		}
	}
	for i, am := range a { //do the multiplication, store values in x
		x[i] = make([]float64, p)
		for j, cm := range c {
			s := float64(0)
			for k, m := range am {
				s += m * cm[k]
			}
			x[i][j] = s
		}
	}
	return x
}

func main() {
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

}
