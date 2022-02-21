package main

import (
	"log"
	"net/http"
	"net/rpc"
)

//声明矩形对象
type Rect struct {
}

//声明参数结构体
type Params struct {
	Width, Height int
}

//定义求矩形面积的方法
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

//定义求矩形面积的方法
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}

func main() {
	rect := new(Rect)
	rpc.Register(rect)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
