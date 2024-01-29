package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Init")
	year := flag.Int("year", 0, "Enter year")
	month := flag.Int("month", 0, "Enter month")

	flag.Parse()

	date := time.Date(*year, time.Month(*month), 1, 0, 0, 0, 0, time.Local)
	totalDays := daysInMonth(year, month)
	startDay := int(date.Weekday())

	rows := int(math.Ceil(float64(totalDays) / 7))
	cols := 7

	table := make([][]int, rows)
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

	for i := 0; i < rows; i++ {
		table[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			if j < startDay-1 && i == 0 {
				table[i][j] = 0
			} else if counter > totalDays {
				table[i][j] = 0
			} else {
				table[i][j] = counter
				counter++
			}
		}
	}

	fmt.Println(int(startDay))

	fmt.Printf("%v-%v\n", *year, *month)

	// table := simpletable.New()

	// table.Header = &simpletable.Header{}

	for _, row := range head {
		// cell := []*simpletable.Cell{{Align: simpletable.AlignCenter, Text: row}}
		fmt.Printf("%s ", row)
		// table.Header.Cells = append(table.Header.Cells, cell...)
	}
	fmt.Println("")
	// Print the table
	for i := 0; i < rows; i++ {
		fmt.Println(table[i])
	}

	// day := 1
	// fmt.Println(totalDays / 7)
	// for c < totalDays {
	// 	for day <= 7 {

	// 		//r := []*simpletable.Cell{{Align: simpletable.AlignCenter, Text: fmt.Sprint(day)}}
	// 		if day >= startDay {
	// 			fmt.Printf("%v ", day)
	// 		}
	// 		//table.Body.Cells = append(table.Body.Cells, r)
	// 		day++
	// 	}
	// }

	//fmt.Println(table.String())
	//fmt.Println(table.String())
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
