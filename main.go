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

func (s *Shop) NextOpeningDate(t time.Time) time.Time {
	oneDay := time.Hour * 24
	nextDay := t.Add(oneDay)
	for {
		if s.IsOpenOn(nextDay) {
			year, month, day := nextDay.Date()
			return time.Date(year, month, day, s.openingHours[0], 0, 0, 0, nextDay.Location())
		}
		nextDay = nextDay.Add(oneDay)
	}
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
