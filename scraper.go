package main

import (
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/rainycape/cld2"
	"github.com/surgebase/porter2"
)

// ScrapePage retrives and parses the html document pointed to by an URL
func ScrapePage(url string) (Page, error) {
	html, err := soup.Get(url)
	if err != nil {
		return Page{}, err
	}
	doc := soup.HTMLParse(html)
	links := parseLinks(&doc, url)
	title := parseTitle(&doc)
	text := parseText(title, &doc)
	return createPage(title, text, url, links), nil
}

func parseTitle(doc *soup.Root) string {
	return doc.Find("title").Text()
}

func parseText(text string, doc *soup.Root) string {
	for _, tag := range getTags() {
		text += " " + getTagText(tag, doc)
	}
	return formatWords(text)
}

func joinWords(words []string) string {
	textBody := strings.Join(words, " ")
	textBody = strings.TrimSpace(textBody)
	re := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	return re.ReplaceAllString(textBody, "")
}

func getTagText(tag string, doc *soup.Root) string {
	text := ""
	for _, node := range doc.FindAll(tag) {
		text += " " + node.Text()
	}
	return text
}

func getTags() []string {
	return []string{"p", "h1", "h2", "h3", "h4", "h5", "a"}
}

func parseLinks(doc *soup.Root, url string) []string {
	links := doc.FindAll("a")
	parsedLinks := make([]string, len(links))
	for index, link := range links {
		parsedLinks[index] = parseLink(link, url)
	}
	return parsedLinks
}

func parseLink(link soup.Root, url string) string {
	href := link.Attrs()["href"]
	if isFullURL(href) {
		return href
	}
	return mergeLink(url, href)
}

func mergeLink(baseURL, href string) string {
	char := "/"
	if strings.HasSuffix(baseURL, char) {
		baseURL = baseURL[:len(baseURL)-len(char)]
	}
	if strings.HasPrefix(href, char) {
		href = href[len(char):]
	}
	return baseURL + char + href
}

func isFullURL(link string) bool {
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

func formatWords(textBody string) string {
	words := strings.Split(textBody, " ")
	cleanWords := make([]string, len(words))
	re := regexp.MustCompile("\\W")
	for index, word := range words {
		cleanWords[index] = stem(removePunctuation(word, re))
	}
	return joinWords(cleanWords)
}

func removePunctuation(str string, re *regexp.Regexp) string {
	return re.ReplaceAllString(str, "")
}

func stem(word string) string {
	return porter2.Stem(word)
}
