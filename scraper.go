package main

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/rainycape/cld2"
	"github.com/surgebase/porter2"
)

func scrapePage(url string) (Page, error) {
	html, err := soup.Get(url)
	if err != nil {
		return Page{}, err
	}
	doc := soup.HTMLParse(html)
	links := parseLinks(doc, url)
	text := "title"
	return createPage(text, url, links), nil
}

func parseLinks(doc soup.Root, url string) []string {
	links := doc.FindAll("a")
	parsedLinks := make([]string, len(links))
	for index, link := range links {
		parsedLinks[index] = parseLink(link, url)
		fmt.Println(parseLink(link, url))
	}
	return parsedLinks
}

func parseLink(link soup.Root, url string) string {
	href := link.Attrs()["href"]
	if isFullUrl(href) {
		return href
	} else {
		return url + href
	}
}

func isFullUrl(link string) bool {
	candidateURL, err := url.Parse(link)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return candidateURL.Host != ""
}

func isEnglish(text string) bool {
	lang := cld2.Detect(text)
	return lang == "en"
}

func formatWords(textBody string) []string {
	words := strings.Split(textBody, " ")
	cleanWords := make([]string, len(words))
	re := regexp.MustCompile("\\W")
	for index, word := range words {
		cleanWords[index] = stem(removePunctuation(word, re))
	}
	return cleanWords
}

func removePunctuation(str string, re *regexp.Regexp) string {
	return re.ReplaceAllString(str, "")
}

func stem(word string) string {
	return porter2.Stem(word)
}
