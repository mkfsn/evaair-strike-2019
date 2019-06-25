package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type panel struct {
	control string
	text    string
}

type Flight struct {
	Origin       string
	Destination  string
	FlightNumber string
	Status       string
	Schedule     string
}

type DateStatus struct {
	Arrival   []*Flight
	Departure []*Flight
}

func main() {
	var start string
	var end string
	var airport string
	flag.StringVar(&start, "start", "20190628", "start date to query")
	flag.StringVar(&end, "end", "20190705", "end date to query")
	flag.StringVar(&airport, "airport", "TPE", "airport to query")
	flag.Parse()

	doc, err := goquery.NewDocument(fmt.Sprintf("https://booking.evaair.com/flyeva/EVA/B2C/flight-status-erc.aspx?lang=zh-tw&cmstitle=erc-note9&date=%s-%s&airport=%sOrderby=&reqtime=&ACTCODE=&REASON=", start, end, airport))
	if err != nil {
		log.Fatal(err)
	}

	panels := []*panel{}
	doc.Find("#MainContent_pnl_tablist > a").Each(func(i int, selection *goquery.Selection) {
		control, ok := selection.Attr("aria-controls")
		if !ok {
			return
		}
		panels = append(panels, &panel{control: control, text: selection.Text()})
	})

	timeTable := make(map[string]*DateStatus)
	for _, panel := range panels {
		items := doc.Find(fmt.Sprintf("#%s div.accordion > div.item", panel.control))

		status := &DateStatus{
			Departure: parseTable(items.Eq(0).Find("table")),
			Arrival:   parseTable(items.Eq(1).Find("table")),
		}

		var year, month, day int
		fmt.Sscanf(panel.text, "%d年%d月%d日", &year, &month, &day)
		timeTable[fmt.Sprintf("%04d%02d%02d", year, month, day)] = status
	}

	b, err := json.MarshalIndent(timeTable, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(b))
}

func parseTable(table *goquery.Selection) []*Flight {
	flights := make([]*Flight, 0)
	table.Find("tr").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			return
		}
		flight := &Flight{
			FlightNumber: strings.TrimSpace(selection.Find("td").Eq(0).Text()),
			Origin:       strings.TrimSpace(selection.Find("td").Eq(1).Text()),
			Destination:  strings.TrimSpace(selection.Find("td").Eq(2).Text()),
			Schedule:     strings.TrimSpace(selection.Find("td").Eq(3).Text()),
			Status:       strings.TrimSpace(selection.Find("td").Eq(5).Text()),
		}
		flights = append(flights, flight)
	})
	return flights
}
