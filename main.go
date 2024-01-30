package main

import (
	"flag"
	"fmt"

	"github.com/alexeyco/simpletable"
	"github.com/pawelwos/go-calendar"
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

	counter := 1

	for i := 0; i < cal.Rows; i++ {
		var rowCells []*simpletable.Cell // Initialize a new slice for each row
		for j := 0; j < cal.Cols; j++ {
			if j < cal.StartDay-1 && i < 1 {
				rowCells = append(rowCells, &simpletable.Cell{Align: simpletable.AlignCenter, Text: ""})
			} else if counter > cal.TotalDays {
				rowCells = append(rowCells, &simpletable.Cell{Align: simpletable.AlignCenter, Text: ""})
			} else {
				var text string
				if cal.Today.Day() == counter && cal.Year == cal.Today.Year() && cal.Month == int(cal.Today.Month()) {
					text = blue(fmt.Sprint(counter))
				} else {
					text = fmt.Sprint(counter)
				}
				rowCells = append(rowCells, &simpletable.Cell{Align: simpletable.AlignCenter, Text: text})
				counter++
			}
		}
		r = append(r, rowCells) // Append the row to the outer slice
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
