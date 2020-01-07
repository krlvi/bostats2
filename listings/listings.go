package listings

import (
	"time"
)

type Entry struct {
	Address  string
	Sqm      float64
	Rooms    float64
	Fee      int
	Link     string
	Price    int
	PriceSqm int
	Date     time.Time
}

func Filter(in []Entry, f func(Entry) bool) []Entry {
	var out []Entry
	for _, l := range in {
		if f(l) {
			out = append(out, l)
		}
	}
	return out
}

func Map(in []Entry, f func(Entry) map[time.Time]int) []map[time.Time]int {
	var out []map[time.Time]int
	for _, l := range in {
		out = append(out, f(l))
	}
	return out
}
