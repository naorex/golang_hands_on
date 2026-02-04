package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Go-lang!")
	fmt.Println(123 * 45)

	fmt.Print("123 * 45 = ") // 出力を改行しない
	fmt.Println(123 * 45)
}
