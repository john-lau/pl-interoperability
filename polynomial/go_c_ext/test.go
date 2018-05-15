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
	"strconv"
)

// type Vec []float64

// func poly(x float64) float64 {
//   mu := 10.0
//   pol := make(Vec, 100)
//   for i := 0; i < 100; i++ {
//     mu = (mu + 2.0)/2.0
//     pol[i] = (mu + 2.0)/2.0
//   }
//   su := 0.0
//   for i := 0; i < 100; i++ {
//     su = x*su + pol[i]
//   }
//   return su
// }

func main() {
	flag.Parse()
	n, _ := strconv.Atoi(flag.Arg(0))
	x, _ := strconv.ParseFloat(flag.Arg(1), 64)
	pu := 0.0
	x_c := C.float(x)
	for i := 0; i < n; i++ {
		pu += float64(C.poly(x_c))
	}
	fmt.Printf("%f\n", pu)
}
