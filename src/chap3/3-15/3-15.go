package main

import "fmt"

type Data interface {
	Initial(name string, data []int)
	PrintData()
}

type Mydata struct {
	Name string
	Data []int
}

func (md *Mydata) Initial(name string, data []int) {
	md.Name = name
	md.Data = data
}

func (md *Mydata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

func main() {
	var ob Data = new(Mydata) // Mydata を Data型として扱う
	ob.Initial("Sachiko", []int{55, 66, 77})
	ob.PrintData()
}
