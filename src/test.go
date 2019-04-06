package main

import (
	"fmt"
	"reflect"
)

type W struct{}

func (t *W) Foo() {
	fmt.Println("foo")
}

func (t *W) Hoo() {
	fmt.Println("hoo")
}

func main() {
	var w W
	reflect.ValueOf(&w).MethodByName("Hoo").Call([]reflect.Value{})
}
