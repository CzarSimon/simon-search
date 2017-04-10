package main

import (
	"time"

	"github.com/hishboy/gocommons/lang"
)

// Links is the type to store urls that the crawler has and will scrape
type Links struct {
	Queue *lang.Queue
	Seen  map[string]time.Time
}

// QueueLink ads a new link to the queue if it has not been previously encounered
func (links *Links) QueueLink(link string) {
	if _, ok := links.Seen[link]; !ok {
		links.Queue.Push(link)
	}
}

//AddToSeen ads link to the set of seen links
func (links *Links) AddToSeen(link string) {
	links.Seen[link] = time.Now().UTC()
}

// HandlePageLinks queue the newly encountered links and sets the visited page url to seen
func (links *Links) HandlePageLinks(page Page) {
	for _, link := range page.Links {
		links.QueueLink(link)
	}
	links.AddToSeen(page.URL)
}

// InitialLinks creates a Links type with the
func InitialLinks(seedLinks []string) *Links {
	links := &Links{
		Queue: lang.NewQueue(),
		Seen:  make(map[string]time.Time),
	}
	for _, link := range seedLinks {
		links.Queue.Push(link)
	}
	return links
}
