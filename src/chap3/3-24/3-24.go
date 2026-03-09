package main

import (
	"fmt"
	"reflect"
)

type General interface {}

type GData interface {
	Set(nm string, g General) GData
	Print()
}

type NData struct {
	Name string
	Data []int
}

func (nd *NData) Set(nm string, g General) GData {
	nd.Name = nm
	if reflect.TypeOf(g) == reflect.SliceOf(reflect.TypeOf(0)) { // 型を判定
		nd.Data = g.([]int)
	}
	return nd
}

func (nd *NData) Print() {
	fmt.Printf("<<%s>> value: %d\n", nd.Name, nd.Data)
}

type SData struct {
	Name string
	Data []string
}

func (sd *SData) Set(nm string, g General) GData {
	sd.Name = nm
	if reflect.TypeOf(g) == reflect.SliceOf(reflect.TypeOf("")) { // 型を判定
		sd.Data = g.([]string)
	}
	return sd
}

func (sd *SData) Print() {
	fmt.Printf("* %s [%s] *\n", sd.Name, sd.Data)
}

func main () {
	var data = []GData{}
	data = append(data, new(NData).Set("Taro", []int{123, 456, 789}))
	data = append(data, new(SData).Set("Jiro", []string{"Hello!", "how low?"}))
	data = append(data, new(NData).Set("Hanako", "98700")) // 型が間違っていても動作する => <<Hanako>> value: []
	data = append(data, new(SData).Set("Taro", "happy?")) // 同上 => * Taro [[]] *
	for _, ob := range data {
		ob.Print()
	}
}
