package main

import (
	"bostats2/listings"
	"bostats2/scrape"
	"fmt"
	h "github.com/aybabtme/uniplot/histogram"
	"github.com/montanaflynn/stats"
	"os"
	"sort"
	"time"
)

func main() {
	hemnetUrl := os.Args[1]
	ls, _ := scrape.AllPages(hemnetUrl)
	overTime := listings.Map(ls, func(listing listings.Entry) map[time.Time]int {
		return map[time.Time]int{listing.Date: listing.PriceSqm}
	})
	buckets := monthBuckets(overTime)

	keys := make([]time.Time, 0, len(buckets))
	for k := range buckets {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Before(keys[j])
	})
	fmt.Println("\nPrice/Sqm per month percentiles:")
	fmt.Printf("|%4s|%10s|%3s|%8s|%8s|%8s|\n", "year", "month", "vol", "50th", "75th", "90th")
	for _, d := range keys {
		data := intToFloat64(buckets[d])
		median, _ := stats.Median(data)
		p75, _ := stats.Percentile(data, 75)
		p90, _ := stats.Percentile(data, 90)
		fmt.Printf("|%4d|%10s|%3d|%8.1f|%8.1f|%8.1f|\n", d.Year(), d.Month(), len(data), median, p75, p90)
		//fmt.Println(d.Year(), d.Month(), len(data), median, p75, p90)
	}
	fmt.Println("\nPrice/Sqm frequency histogram for the period", keys[0].Month(), keys[0].Year(), "to", keys[len(keys)-1].Month(), keys[len(keys)-1].Year())
	hist := h.Hist(20, listings.PriceSqm(ls))
	h.Fprintf(os.Stdout, hist, h.Linear(20), func(v float64) string {
		return fmt.Sprintf("%.0f", v)
	})
}

func intToFloat64(in []int) []float64 {
	out := []float64{}
	for _, i := range in {
		out = append(out, float64(i))
	}
	return out
}

func monthBuckets(entries []map[time.Time]int) map[time.Time][]int {
	out := map[time.Time][]int{}
	for _, entry := range entries {
		for d, e := range entry {
			justMonth := time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, time.UTC)
			out[justMonth] = append(out[justMonth], e)
		}
	}
	return out
}
