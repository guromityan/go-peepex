package main

import (
	"log"

	"github.com/alecthomas/kingpin"
	"github.com/guromityan/go-peepex/app"
)

var (
	cli   = kingpin.New("peepex", "A handy tool to peek at the contents of an Excel workbook.")
	files = cli.Arg("files", "Excel files (ext: .xlsx)").ExistingFilesOrDirs()
)

func main() {
	filename := "samples/test1.xlsx"
	peepDatas, err := app.Peep(filename, "Sheet1,Sheet2,Sheet3", "A1,B1")
	if err != nil {
		log.Fatalln(err)
	}

	app.TableView(peepDatas)
}
