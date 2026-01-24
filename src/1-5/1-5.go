package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	name := input("type your name") // := 変数の定義と値の代入を同時に
	fmt.Println("Hello, " + name + "!!")
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
