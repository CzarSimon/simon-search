package main

import (
  "github.com/rainycape/cld2"
  "github.com/surgebase/porter2"
  "fmt"
  "github.com/anaskhan96/soup"
  "strings"
  "regexp"
)

func scrapePage(url string) (string, error) {
  html, err := soup.Get(url)
  if err != nil {
    return "Failed to scrape", err
  }
  return html, nil
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

func main() {
  url := "https://news.ycombinator.com/"
  res, err := scrapePage(url)
  checkErr(err)
  fmt.Println(res)
}
