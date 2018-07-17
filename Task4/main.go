package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	threads := 8
	for i := 0; i < threads; i++ {
		wg.Add(1)
		start := i + 1
		go func() {
			find(start, threads)
			wg.Done()
		}()
	}
	// Wait until all threads are done
	wg.Wait()
}


func find(start, step int) {
	for i := start; i < 10000; i += step {
		for j := start + 100000; i < 100000*10; j++ {
			num := palindrom(i, i+j)
			if isPalindromic(num) {
				fmt.Println(i, j, ":", num)
			}
		}
	}
}

func palindrom(start, end int) int {
	return squareSum(end) - squareSum(start-1)
}

func squareSum(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func isPalindromic(num int) bool {
	s := strconv.Itoa(num)
	return s == reverse(s)
}

func reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}
