package main

import "fmt"

// Mydata is structure
type Mydata struct {
	Name string
	Data []int
}

func main() {
	taro := Mydata {
		"Taro",
		[]int{10, 20, 30},
	}
	fmt.Println(taro) // {Taro [10 20 30]}
	taro = reverse(taro) // 関数の引数で用いると値渡し（別アドレスでメモリ領域を確保して値コピー）
	fmt.Println(taro) // {Taro [30 20 10]}
}

func reverse(md Mydata) Mydata {
	od := md.Data
	nd := []int{}
	for i := len(od)-1; i>=0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
	return md
}
