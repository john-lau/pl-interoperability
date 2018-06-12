//usr/bin/env go run "$0" "$@"; exit

package main

import (
   "flag"
   "fmt"
   "strconv"
   // "time"
   // "runtime"
)


type Vec []float64


func poly(x float64) float64 {
  mu := 10.0
  pol := make(Vec, 100)
  for i := 0; i < 100; i++ {
    mu = (mu + 2.0)/2.0
    pol[i] = (mu + 2.0)/2.0
  }
  su := 0.0
  for i := 0; i < 100; i++ {
    su = x*su + pol[i]
  }
  return su
}

func main() {
  // start := time.Now()
  flag.Parse()
  n, _ := strconv.Atoi(flag.Arg(0))
  x, _ := strconv.ParseFloat(flag.Arg(1), 64)
  pu := 0.0
  for i := 0; i < n; i++ {
    pu += poly(x)
  }
  fmt.Printf("%f\n", pu)
	// fmt.Println(float32(time.Since(start)) / 1000000000.0)
   // var m runtime.MemStats
   // runtime.ReadMemStats(&m)
   // fmt.Println((m.Sys + m.TotalAlloc)/1024.0)
}
