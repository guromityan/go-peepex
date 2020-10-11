package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/guromityan/go-peepex/app"
)

const version = "0.9.0"

var (
	cli    = kingpin.New("peepex", "A handy tool to peek at the contents of an Excel workbook.")
	files  = cli.Arg("files", "Excel files").ExistingFiles()
	sheets = cli.Flag("sheet", "Comma-separated sheet names.  (ex)Sheet1,Sheet2").Short('s').Required().String()
	cells  = cli.Flag("cell", "Comma-separated cell coordinate.  (ex)A1,B2").Short('c').Required().String()
)

func main() {
	cli.Version(version)
	cli.Parse(os.Args[1:])

	if len(*files) == 0 {
		fmt.Println("Please specify any files.")
		return
	}

	if *sheets == "" {
		fmt.Println("Please specify any sheets.")
		return
	}

	if *cells == "" {
		fmt.Println("Please specify any cells.")
		return
	}

	peepDataWithFiles := [][]*app.PeepData{}
	for _, filename := range *files {
		peepDatas, err := app.Peep(filename, *sheets, *cells)
		if err != nil {
			log.Fatalln(err)
		}
		peepDataWithFiles = append(peepDataWithFiles, peepDatas)
	}
	app.TableView(peepDataWithFiles, len(*files))
}
