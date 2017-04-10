package main

import "fmt"

// Page is the contrainer type for information scraped from a webpage
type Page struct {
	Title, Text, URL string
	Links            []string
}

func createPage(title, url string, links []string) Page {
	return Page{
		Title: title,
		Text:  "text",
		URL:   url,
		Links: links,
	}
}

func handlePage(page Page) {
	fmt.Println(page.Title)
	fmt.Println(page.Text)
	fmt.Println(page.URL)
	fmt.Println("Links:", len(page.Links))
}
