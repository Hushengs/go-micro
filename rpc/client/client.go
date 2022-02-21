package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//声明参数结构体
type Params struct {
	Width, Height int
}

func main() {
	//1.连接远程rpc服务
	rp, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal()
	}
	//2.调用远程方法
	//定义接收服务端传回来的计算结果的变量
	ret := 0
	//求面积
	err1 := rp.Call("Rect.Area", Params{50, 100}, &ret)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println("面积：", ret)
	//求周长
	err2 := rp.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err1 != nil {
		log.Fatal(err2)
	}
	fmt.Println("周长：", ret)
}
