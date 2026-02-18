package main

import "fmt"

// "Mydata 型" を独自に定義して構造体を利用する
type Mydata struct {
	Name string
	Data []int
}

func main() {
	taro := Mydata{"Taro", []int{10, 20, 30}}
	hanako := Mydata{
		Name: "Hanako",
		Data: []int{90, 80, 70},
	}
	fmt.Println(taro) // {Taro [10 20 30]}
	fmt.Println(hanako) //{Hanako [90 80 70]}
}
