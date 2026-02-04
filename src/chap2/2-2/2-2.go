package main

import (
	"fmt"
	"hello"
	"strconv"
)

func main() {
	x := hello.Input("type a price")
	n, err := strconv.Atoi(x) // "ASCII to integer" メソッド（文字列を数値に変換）

	if err != nil {
		fmt.Println("ERROR!!")
		return
	}

	p := float64(n)
	fmt.Println(int(p * 1.1))
}
