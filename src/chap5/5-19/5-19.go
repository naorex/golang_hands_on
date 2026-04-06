package main

import (
	"database/sql"
	"fmt"
	"hello"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Mydata struct {
	ID int
	Name string
	Mail string
	Age int
}

func (m *Mydata) Str() string {
	return "<\"" + strconv.Itoa(m.ID) + ":" + m.Name + "\"" + m.Mail + "," + strconv.Itoa(m.Age) + ">"
}
var qry string = "SELECT * FROM mydata WHERE id = ?" // プレースホルダ

func main() {
	con, er := sql.Open("sqlite3", "../data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	for true {
		s := hello.Input("id")
		if s == "" {
			break
		}
		n, er := strconv.Atoi(s)
		if er != nil {
			panic(er)
		}
		rs := con.QueryRow(qry, n) // 1行処理だけなら、QueryRowをつかった方が処理がすっきりする
		var md Mydata
		er2 := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er2 != nil {
			panic(er2)
		}
		fmt.Println(md.Str())
		}
	fmt.Println("***end***")
}
