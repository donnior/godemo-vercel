package internal

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Content struct {
	Title   string
	Content string
}

func GetContent(chapterLink string, config SiteConfig) ([]string, error) {
	res, nil := fetchContent(chapterLink, config.ContentSelector, "")
	return res, nil
}

func fetchContent(chapterLink string, contentSelector string, alertiveSelector string) ([]string, error) {

	f := DefaultWebAgentFetcher()
	doc, err := f.FetchHtmlAsDoc(chapterLink)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	fmt.Print(doc.Text())

	doc.Find("style").Remove()
	doc.Find("script").SetHtml("")

	var texts []string
	doc.Find(contentSelector).Contents().Each(func(i int, s *goquery.Selection) {
		if goquery.NodeName(s) == "#text" {
			t := strings.TrimRight(s.Text(), "\n\r")
			if len(t) > 0 {
				texts = append(texts, t)
			}
		}
	})

	return texts, nil

	// return []Chapter{
	// 	{Link: "some link", Title: "some title 2"},
	// }, nil
}
