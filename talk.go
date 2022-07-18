package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	/*
		Set Up Colly With a Target Website
		Let’s create a function that initializes the collector and then calls the target website.
		Later, we can extend the function and break it down into subparts based on the requirements.
	*/

	c := colly.NewCollector(
		colly.AllowedDomains("wikipedia.org", "en.wikipedia.org"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit("https://en.wikipedia.org/wiki/List_of_Australian_and_Antarctic_dinosaurs")

	/*
		The snippet above initializes a collector and restricts it to the Wikipedia domain.
		We have also attached an
		* OnRequest to the collector to know when they start running.
		* OnHTML that will print the page title
		Finally, we call c.Visit with a URL that opens the article 'List_of_Australian_and_Antarctic_dinosaurs'
	*/
}
