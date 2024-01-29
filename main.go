package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"
)

func main() {
	year := flag.Int("year", 0, "Enter year")
	month := flag.Int("month", 0, "Enter month")

	flag.Parse()

	if *year == 0 {
		*year = time.Now().Year()
	}

	if *month == 0 {
		*month = int(time.Now().Month())
	}

	totalDays := daysInMonth(year, month)

	date := time.Date(*year, 1, 1, 0, 0, 0, 0, time.UTC)

	startDay := int(date.Weekday())

	rows := int(math.Ceil(float64(totalDays) / 7))
	cols := 7

	counter := 1

	head := [...]string{
		"Po",
		"Wt",
		"Śr",
		"Cz",
		"Pt",
		"So",
		"Ni",
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{}

	for _, row := range head {
		cell := []*simpletable.Cell{{Align: simpletable.AlignCenter, Text: row}}
		// fmt.Printf("%s ", row)
		table.Header.Cells = append(table.Header.Cells, cell...)
	}

	var r [][]*simpletable.Cell

	for i := 0; i < rows; i++ {
		var rowCells []*simpletable.Cell // Initialize a new slice for each row
		for j := 0; j < cols; j++ {
			if j < startDay && i < 1 {
				rowCells = append(rowCells, &simpletable.Cell{Align: simpletable.AlignCenter, Text: ""})
			} else if counter > totalDays {
				rowCells = append(rowCells, &simpletable.Cell{Align: simpletable.AlignCenter, Text: ""})
			} else {
				rowCells = append(rowCells, &simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(counter)})
				counter++
			}
		}
		r = append(r, rowCells) // Append the row to the outer slice

		table.Body.Cells = append(table.Body.Cells, rowCells...)
	}

	fmt.Printf("Showing calendar for date: %v / %v\n", date.Month(), date.Year())

	fmt.Println("")

	fmt.Println(table.String())
}

func daysInMonth(y *int, m *int) int {

	if *m < 1 || *m > 12 {
		*m = int(time.Now().Month())
	}
	// add year validation
	if *y <= 1970 && len(strconv.Itoa(*y)) != 4 {
		*y = time.Now().Year()
	}
	if *m == 2 {
		if *y%400 == 0 || (*y%4 == 0 && *y%100 != 0) {
			return 29
		}
	}
	daysInMonth := [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	return daysInMonth[*m-1]
}
