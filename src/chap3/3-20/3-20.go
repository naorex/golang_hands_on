package main

import "fmt"

type General interface {} // 空のインターフェースにするとなんでも受け入れる型になる

type GData interface {
	Set(nm string, g General)
	Print()
}

type GDataImpl struct {
	Name string
	Data General
}

func (gd *GDataImpl) Set(nm string, g General) {
	gd.Name = nm
	gd.Data = g
}

func (gd *GDataImpl) Print() {
	fmt.Printf("<<%s>> ", gd.Name)
	fmt.Println(gd.Data)
}

func main() {
	var data = []GDataImpl{}
	data = append(data, GDataImpl{"Taro", 123})
	data = append(data, GDataImpl{"Hanako", "hello!"}) // どんな型でもOK
	data = append(data, GDataImpl{"Sachiko", []int{123, 456, 789}})
	for _, ob := range data {
		ob.Print()
	}
}
