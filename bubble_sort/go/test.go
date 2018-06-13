//usr/bin/env go run "$0" "$@"; exit

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	// "time"
)

func swap(arrayzor []int, i, j int) {
	tmp := arrayzor[j]
	arrayzor[j] = arrayzor[i]
	arrayzor[i] = tmp
}

func bubbleSort(arrayzor []int) {

	swapped := true;
	for swapped {
		swapped = false
		for i := 0; i < len(arrayzor) - 1; i++ {
			if arrayzor[i + 1] < arrayzor[i] {
				swap(arrayzor, i, i + 1)
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
	// start := time.Now()
	// CPU profiling by default
	// defer profile.Start(profile.MemProfile).Stop()

	// arrayzor := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}

	//generate random array
	n := int(100)
	flag.Parse()
	if flag.NArg() > 0 {
		n, _ = strconv.Atoi(flag.Arg(0))
	}

	arrayzor := rangeInt(0, 1000, n)

	// fmt.Println("Unsorted array: ", arrayzor)
	bubbleSort(arrayzor)
	// fmt.Println("Sorted array: ", arrayzor)

	// fmt.Println(time.Since(start))

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println((m.Sys + m.TotalAlloc)/1024.0)
}
