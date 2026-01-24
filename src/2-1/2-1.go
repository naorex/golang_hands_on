package main

import "fmt"

func main() {
	var a, b, c int = 100, 200, 300
	// a, b, c:= 100, 200, 300 // 短縮表記（暗黙的な型宣言）
	fmt.Print("total: ")
	fmt.Println(a+b+c)
}
