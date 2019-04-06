package main

import (
	"application/controller/mw"
	"frame"
)

func main() {
	frame := frame.Default()
	frame.GET("act/foo", &controllers.Act{})
	frame.GET("act/index", &controllers.Act{})
	frame.GET("books/infos", &controllers.Books{})
	frame.RUN(":80")
}