package api

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type LottoData struct {
	Date             time.Time `json:"date"`
	FirstPrize       string    `json:"first_prize"`
	TwentySevenPrize string    `json:"twenty_seven_prize"`
}

func StartCrawl() {
	c := colly.NewCollector()

	crawlUrl := "https://xoso.me/"

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div.one-city table.kqmb", func(e *colly.HTMLElement) {
		selector := e.DOM

		data := LottoData{}

		data.Date = time.Now().UTC()
		data.FirstPrize = selector.Find("span.v-gdb").Text()
		data.TwentySevenPrize = ""

		fmt.Println(data.FirstPrize)

	})

	c.Visit(crawlUrl)
}
