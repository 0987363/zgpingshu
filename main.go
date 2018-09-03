package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape(url string) *goquery.Document {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc

}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Unable read args.")
	}

	doc := ExampleScrape(os.Args[1])

// <div id="content" class="grid-8-1">
	// Find the review items
	doc.Find("div").Each(func(i int, s0 *goquery.Selection) {
		s0.Find("a").Each(func(i int, s *goquery.Selection) {
			if val, ok := s.Attr("id"); ok {
				if val == "down" {
					if val, ok := s.Attr("href"); ok {
						img, _ := s0.Attr("img")
						fmt.Println(val, img)
						return
					}
				}
			}
		})


		

	})

}
