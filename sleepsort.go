package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var numGoroutine int
var sortedInt []int

// SleepAndPrint prints a number of sleeping for n seconds
func SleepAndPrint(x int, wg *sync.WaitGroup) {
	defer wg.Done()

	// *Sleeping for time proportional to value
	time.Sleep(time.Duration(x) * time.Second)

	// *Printing the value
	fmt.Printf("%d ", x)

	// Append to sorted int
	sortedInt = append(sortedInt, x)
}

// SleepSort algorithm which takes integer slice
// as an argument
func SleepSort(numbers []int) {
	var wg sync.WaitGroup

	// *Creating wait group that waits of len(numbers) of go routines to finish
	wg.Add(len(numbers))

	for _, x := range numbers {
		// Preview of sleep time
		fmt.Printf("[%d]\tSleep time: ", x)
		for i := 1; i <= x; i++ {
			fmt.Printf(" ------- %ds", i)
		}
		fmt.Println()

		// *Spinning a Go routine
		go SleepAndPrint(x, &wg)

		numGoroutine++
	}

	// *Waiting for all go routines to finish
	wg.Wait()
}

func main() {
	args := os.Args[1:]
	var numbers = []int{}
	for _, x := range args {
		xi, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, xi)
	}

	startTimer := time.Now()
	SleepSort(numbers)
	endTimer := time.Since(startTimer)

	fmt.Println("\n\n======STATS======\n")
	fmt.Printf("1]Sorted array: %d\n", sortedInt)
	fmt.Printf("2]Num of goroutine spawned: %d (%d if main Goroutine is included)\n", numGoroutine, numGoroutine+1)
	fmt.Printf("3]Time spent running sleep sort: %s\n", endTimer)
}
