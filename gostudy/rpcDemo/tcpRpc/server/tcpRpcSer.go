package main

import (
	"log"
	"net"
	"net/rpc"
)

//注意字段必须是导出
type Params struct {
	Width, Height int
}

type Rect struct {
}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil

}

func chkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	rect := new(Rect)
	//注册rpc服务
	rpc.Register(rect)
	//获取tcpaddr
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8082")
	chkError(err)

	//监听端口
	tcplisten, err2 := net.ListenTCP("tcp", tcpaddr)
	chkError(err2)
	//死循环处理连接请求
	for {
		conn, err3 := tcplisten.Accept()
		if err3 != nil {
			continue
		}
		//使用goroutine单独处理rpc连接请求
		go rpc.ServeConn(conn)
	}
}
