package main

import (
	"log"
	"net/http"
	"net/rpc"
)

//go对RPC的支持，支持三个级别：TCP、HTTP、JSONRPC
//go的RPC只支持GO开发的服务器与客户端之间的交互，因为采用了gob编码

//注意字段必须是导出
type ParamsSer struct {
	Width, Height int
}

type Rect struct{}

//函数必须是导出的
//必须有两个导出类型参数
//第一个参数是接收参数
//第二个参数是返回给客户端参数，必须是指针类型
//函数还要有一个返回值error
func (r *Rect) Area(p ParamsSer, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

func (r *Rect) Perimeter(p ParamsSer, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}

func main() {

	rect := new(Rect)
	//注册一个rect服务
	rpc.Register(rect)
	//把服务处理绑定到http协议上
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		log.Fatal(err)
	}
}
