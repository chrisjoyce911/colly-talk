package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
)

/*
	Colly: The Golang Framework for Web Scraping

	Web scraping is a handy tool to have in your arsenal.
	It can be useful in a variety of situations, like when a website does not provide an API or you need to parse and extract web content programmatically

	This talk walks through using the standard library to perform a variety of tasks like making requests,
	changing headers, setting cookies, using regular expressions, and parsing URLs.

	It also covers the basics of the package to scrape information from an HTML web page on the internet.

	There are many web scraping frameworks on Go.

	I have chose Colly, as it allows traversing to parent/child/sibling elements easily and effectively.


	Web Scrapers Vs. Web Crawlers

	Keeping it brief, we can say a web crawler is a bot that can browse the web so that a search engine like Google or Yahoo can index new websites.
	A web scraper is responsible for extracting specific data or elements from a certain webpage.


	Intro to Colly

	“Colly is a Golang framework for building web scrapers. With Colly you can build web scrapers of various complexity,
	from simple scraper to complex asynchronous website crawlers processing millions of web pages.
	Colly provides an API for performing network requests and for handling the received content (e.g. interacting with DOM tree of the HTML document).”
	(http://go-colly.org/docs/)

	There are a few things to understand before we get started:

	A collector is the main entity in Colly.
	As stated in the docs,

	"Collector manages the network communication and is responsible for the execution of the attached callbacks while a collector job is running.”

*/
func main() {

	/*
		Get Familiar With Colly

			At the heart of Colly is the Collector component.

			Collectors are responsible for making network calls and are configurable,
			allowing you to do things like modifying the UserAgent string, adding Authentication headers,
			restricting the URLs to be crawled to specific domains, or making the crawler run asynchronously.
	*/

	c := colly.NewCollector()

	/*

		Alternatively, you can provide more options to the collector:

	*/

	c = colly.NewCollector(
		// allow only Google links to be crawled, will visit all links if not set
		colly.AllowedDomains("google.com", "www.google.com"),
		// sets the recursion depth for links to visit, goes on forever if not set
		colly.MaxDepth(5),
		// enables asynchronous network requests to the destination
		colly.Async(true),
	)

	/*
		Limiting Colly
			We might also want to place specific limits on our crawler’s behavior to be undetected by bad bot programs and get our IP banned.

			Some websites are more picky than others when it comes to the amount of traffic they allow before cutting you off.
			Generally, setting a delay of a couple seconds should keep you off the “naughty list”.

			Colly makes it easy to introduce rate limiting:
	*/

	c.Limit(&colly.LimitRule{
		// Filter domains affected by this rule
		DomainGlob: "godoc.org/*",
		// Set a delay between requests to these domains
		Delay: 1 * time.Second,
		// Add an additional random delay
		RandomDelay: 1 * time.Second,
	})

	/*
		Setting Up Callbacks

		Collectors can also have callbacks such as OnRequest and OnHTML attached to them.
		These callbacks are executed at different periods in the collection’s lifecycle (similar to React’s lifecycle methods),
		for instance, Colly calls the OnRequest method just before the collector makes an HTTP request.

		You can find a complete list of supported callbacks on Colly’s godoc page.
	*/

	// OnRequest registers a function. Function will be executed on every request made by the Collector
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// OnError registers a function. Function will be executed if an error occurs during the HTTP request.
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// OnResponse registers a function. Function will be executed on every response
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	// OnHTML registers a function. Function will be executed on every HTML element matched by the GoQuery Selector parameter.
	// GoQuery Selector is a selector used by https://github.com/PuerkitoBio/goquery
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
		fmt.Println("First column of a table row:", e.Text)
	})

	// OnScraped registers a function. Function will be executed after OnHTML, as a final part of the scraping.
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	/*
	   How Does OnHTML Work?

	   The OnHTML method allows you to register a callback for when the collector reaches a portion of a page that matches a specific HTML tag specifier.
	   For starters, we can get a callback whenever our crawler sees an <body> tag that contains an name attribute.
	*/
	c.OnHTML("body[name]", func(e *colly.HTMLElement) {
		// Extract the Class Name from the HTML Body element
		name := e.Attr("name")
		fmt.Println(name)
		link := "www.google.com"
		// Open the Link
		c.Visit(e.Request.AbsoluteURL(link))
	})

	/*
		OnHTML is one of the best callbacks.
		Rather than searching for tag attributes, you can search for the text or image.
	*/
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
}
