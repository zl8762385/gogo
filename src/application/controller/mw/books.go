package controllers

import (
	"fmt"
	"frame"
)

// 继承鸡肋
type Books struct {
	frame.Controller
}


func (b *Books) Infos() {
	fmt.Fprint(b.Response, "书的详细信息")
}
