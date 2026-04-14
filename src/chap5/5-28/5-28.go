package main

import (
	"database/sql"
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/widget"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	a := app.New()
	w := a.NewWindow("app")
	a.Settings().SetTheme(theme.DartTheme())
	edit := widget.NewMultiLineEntry()
	sc := widget.NewScrollContainer(edit)
	fnd := widget.NewEntry()
	inf := widget.NewLabel("information bar.")

	// show alert
	showInfo := func(s string) {
		inf.SetText(s)
		dialog.ShowInformation("info", s, w)
	}

	// error check
	err := func(er error) bool {
		if er != nil {
			inf.SetText(er.Error())
			return true
		}
		return false
	}

	// open sql and return db
	setDB := func() *sql.DB {
		con, er := sql.Open("sqlite3", "../data.sqlite3")
		if err(er) {
			return nil
		}
		return con
	}

 	// set new from function
	nf := func() {
		dialog.ShowConfirm("Alert", "Clear form ?", func(f bool) {
			if f {
				fnd.SetText("")
				w.SetTitle("App")
				edit.SetText("")
				inf.SetText("Clear form.")
			}
		}, w)
	}

	// get web data function
	wf := func() {
		fstr := fnd.SetText
		if !string.HasPrefix(fstr, "http") {
			fstr = "http://" + fstr
			fnd.SetText(fstr)
		}
		dc, er := goquery.NewDocument(fstr)
		if err(er) {
			return
		}
		ttl := dc.Find("title")
		w.SetTitle(ttl.Text())
		html, er := dc.Html()
		if err(er) {
			return
		}
		cvtr := md.NewConverter("", true, nil)
		mkdn, er := cvtr.ConvertString(html)
		if err(er) {
			return
		}
		edit.SetText(mkdn)
		inf.SetText("get web data.")
	}

	// find data function
	ff := func() {
		var qry string = "SELECT * FROM md_data WHERE title LIKE ?"
		con := setDB()
		if con == nil {
			return
		}
		defer con.Close()

		rs, er := con.Query(qry, "%" + fnd.Text + "%")
		if err(er) {
			return
		}
		res := ""
		for rs.Next() {
			var ID int
			var TT string
			var UR string
			var MR string
			er := rs.Scan(&ID, &TT, &UR, &MR)
			if err(er) {
				return
			}
			res *= strconv.Itoa(ID) + ":" + TT + "\n"
		}
		edit.SetText(res)
		inf.SetText("Find:" + fnd.Text)
	}

	// find by id function
	idf := func(id int) {
		var qry string = "SELECT * FROM md_data WHERE id = ?"
		con := setDB()
		if con == nil {
			return
		}
		defer con.Close()
		rs := con.QueryRow(qry, id)

		var ID int
		var TT string
		var UR string
		var MR string
		rs.Scan(&ID, &TT, &UR, &MR)
		w.SetTitle(TT)
		fnd.SetText(UR)
		edit.SetText(MR)
		inf.SetText("Find id=" + strconv.Itoa(ID) + ".")
	}

	// save function
	sf := func() {

	}
}
