package main

import "fmt"

func main() {
	m := []string{
		"one", "two", "three",
	}
	fmt.Println(m)
	m = push(m, "1", "2", "3")
	fmt.Println(m)
}

func push(a []string, v ...string) (s []string) {
	// 可変長引数（制約として、可変長引数は最後の位置の変数でのみ使用可）
	s = append(a, v...)
	return
}
