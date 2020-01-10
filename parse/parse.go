package parse

import (
	"bostats2/listings"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

func FindListings(doc *goquery.Document) ([]listings.Entry, error) {
	var out []listings.Entry
	doc.Find(".sold-results__normal-hit").Each(func(_ int, s *goquery.Selection) {
		l, err := parseListing(s)
		if err != nil {
			log.Println(err)
			return
		}
		out = append(out, l)
	})
	return out, nil
}

func parseListing(s *goquery.Selection) (listings.Entry, error) {
	ap := listings.Entry{}
	addr, err := contentAt(s, ".sold-property-listing__location .item-result-meta-attribute-is-bold")
	if err != nil {
		return listings.Entry{}, err
	}
	ap.Address = strings.TrimSpace(addr)
	sqm, rooms, err := sqmRoom(s)
	if err != nil {
		return listings.Entry{}, err
	}
	ap.Sqm = sqm
	ap.Rooms = rooms
	fee, err := fee(s)
	if err != nil {
		return listings.Entry{}, err
	}
	ap.Fee = fee
	link, err := attrAt(s, ".item-link-container", "href")
	if err != nil {
		return listings.Entry{}, err
	}
	ap.Link = link
	price, err := price(s)
	if err != nil {
		return listings.Entry{}, err
	}
	ap.Price = price
	pSqm, err := priceSqm(s)
	if err != nil {
		return listings.Entry{}, err
	}
	ap.PriceSqm = pSqm
	date, err := date(s)
	if err != nil {
		return listings.Entry{}, err
	}
	ap.Date = date
	return ap, nil
}

func PagesAvailable(doc *goquery.Document) (int, error) {
	s := doc.Find(".padded-container .result-tools .centered").First()
	showing := s.Nodes[0].FirstChild.NextSibling.FirstChild.Data
	from := s.Nodes[0].LastChild.Data
	return totalPages(showing, from)
}

func totalPages(showing, last string) (int, error) {
	to, err := strconv.Atoi(strings.TrimSpace(strings.Split(showing, "-")[1]))
	if err != nil {
		return 0, err
	}
	total, err := strconv.Atoi(strings.TrimSpace(strings.Split(last, "av")[1]))
	if err != nil {
		return 0, err
	}
	return int(math.Ceil(float64(total) / float64(to))), nil
}

func sqmRoom(s *goquery.Selection) (sqm, rooms float64, err error) {
	raw, err := contentAt(s, ".sold-property-listing__size .sold-property-listing__subheading")
	if err != nil {
		return 0, 0, err
	}
	left := strings.TrimSpace(strings.Split(raw, "m")[0])
	sqm, err = strconv.ParseFloat(strings.ReplaceAll(left, ",", "."), 64)
	if err != nil {
		return 0, 0, err
	}
	right := strings.ReplaceAll(strings.Split(strings.TrimSpace(strings.Split(raw, "m")[1][2:]), "r")[0], ",", ".")
	rooms, err = strconv.ParseFloat(strings.TrimSpace(right), 64)
	if err != nil {
		return 0, 0, err
	}
	return
}

func price(s *goquery.Selection) (int, error) {
	raw, err := contentAt(s, ".sold-property-listing__price .sold-property-listing__subheading")
	if err != nil {
		return 0, err
	}
	right := strings.Fields(strings.TrimSpace(strings.Split(strings.Split(raw, "Slutpris")[1], "kr")[0]))
	px, err := strconv.Atoi(strings.Join(right, ""))
	if err != nil {
		log.Fatal(err)
	}
	return px, nil
}

func priceSqm(s *goquery.Selection) (int, error) {
	raw, err := contentAt(s, ".sold-property-listing__price .sold-property-listing__price-per-m2")
	if err != nil {
		return 0, err
	}
	fields := strings.Fields(strings.TrimSpace(strings.Split(raw, "kr")[0]))
	i, _ := strconv.Atoi(fields[0] + fields[1])
	return i, nil
}

func fee(s *goquery.Selection) (int, error) {
	raw, err := contentAt(s, ".sold-property-listing__size .sold-property-listing__fee")
	if err != nil {
		return 0, err
	}
	fields := strings.Fields(strings.TrimSpace(strings.Split(raw, "kr")[0]))
	feeString := fields[0]
	if len(fields) == 2 {
		feeString = fields[0] + fields[1]
	}
	i, _ := strconv.Atoi(feeString)
	return i, nil
}

func date(s *goquery.Selection) (time.Time, error) {
	raw, err := contentAt(s, ".sold-property-listing__price .sold-property-listing__sold-date")
	if err != nil {
		return time.Time{}, err
	}
	fields := strings.Fields(strings.TrimSpace(strings.Split(raw, "Såld")[1]))
	months := map[string]string{
		"januari":   "01",
		"februari":  "02",
		"mars":      "03",
		"april":     "04",
		"maj":       "05",
		"juni":      "06",
		"juli":      "07",
		"augusti":   "08",
		"september": "09",
		"oktober":   "10",
		"november":  "11",
		"december":  "12",
	}
	day := fields[0]
	if len(day) == 1 {
		day = "0" + day
	}
	t, err := time.Parse("2006-01-02", fields[2]+"-"+months[fields[1]]+"-"+day)
	if err != nil {
		log.Fatal(err)
	}
	return t, nil
}

func contentAt(selection *goquery.Selection, class string) (string, error) {
	c := selection.Find(class).Map(func(_ int, selection *goquery.Selection) string {
		if len(selection.Nodes) == 0 {
			return ""
		}
		return selection.Nodes[0].FirstChild.Data
	})
	if len(c) == 0 {
		return "", errors.New("could not parse content for " + class)
	}
	return c[0], nil
}

func attrAt(selection *goquery.Selection, class, attr string) (string, error) {
	c := selection.Find(class).Map(func(_ int, selection *goquery.Selection) string {
		href, _ := selection.Attr(attr)
		return href
	})
	if len(c) == 0 {
		return "", errors.New("could not parse attr " + attr + " for " + class)
	}
	return c[0], nil
}
