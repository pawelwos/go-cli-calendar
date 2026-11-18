package main

import (
	"flag"
	"fmt"

	"github.com/alexeyco/simpletable"
	"example.com/calendar-src/calendar"
)

const (
	ColorDefault = "\x1b[39m"
	ColorBlue    = "\x1b[94m"
)

func main() {
	year := flag.Int("year", 0, "Enter year")
	month := flag.Int("month", 0, "Enter month")

	flag.Parse()

	var cal = calendar.Create(*year, *month)

	table := simpletable.New()

	table.Header = &simpletable.Header{}

	for _, row := range calendar.GetHead() {
		cell := []*simpletable.Cell{{Align: simpletable.AlignCenter, Text: row}}
		table.Header.Cells = append(table.Header.Cells, cell...)
	}

	var r [][]*simpletable.Cell

	for _, week := range cal.GetBody() {
		var rowCells []*simpletable.Cell
		for _, day := range week {
			var text string
			if day != 0 {
				if cal.Today.Day() == day && cal.Year == cal.Today.Year() && cal.Month == int(cal.Today.Month()) {
					text = blue(fmt.Sprint(day))
				} else {
					text = fmt.Sprint(day)
				}
			}
			rowCells = append(rowCells, &simpletable.Cell{Align: simpletable.AlignCenter, Text: text})
		}
		r = append(r, rowCells)
	}

	table.Body.Cells = append(table.Body.Cells, r...)

	fmt.Println("")

	fmt.Printf("Current date is: %v / %s / %v", cal.Today.Day(), cal.Today.Month(), cal.Today.Year())

	fmt.Println("")

	fmt.Printf("Showing calendar for date: %v / %v\n", cal.Month, cal.Year)

	fmt.Println("")

	fmt.Println(table.String())

	fmt.Println("")

}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}
