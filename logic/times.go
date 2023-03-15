package logic

import (
	"fmt"
	"time"
)

func TimeParse(date, layout string) (t time.Time) {
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

func TimeDiff(actual, previous time.Time) float64 {
	diff := actual.Sub(previous)
	infloat := diff.Seconds()

	return infloat
}
