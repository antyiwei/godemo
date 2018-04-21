package main

import "fmt"

// 小心append slice
func main() {

	s0 := []int{1, 2, 3, 4}
	s1 := s0[:2]
	s2 := append(s1, 5)
	s3 := append(s1, 6, 7)
	s4 := append(s1, 8, 9, 10)
	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	println(s0)
	println(s1)
	println(s2)
	println(s3)
	println(s4)

}
