package main

import "fmt"

func main() {
	// 無名関数
	// 先頭から順に取り出すモジュールをここでは定義
	f := func(a []string) ([]string, string) {
		return a[1:], a[0]
	}
	m := []string{
		"one", "two", "three",
	}
	s := ""
	fmt.Println(m)
	for len(m) > 0 {
		m, s = f(m)
		fmt.Println(s + " ->", m)
	}
}
