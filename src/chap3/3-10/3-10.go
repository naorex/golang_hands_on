package main

import "fmt"

// Mydata is structure
type Mydata struct {
	Name string
	Data []int
}

func main() {
	taro := new(Mydata)
	fmt.Println(taro) // &{ []} ポインタ
	taro.Name = "Taro"
	taro.Data = make([]int, 5)
	fmt.Println(taro) // &{Taro [0 0 0 0 0]} ポインタ
}
