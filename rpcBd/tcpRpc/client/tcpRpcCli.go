package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

func chkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	//连接远程rpc服务
	//这里使用Dial，http方式使用DialHTTP，其他代码都一样
	rpc, err := rpc.Dial("tcp", "127.0.0.1:8082")
	chkError(err)

	ret := 0
	//调用远程方法
	//注意第三个参数是指针类型
	err2 := rpc.Call("Rect.Area", Params{50, 100}, &ret)
	chkError(err2)
	fmt.Println(ret)

	err3 := rpc.Call("Rect.Perimeter", Params{50, 100}, &ret)
	chkError(err3)
	fmt.Println(ret)
}
