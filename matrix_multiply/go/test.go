//usr/bin/env go run "$0" "$@"; exit

package main

import (
	"fmt"
	"flag"
	"runtime"
	"strconv"
	// "time"
)

func matgen(n int) [][]float64 {
	a := make([][]float64, n)
	tmp := float64(1.0) / float64(n) / float64(n) //this is not the smartest
	for i := 0; i < n; i++ {
		a[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			a[i][j] = tmp * float64(i-j) * float64(i+j)
		}
	}
	return a
}

func matmul(a [][]float64, b [][]float64){
	m := len(a)
	n := len(a[0])
	p := len(b[0])
	x := make([][]float64, m)
	c := make([][]float64, p)
	for i := 0; i < p; i++ {
		c[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			c[i][j] = b[j][i]
		}
	}
	for i, am := range a {
		x[i] = make([]float64, p)
		for j, cm := range c {
			s := float64(0)
			for k, m := range am {
				s += m * cm[k]
			}
			x[i][j] = s
		}
	}
	// return x
}

func main() {
	// start := time.Now()
	// CPU profiling by default
	// defer profile.Start(profile.MemProfile).Stop()
	n := int(100)
	flag.Parse()
	if flag.NArg() > 0 {
		n, _ = strconv.Atoi(flag.Arg(0))
	}
	a := matgen(n)
	b := matgen(n)
	matmul(a, b)
	// fmt.Printf("%f\n", x[n/2][n/2])

	// fmt.Println(time.Since(start))

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println((m.Sys + m.TotalAlloc)/1024.0)
}
