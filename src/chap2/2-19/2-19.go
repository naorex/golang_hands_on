package main

import "fmt"

func main() {
	a := []int {10,20,30}
	fmt.Println(a) // [10 20 30]
	a = push(a, 1000)
	fmt.Println(a) // [10 20 30 1000]
	a = pop(a)
	fmt.Println(a) // [10 20 30]
	a = unshift(a, 1000)
	fmt.Println(a) // [1000 10 20 30]
	a = shift(a)
	fmt.Println(a) // [10 20 30]
	a = insert(a, 1000, 2)
	fmt.Println(a) // [10 20 1000 30]
	a = remove(a, 2)
	fmt.Println(a) // [10 20 30]
}

// スライスに関する機能は append のみなので、以下で実装

func push(a []int, v int) []int {
	// 最後に追加
	return append(a, v)
}

func pop(a []int) []int {
	// 最後を削除
	return a[:len(a)-1]
}

func unshift(a []int, v int) []int {
	// 最初に追加
	return append([]int{v}, a...) // a... => a[0], a[1], a[2]
}

func shift(a []int) []int {
	// 最初を削除
	return a[1:]
}

func insert(a []int, v int, p int) []int {
	// 指定の位置に追加
	a = append(a, 0)
	a = append(a[:p+1], a[p:len(a)-1]...)
	a[p] = v
	return a
}

func remove(a []int, p int) []int {
	// 指定の位置を削除
	return append(a[:p], a[p+1:]...)
}
