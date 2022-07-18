package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

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

			foundURL := el.Request.AbsoluteURL(el.Attr("href"))
			fmt.Printf("%s\n\t%s\n\n", strings.TrimSpace(el.Text), foundURL)

		})
	})

	c.Visit("https://en.wikipedia.org/wiki/List_of_Australian_and_Antarctic_dinosaurs")

}
