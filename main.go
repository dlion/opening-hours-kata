package main

import "time"

type Shop struct {
	openingDays  []time.Weekday
	openingHours []int
}

func (s *Shop) IsOpenOn(t time.Time) bool {
	if isAnOpeningDay(s.openingDays, t.Weekday()) && isInOpeningHours(s.openingHours, t.Hour()) {
		return true
	}
	return false
}

func isAnOpeningDay(o []time.Weekday, t time.Weekday) bool {
	for _, d := range o {
		if d == t {
			return true
		}
	}
	return false
}

func isInOpeningHours(o []int, t int) bool {
	openingHour := o[0]
	closingHour := o[1]
	if t >= openingHour && t <= closingHour {
		return true
	}
	return false
}
