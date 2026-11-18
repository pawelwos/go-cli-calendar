package calendar

import (
	"reflect"
	"testing"
	"time"
)

func TestCreateExplicitDate(t *testing.T) {
	cal := Create(2025, 2)
	if cal.Year != 2025 || cal.Month != 2 {
		t.Fatalf("got %d-%02d, want 2025-02", cal.Year, cal.Month)
	}
	if cal.TotalDays != 28 {
		t.Fatalf("got TotalDays=%d, want 28", cal.TotalDays)
	}
	// 2025-02-01 is Saturday (Mon=1 .. Sun=7 -> 6)
	if cal.StartDay != 6 {
		t.Fatalf("got StartDay=%d, want 6", cal.StartDay)
	}
	if cal.Rows != 5 {
		t.Fatalf("got Rows=%d, want 5", cal.Rows)
	}
}

func TestCreateDefaultsAndValidation(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name      string
		year      int
		month     int
		wantYear  int
		wantMonth int
	}{
		{"zero defaults", 0, 0, now.Year(), int(now.Month())},
		{"negative year", -5, 3, now.Year(), 3},
		{"year above 9999", 10000, 3, now.Year(), 3},
		{"month below 1", 2024, -1, 2024, int(now.Month())},
		{"month above 12", 2024, 13, 2024, int(now.Month())},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cal := Create(tc.year, tc.month)
			if cal.Year != tc.wantYear || cal.Month != tc.wantMonth {
				t.Fatalf("Create(%d, %d) got %d-%02d, want %d-%02d",
					tc.year, tc.month, cal.Year, cal.Month, tc.wantYear, tc.wantMonth)
			}
			expectedDate := time.Date(cal.Year, time.Month(cal.Month), 1, 0, 0, 0, 0, time.UTC)
			expectedStartDay := (int(expectedDate.Weekday())+6)%7 + 1
			if cal.StartDay != expectedStartDay {
				t.Fatalf("StartDay=%d, want %d (must match validated year/month)", cal.StartDay, expectedStartDay)
			}
		})
	}
}

func TestLeapYears(t *testing.T) {
	tests := []struct {
		year     int
		wantDays int
	}{
		{2024, 29},
		{2000, 29},
		{2025, 28},
		{1900, 28},
	}

	for _, tc := range tests {
		cal := Create(tc.year, 2)
		if cal.TotalDays != tc.wantDays {
			t.Errorf("Feb %d: got %d days, want %d", tc.year, cal.TotalDays, tc.wantDays)
		}
	}
}

func TestSundayStartAlignmentAndRows(t *testing.T) {
	// Nov 2026 starts on Sunday (StartDay=7), 30 days -> 6 rows, 1st at [0][6]
	cal := Create(2026, 11)
	if cal.StartDay != 7 {
		t.Fatalf("Nov 2026: got StartDay=%d, want 7", cal.StartDay)
	}
	if cal.Rows != 6 {
		t.Fatalf("Nov 2026: got Rows=%d, want 6", cal.Rows)
	}

	want := [][]int{
		{0, 0, 0, 0, 0, 0, 1},
		{2, 3, 4, 5, 6, 7, 8},
		{9, 10, 11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20, 21, 22},
		{23, 24, 25, 26, 27, 28, 29},
		{30, 0, 0, 0, 0, 0, 0},
	}
	got := cal.GetBody()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Nov 2026 GetBody() mismatch:\ngot:  %v\nwant: %v", got, want)
	}
}

func TestExactRowCountsNoTrailingEmptyRow(t *testing.T) {
	tests := []struct {
		name     string
		year     int
		month    int
		wantRows int
	}{
		{"Feb 2021 (Mon start, 28 days)", 2021, 2, 4},
		{"Apr 2023 (Sat start, 30 days)", 2023, 4, 5},
		{"May 2021 (Sat start, 31 days)", 2021, 5, 6},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cal := Create(tc.year, tc.month)
			if cal.Rows != tc.wantRows {
				t.Fatalf("got Rows=%d, want %d", cal.Rows, tc.wantRows)
			}
			body := cal.GetBody()
			if len(body) != tc.wantRows {
				t.Fatalf("len(GetBody())=%d, want %d", len(body), tc.wantRows)
			}
			lastRow := body[len(body)-1]
			hasDay := false
			for _, d := range lastRow {
				if d != 0 {
					hasDay = true
					break
				}
			}
			if !hasDay {
				t.Fatalf("last row is completely empty: %v", lastRow)
			}
		})
	}
}
