package main

import (
	"fmt"
	"hello"
)

func main() {
	name := hello.Input("type your name") // := 変数の定義と値の代入を同時に
	fmt.Println("Hello, " + name + "!!")
}
