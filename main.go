package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"
)

const (
	ColorDefault = "\x1b[39m"
	ColorBlue    = "\x1b[94m"
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

	date := time.Date(*year, time.Month(*month), 1, 0, 0, 0, 0, time.UTC)
	today := time.Now()

	startDay := int(date.Weekday())
	totalDays := daysInMonth(year, month)

	rows := int(math.Ceil(float64(totalDays+startDay) / 7))

	cols := 7

	counter := 1

	head := [...]string{
		"Mon",
		"Tue",
		"Wed",
		"Thu",
		"Fri",
		"Sat",
		"Sun",
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
			if j < startDay-1 && i < 1 {
				rowCells = append(rowCells, &simpletable.Cell{Align: simpletable.AlignCenter, Text: ""})
			} else if counter > totalDays {
				rowCells = append(rowCells, &simpletable.Cell{Align: simpletable.AlignCenter, Text: ""})
			} else {
				var text string
				if today.Day() == counter && *year == today.Year() && *month == int(today.Month()) {
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

	fmt.Printf("Current date is: %v / %s / %v", today.Day(), today.Month(), today.Year())

	fmt.Println("")

	fmt.Printf("Showing calendar for date: %v / %v\n", date.Month(), date.Year())

	fmt.Println("")

	fmt.Println(table.String())

	fmt.Println("")

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

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}
