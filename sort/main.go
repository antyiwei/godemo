package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/psilva261/timsort"
)

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}
func quickSort_go(a []int, lo, hi int, done chan struct{}, depth int) {
	if lo >= hi {
		done <- struct{}{}
		return
	}
	depth--
	p := partition(a, lo, hi)
	if depth > 0 {
		childDone := make(chan struct{}, 2)
		go quickSort_go(a, lo, p-1, childDone, depth)
		go quickSort_go(a, p+1, hi, childDone, depth)
		<-childDone
		<-childDone
	} else {
		quickSort(a, lo, p-1)
		quickSort(a, p+1, hi)
	}
	done <- struct{}{}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	testData1, testData2, testData3, testData4 := make([]int, 0, 100000000), make([]int, 0, 100000000), make([]int, 0, 100000000), make([]int, 0, 100000000)
	times := 100000000
	for i := 0; i < times; i++ {
		val := rand.Intn(20000000)
		testData1 = append(testData1, val)
		testData2 = append(testData2, val)
		testData3 = append(testData3, val)
		testData4 = append(testData4, val)
	}
	start := time.Now()
	quickSort(testData1, 0, len(testData1)-1)
	fmt.Println("single goroutine: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData1) {
		fmt.Println("wrong quick_sort implementation")
	}
	done := make(chan struct{})
	start = time.Now()
	go quickSort_go(testData2, 0, len(testData2)-1, done, 5)
	<-done
	fmt.Println("multiple goroutine: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData2) {
		fmt.Println("wrong quickSort_go implementation")
	}
	start = time.Now()
	sort.Ints(testData3)
	fmt.Println("std lib: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData3) {
		fmt.Println("wrong std lib implementation")
	}
	start = time.Now()
	timsort.Ints(testData4, func(a, b int) bool { return a <= b })
	fmt.Println("timsort: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData4) {
		fmt.Println("wrong timsort implementation")
	}
}
