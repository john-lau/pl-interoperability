//usr/bin/env go run "$0" "$@"; exit

package main

import (
   "flag"
   "fmt"
   // "github.com/pkg/profile"
   // "math"
   "strconv"
   // "time"
   "runtime"
)

// var n = flag.Int("n", 2000, "count")
var n = 0

func PrintMemUsage() {
  var m runtime.MemStats
  runtime.ReadMemStats(&m)
  // For info on each, see: https://golang.org/pkg/runtime/#MemStats
  fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
  fmt.Printf("\tTotalAlloc = %v ", m.TotalAlloc)
  fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
  fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func evalA(i, j int) int { return ((i+j)*(i+j+1)/2 + i + 1) }

type Vec []float64

func (v Vec) Times(u Vec) {
   for i := 0; i < len(v); i++ {
      v[i] = 0
      for j := 0; j < len(u); j++ {
         v[i] += u[j] / float64(evalA(i, j))
      }
   }
}

func (v Vec) TimesTransp(u Vec) {
   for i := 0; i < len(v); i++ {
      v[i] = 0
      for j := 0; j < len(u); j++ {
         v[i] += u[j] / float64(evalA(j, i))
      }
   }
}

func (v Vec) ATimesTransp(u Vec) {
   x := make(Vec, len(u))
   x.Times(u)
   v.TimesTransp(x)
}

func bToMb(b uint64) uint64 {
  return b / 1024 / 1024
}

func main() {
   // CPU profiling by default
   // defer profile.Start(profile.MemProfile).Stop()
   // var m runtime.MemStats
   // runtime.ReadMemStats(&m)
   // start := time.Now()
   flag.Parse()
   if flag.NArg() > 0 {
     n, _ = strconv.Atoi(flag.Arg(0))
   }

   N := n
   u := make(Vec, N)
   for i := 0; i < N; i++ {
      u[i] = 1
   }
   v := make(Vec, N)
   for i := 0; i < 10; i++ {
      v.ATimesTransp(u)
      u.ATimesTransp(v)
   }
   var vBv, vv float64
   for i := 0; i < N; i++ {
      vBv += u[i] * v[i]
      vv += v[i] * v[i]
      // PrintMemUsage()
   }
   // fmt.Println(float32(time.Since(start))/1000000000.0)
   // fmt.Printf("%0.9f\n", math.Sqrt(vBv/vv))
   var m runtime.MemStats
   runtime.ReadMemStats(&m)
   fmt.Println((m.Sys + m.TotalAlloc)/1024.0)
}
