package main

import "fmt"

func main() {
	m := []string{}
	m, _ = push(m, "apple")
	m, _ = push(m, "banana")
	m, _ = push(m, "orange")
	fmt.Println(m)
	m, v := pop(m)
	fmt.Println("get " + v + " ->", m)
}

func push(a []string, v string) ([]string, int) {
	// 最後に追加
	return append(a, v), len(a)
}

func pop(a []string) ([]string, string) {
	// 最後を削除
	return a[:len(a)-1], a[len(a)-1]
}
