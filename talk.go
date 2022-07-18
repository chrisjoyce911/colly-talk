package main

import (
	"fmt"

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
		e.ForEach("tr td:first-child i", func(_ int, td *colly.HTMLElement) {
			fmt.Println(td.Text)
		})
	})

	c.Visit("https://en.wikipedia.org/wiki/List_of_Australian_and_Antarctic_dinosaurs")

	/*

		Developer tools should be totally used. Many times, elements won’t be present with IDs and class names (e.g. Facebook after their new UI update).
		All elements have randomized classes and IDs. Use XPath for these scenarios.

		Lets get a list of Australian and Antarctic dinosaurs
	*/

}
