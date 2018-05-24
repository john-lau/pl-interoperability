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
	"github.com/pkg/profile"
	"strconv"
)

func main() {
	// CPU profiling by default
	defer profile.Start(profile.MemProfile).Stop()
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
