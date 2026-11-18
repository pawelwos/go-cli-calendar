package calendar

import (
	"math"
	"time"
)

type Calendar struct {
	Year      int
	Month     int
	Today     time.Time
	StartDay  int
	TotalDays int
	Rows      int
	Cols      int
}

func Create(year int, month int) Calendar {
	var c = Calendar{}

	c.Today = time.Now()
	c.Year = year
	c.Month = month

	if c.Year < 1 || c.Year > 9999 {
		c.Year = c.Today.Year()
	}

	if c.Month < 1 || c.Month > 12 {
		c.Month = int(c.Today.Month())
	}

	date := time.Date(c.Year, time.Month(c.Month), 1, 0, 0, 0, 0, time.UTC)

	leadingBlanks := (int(date.Weekday()) + 6) % 7
	c.StartDay = leadingBlanks + 1
	c.TotalDays = daysInMonth(c.Year, c.Month)

	c.Rows = int(math.Ceil(float64(c.TotalDays+leadingBlanks) / 7))
	c.Cols = 7

	return c
}

func daysInMonth(y int, m int) int {
	// Leap year check
	if m == 2 {
		if y%400 == 0 || (y%4 == 0 && y%100 != 0) {
			return 29
		}
	}
	days := [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	return days[m-1]
}

func GetHead() [7]string {
	head := [7]string{
		"Mon",
		"Tue",
		"Wed",
		"Thu",
		"Fri",
		"Sat",
		"Sun",
	}

	return head
}

func (cal Calendar) GetBody() [][]int {
	counter := 1
	table := make([][]int, cal.Rows)

	for i := 0; i < cal.Rows; i++ {
		table[i] = make([]int, cal.Cols)

		for j := 0; j < cal.Cols; j++ {
			if i == 0 && j < cal.StartDay-1 || counter > cal.TotalDays {
				table[i][j] = 0
			} else {
				table[i][j] = counter
				counter++
			}
		}
	}

	return table
}
