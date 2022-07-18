package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	var dinosaurs = map[string]string{}

	c := colly.NewCollector(
		colly.AllowedDomains("wikipedia.org", "en.wikipedia.org"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnHTML("table.wikitable tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr td:first-child i a", func(_ int, el *colly.HTMLElement) {

			fmt.Printf("%s - %s\n", strings.TrimSpace(el.Text), el.Attr("href"))
			dinosaurs[strings.TrimSpace(el.Text)] = el.Attr("href")
		})
	})

	c.Visit("https://en.wikipedia.org/wiki/List_of_Australian_and_Antarctic_dinosaurs")

	/*
		Get a list of Australian and Antarctic dinosaurs
		Get link for each dinosaur
	*/

}
