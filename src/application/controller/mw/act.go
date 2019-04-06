package controllers

import (
	"fmt"
	"frame"
)

// 继承鸡肋
type Act struct {
	frame.Controller
}

// act
func (p *Act) Foo() {
	fmt.Fprint(p.Response, "foooooo")
	fmt.Println("Fooo")
}

// 实现方法 2
func (p *Act) Index() {
	fmt.Fprint(p.Response, "zhangxiaoliang")
}
