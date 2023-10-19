package main

import (
	"fmt"
	"time"
	"errors"
)

type DateRange struct {
	Start time.Time
	End   time.Time
}

func NewDateRange(start, end time.Time) (DateRange, error)  {
	if start.IsZero() || end.IsZero() {
		return DateRange{}, errors.New("Invalid operation")
	}

	if end.Before(start) {
		return DateRange{}, errors.New("Invalid operation")
	}

	return DateRange{Start:start, End:end}, nil

}

func (d DateRange) Hours() float64 {
	return d.End.Sub(d.Start).Hours()
}

func main() {
	lifetime, _ := NewDateRange(time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC), time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC))

	fmt.Println(lifetime.Hours())

	travelInTime, _ := NewDateRange(time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC), time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC))

	fmt.Println(travelInTime.Hours())
}
