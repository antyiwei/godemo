package main

import "fmt"



// 遍历slice时区分用值还是索引
func main() {
	arr:=[]struct{v int }{{1},{2},{3}}

	for _,e :=range arr{
		e.v++
	}
	fmt.Printf("%v\n",arr )

	for idx:= range arr {
		arr[idx].v++

	}
	fmt.Printf("%v\n",arr)
}