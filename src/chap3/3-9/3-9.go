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
	reverse(&taro) // メモリアドレスの参照渡し
	fmt.Println(taro) // {Taro [30 20 10]}　戻り値が無い関数で操作したが、影響を受けた
}

func reverse(md *Mydata) {
	// 戻り値無し
	od := (*md).Data
	nd := []int{}
	for i := len(od)-1; i>=0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
}
