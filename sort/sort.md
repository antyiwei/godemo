
# 使用 Go 实现快速排序

## [转载来自`使用 Go 实现快速排序`](http://colobu.com/2018/06/26/implement-quick-sort-in-golang/#more)

快速排序(quick sort)号称是二十世纪最伟大的十大算法之一(`The Best of the 20th Century: Editors Name Top 10 Algorithms`), 但是快速排序也是最不容易实现的排序算法之一 ()。虽然它的原理非常的简单，但实现起来很容易出错。 也曾因为快排导致腥风血雨甚至网站攻击事件。

快速排序由C. A. R. Hoare在1962年提出。它的基本思想是：通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列。
- - 分治法：将问题分解为若干个规模更小但结构与原问题相似的子问题。递归地解这些子问题，然后将这些子问题的解组合为原问题的解。
利用分治法可将快速排序的分为三步：

* 在数据集之中，选择一个元素作为”基准”（pivot）。
* 所有小于”基准”的元素，都移到”基准”的左边；所有大于”基准”的元素，都移到”基准”的右边。这个操作称为分区 (partition) 操作，分区操作结束后，基准元素所处的位置就是最终排序后它的位置。
* 对”基准”左边和右边的两个子集，不断重复第一步和第二步，直到所有子集只剩下一个元素为止。
快速排序平均时间复杂度为O(n log n),最坏情况为O(n2)，不稳定排序。

快速排序一般实现为原地排序(in-place),因为非原地排序会设计到大量的容器创建和对象复制。

本文实现了两种快速排序，一种是单线程的快速排序,一种是一定数量的goroutine并行的快速排序。

同时也增加了标准库排序算法和timsort算法的比较。

下面是算法实现：
```go
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
```