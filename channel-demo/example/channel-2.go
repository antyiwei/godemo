package main

import "fmt"

func main() {
	var c chan string
	c <- " let's get started  "

	fmt.Println(<-c)
}
