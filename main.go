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
	// http://www.zgpingshu.com/down/5109/
	doc.Find("div a").Each(func(i int, s0 *goquery.Selection) {
		if val, ok := s0.Attr("id"); ok && val == "down" {
			if val, ok := s0.Attr("href"); ok {
				 s0.Children().Each(func(i int, s0 *goquery.Selection) {
					img, _ := s0.Attr("img")
					fmt.Println("img:", img, i)
				})

				img, _ := s0.Attr("img")
				fmt.Println("img:", img, i)

				fmt.Println(val, i)
				return
			}
		}

		/*
		s0.Children().Each(func(i int, s *goquery.Selection) {
			if val, ok := s.Attr("id"); ok {
				if val == "down" {
					s2 := s.Next()
					if val, ok := s.Attr("href"); ok {
						img, _ := s2.Attr("img")
						fmt.Println(val, img)
						return
					}
				}
			}
		})
		*/


		

	})

}
