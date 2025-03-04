package main

import (
	"fmt"
	"runtime"
)

func main() {
	var p *int
	//*p = 5 //panic: runtime error: invalid memory address or nil pointer dereference
	x := 10
	p = &x
	fmt.Println(p)
	fmt.Println(*p) // To print the dereference value
	*p = 100
	fmt.Println(x)

	p1 := new(int)
	fmt.Println(p1)
	fmt.Println(*p1)
	*p1 = 42
	fmt.Println(*p1)

	n := make([]int, 5)
	fmt.Println(n)

	n[0] = 10
	fmt.Println(n)

	ptr := createPointerS()
	fmt.Println(*ptr)

	pt := createPointer()
	fmt.Println(pt)

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	memoryUsage(memStats)

	_ = make([]int, 100000000)
	memoryUsage(memStats)
	runtime.GC()

	_ = make([]int, 100000000)
	memoryUsage(memStats)
	runtime.GC()
	memoryUsage(memStats)
	memoryLeak()

	for i := 0; i < 1000000; i++ {
		leakGoroutine() // Starting too many goroutines without closing channels
	}
}

func leakGoroutine() {
	ch := make(chan int)
	go func() {
		// Goroutine is blocked forever because channel is not closed
		for {
			_, open := <-ch
			if !open {
				break
			}
		}
	}()
	// Forgetting to close the channel causes a memory leak, use close(ch)
}

func memoryLeak() {
	largeSlice := make([]byte, 10<<30)
	fmt.Println(len(largeSlice))
	largeSlice = nil // Dereference the largeSlice to allow garbage collection
}

func memoryUsage(memStats runtime.MemStats) {
	fmt.Printf("Alloc = %v  || TotalAlloc = %v  || Sys = %v  || NumGC = %v\n",
		memStats.Alloc,
		memStats.TotalAlloc,
		memStats.Sys,
		memStats.NumGC)
}

func createPointerS() *int { //Heap memory allocation
	x := 42
	return &x
}

func createPointer() int { //Stack memory allocation
	x := 42
	return x
}
