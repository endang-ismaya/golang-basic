package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"endangismaya.com/gotest/utils"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 started")

	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		time.Sleep(time.Second * 2)
	}
	fmt.Println("f1 exited")
	wg.Done()
}

func f2() {
	fmt.Println("f2 started")

	for i := 5; i < 8; i++ {
		fmt.Println("f2, i=", i)
	}
	fmt.Println("f2 exited")
}

func cpuInfo() {
	utils.PrintTitle("CPU Information")
	fmt.Println("main started")
	fmt.Println("No. of CPUs", runtime.NumCPU())
	fmt.Println("No. of GOroutines", runtime.NumGoroutine())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("Go MAx Procs:", runtime.GOMAXPROCS(0))
}

func main() {
	// run cpu infomation
	cpuInfo()

	// create var for WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)

	utils.PrintTitle("Running f1()")
	go f1(&wg)
	fmt.Println("No. of GOroutines after f1():", runtime.NumGoroutine())

	f2()

	// exiting main
	wg.Wait()
	fmt.Println("main stopped")
}
