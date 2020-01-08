package scrape

import (
	"bostats2/listings"
	"bostats2/parse"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func AllPages(url string) ([]listings.Entry, error) {
	rsp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(rsp.Body)
	if err != nil {
		return nil, err
	}

	pages, err := parse.PagesAvailable(doc)
	if err != nil {
		// Hemnet caps out at 50 pages
		//return nil, err
		pages = 50
	}

	var wg sync.WaitGroup
	var out []listings.Entry
	c := make(chan []listings.Entry)
	done := make(chan bool)
	go func() {
		for n := range c {
			out = append(out, n...)
		}
		done <- true
	}()
	for i := 1; i <= pages; i++ {
		wg.Add(1)
		go parsePage(url, i, c, &wg)
	}
	wg.Wait()
	close(c)
	<- done
	return out, nil
}

func parsePage(url string, page int, c chan []listings.Entry, wg *sync.WaitGroup) {
	defer wg.Done()
	url = url + "&page=" + strconv.Itoa(page)
	rsp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(rsp.Body)
	if err != nil {
		return
	}
	aps, err := parse.FindListings(doc)
	if err != nil {
		return
	}
	c <- aps
	return
}
