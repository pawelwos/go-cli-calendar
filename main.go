package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"
)

func main() {
	fmt.Println("Init")
	year := flag.Int("year", 0, "Enter year")
	month := flag.Int("month", 0, "Enter month")

	flag.Parse()

	// if no year and month provided get current date
	days := daysInMonth(year, month)

	fmt.Println(days)
	fmt.Printf("%v-%v\n", *year, *month)

	head := [...]string{
		"Poniedziałek",
		"Wtorek",
		"Środa",
		"Czwartek",
		"Piątek",
		"Sobota",
		"Niedziela",
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{}

	for _, row := range head {
		cell := []*simpletable.Cell{{Align: simpletable.AlignCenter, Text: row}}
		table.Header.Cells = append(table.Header.Cells, cell...)
	}

	fmt.Println(table.String())
}

func daysInMonth(m *int, y *int) int {
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
