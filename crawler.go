package main

import (
	"log"
	"time"
)

// Crawl continously scrapes pages, putting new links on the link queue and storing
func Crawl(links *Links) {
	for {
		if links.Queue.Len() < 1 {
			time.Sleep(1000)
			continue
		}
		CrawlURL(links.GetLink(), links)
	}
}

// CrawlURL performs crawl on a specified url
func CrawlURL(url string, links *Links) {
	page, err := ScrapePage(url)
	if err != nil {
		log.Println(err)
	} else {
		links.HandlePageLinks(page)
		handlePage(page)
	}
}

func main() {
	seedLinks := []string{"https://xkcd.com/", "https://news.ycombinator.com/", "https://www.nytimes.com"}
	links := InitialLinks(seedLinks)
	Crawl(links)
	links.Report()
}
