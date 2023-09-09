package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type quote struct {
	quote  string
	author string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/116.0")
		fmt.Println("visting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response code", r.StatusCode)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error", err.Error())
	})

	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		div := h.DOM
		quote := div.Find(".text").Text()
		author := div.Find(".author").Text()
		fmt.Printf("quote: %s\nBy %s\n\n", quote, author)

	})

	// c.OnHTML(".text", func(h *colly.HTMLElement) {
	// 	fmt.Println("quote", h.Text)
	// })
	// c.OnHTML(".author", func(h *colly.HTMLElement) {
	// 	fmt.Println("author", h.Text)
	// })
	c.Visit("http://quotes.toscrape.com")
}
