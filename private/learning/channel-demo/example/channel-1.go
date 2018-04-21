package main

import "fmt"

// 操作空channel 会死锁而非panic
func main() {
	var c chan string
	fmt.Println(<-c) // deadlocks
}
