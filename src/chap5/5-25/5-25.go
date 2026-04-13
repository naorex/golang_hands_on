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
var qry string = "SELECT * FROM mydata WHERE name like ? OR mail LIKE ?"

func main() {
	con, er := sql.Open("sqlite3", "../data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	ids := hello.Input("update ID")
	id, _ := strconv.Atoi(ids)
	qry := "SELECT * FROM mydata WHERE id = ?"
	rw := con.QueryRow(qry, id)
	tgt := mydatafmRw(rw)
	ae := strconv.Itoa(tgt.Age)
	nm := hello.Input("name(" + tgt.Name + ")")
	ml := hello.Input("mail(" + tgt.Mail + ")")
	ge := hello.Input("age(" + ae +")")
	ag, _ := strconv.Atoi(ge)

	if nm == "" {
		nm = tgt.Name
	}
	if ml == "" {
		ml = tgt.Mail
	}
	if ge == "" {
		ag = tgt.Age
	}
	qry = "UPDATE mydata SET name=?, mail=?, age=? WHERE id=?"
	con.Exec(qry, nm, ml, ag, id)
	showRecord(con)
}

func showRecord(con *sql.DB) {
	qry = "SELECT * FROM mydata"
	rs, _ := con.Query(qry)
	for rs.Next() {
		fmt.Println(mydatafmRws(rs).Str())
	}
}

func mydatafmRws(rs *sql.Rows) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if er != nil {
		panic(er)
	}
	return &md
}

func mydatafmRw(rs *sql.Row) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if er != nil {
		panic(er)
	}
	return &md
}
