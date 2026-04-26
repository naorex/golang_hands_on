package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	a := app.New()
	w := a.NewWindow("Markdown Gatherer")

	edit := widget.NewMultiLineEntry()
	sc := container.NewScroll(edit)
	fnd := widget.NewEntry()
	inf := widget.NewLabel("Information bar.")

	// --- Helper Functions ---

	showInfo := func(s string) {
		inf.SetText(s)
		dialog.ShowInformation("Info", s, w)
	}

	checkError := func(er error) bool {
		if er != nil {
			inf.SetText(er.Error())
			return true
		}
		return false
	}

	setDB := func() *sql.DB {
		con, er := sql.Open("sqlite3", "./data.sqlite3") // パスをカレントに変更
		if checkError(er) {
			return nil
		}
		return con
	}

	// --- Action Functions ---

	nf := func() {
		dialog.ShowConfirm("Alert", "Clear form?", func(f bool) {
			if f {
				fnd.SetText("")
				w.SetTitle("App")
				edit.SetText("")
				inf.SetText("Clear form.")
			}
		}, w)
	}

	wf := func() {
		fstr := fnd.Text
		if fstr == "" {
			return
		}
		if !strings.HasPrefix(fstr, "http") {
			fstr = "https://" + fstr
			fnd.SetText(fstr)
		}

		// http.Get を使用してドキュメントを取得
		res, er := http.Get(fstr)
		if checkError(er) {
			return
		}
		defer res.Body.Close()

		dc, er := goquery.NewDocumentFromReader(res.Body)
		if checkError(er) {
			return
		}

		ttl := dc.Find("title").Text()
		w.SetTitle(ttl)

		html, er := dc.Html()
		if checkError(er) {
			return
		}

		cvtr := md.NewConverter("", true, nil)
		mkdn, er := cvtr.ConvertString(html)
		if checkError(er) {
			return
		}
		edit.SetText(mkdn)
		inf.SetText("Get web data success.")
	}

	ff := func() {
		qry := "SELECT id, title FROM md_data WHERE title LIKE ?"
		con := setDB()
		if con == nil {
			return
		}
		defer con.Close()

		rs, er := con.Query(qry, "%"+fnd.Text+"%")
		if checkError(er) {
			return
		}
		defer rs.Close()

		resText := ""
		for rs.Next() {
			var id int
			var title string
			er := rs.Scan(&id, &title)
			if checkError(er) {
				return
			}
			resText += fmt.Sprintf("%d: %s\n", id, title)
		}
		edit.SetText(resText)
		inf.SetText("Find: " + fnd.Text)
	}

	idf := func(id int) {
		qry := "SELECT id, title, url, markdown FROM md_data WHERE id = ?"
		con := setDB()
		if con == nil {
			return
		}
		defer con.Close()

		var ID int
		var TT, UR, MR string
		er := con.QueryRow(qry, id).Scan(&ID, &TT, &UR, &MR)
		if checkError(er) {
			return
		}

		w.SetTitle(TT)
		fnd.SetText(UR)
		edit.SetText(MR)
		inf.SetText("Found ID=" + strconv.Itoa(ID))
	}

	sf := func() {
		dialog.ShowConfirm("Alert", "Save data?", func(f bool) {
			if f {
				con := setDB()
				if con == nil {
					return
				}
				defer con.Close()

				qry := "INSERT INTO md_data (title, url, markdown) VALUES (?, ?, ?)"
				_, er := con.Exec(qry, w.Title(), fnd.Text, edit.Text)
				if checkError(er) {
					return
				}
				showInfo("Saved to database!")
			}
		}, w)
	}

	xf := func() {
		dialog.ShowConfirm("Alert", "Export this data?", func(f bool) {
			if f {
				fn := w.Title() + ".md"
				ctt := "# " + w.Title() + "\n\n"
				ctt += "## " + fnd.Text + "\n\n"
				ctt += edit.Text
				er := os.WriteFile(fn, []byte(ctt), 0644)
				if checkError(er) {
					return
				}
				showInfo("Exported to " + fn)
			}
		}, w)
	}

	// --- Theme Toggle ---
	isDark := true
	cf := func() {
		if isDark {
			a.Settings().SetTheme(theme.LightTheme())
			inf.SetText("Changed to Light Theme")
		} else {
			a.Settings().SetTheme(theme.DarkTheme())
			inf.SetText("Changed to Dark Theme")
		}
		isDark = !isDark
	}

	// --- UI Components ---

	cbtn := widget.NewButton("Clear", nf)
	wbtn := widget.NewButton("Get Web", wf)
	fbtn := widget.NewButton("Find", ff)
	ibtn := widget.NewButton("ID Search", func() {
		rid, er := strconv.Atoi(fnd.Text)
		if checkError(er) {
			return
		}
		idf(rid)
	})
	sbtn := widget.NewButton("Save", sf)
	xbtn := widget.NewButton("Export", xf)

	// Menu
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("New", nf),
		fyne.NewMenuItem("Save", sf),
		fyne.NewMenuItem("Export", xf),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Quit", func() { a.Quit() }),
	)
	themeMenu := fyne.NewMenu("View",
		fyne.NewMenuItem("Toggle Theme", cf),
	)
	w.SetMainMenu(fyne.NewMainMenu(fileMenu, themeMenu))

	// Toolbar
	tb := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), nf),
		widget.NewToolbarAction(theme.SearchIcon(), ff),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), sf),
	)

	// Layout
	header := container.NewVBox(
		tb,
		widget.NewForm(widget.NewFormItem("URL/ID/Title", fnd)),
		container.NewHBox(cbtn, wbtn, fbtn, ibtn, sbtn, xbtn),
	)

	content := container.New(layout.NewBorderLayout(header, inf, nil, nil),
		header, inf, sc)

	w.SetContent(content)
	w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()
}
