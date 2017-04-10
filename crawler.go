package main

import "log"

// Crawl continously scrapes pages, putting new links on the link queue and storing
func Crawl(links *Links) {
	var url string
	url = links.Queue.Poll().(string)
	page, err := scrapePage(url)
	if err != nil {
		log.Println(err)
	} else {
		links.HandlePageLinks(page)
	}
}

func main() {
	seedLinks := []string{"https://xkcd.com/", "https://news.ycombinator.com/", "https://www.nytimes.com"}
	links := InitialLinks(seedLinks)
	Crawl(links)
}
