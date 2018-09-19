package main

import (
	"testing"
	"time"
)

var (
	wednesday        = time.Date(2016, time.May, 11, 12, 22, 11, 824, time.Local)
	wednesdayEvening = time.Date(2016, time.May, 11, 23, 22, 11, 824, time.Local)
	thursday         = time.Date(2016, time.May, 12, 12, 22, 11, 824, time.Local)
	fridayMorning    = time.Date(2016, time.May, 13, 8, 0, 0, 0, time.Local)

	shop = &Shop{
		openingDays:  []time.Weekday{time.Monday, time.Wednesday, time.Friday},
		openingHours: []int{8, 16},
	}
)

func TestIsOpenOn(t *testing.T) {
	t.Run("Should be open on an opening day", func(t *testing.T) {
		expected := true
		got := shop.IsOpenOn(wednesday)

		if got != expected {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	})

	t.Run("Should be closed on an closing day", func(t *testing.T) {
		expected := false
		got := shop.IsOpenOn(thursday)

		if got != expected {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	})

	t.Run("Should be closed in a closing hour", func(t *testing.T) {
		expected := false
		got := shop.IsOpenOn(wednesdayEvening)

		if got != expected {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	})
}

func TestNextOpeningDate(t *testing.T) {
	t.Run("Should have a next opening day", func(t *testing.T) {
		expected := fridayMorning
		got := shop.NextOpeningDate(wednesday)

		if got != expected {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	})
}
