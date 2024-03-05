package pkg

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/donnior/godemo-vercel/util"
	logger "github.com/sirupsen/logrus"
)

type Chapter struct {
	Link  string
	Title string
}

type SiteConfig struct {
	Domain          string
	ChapterSelector string
	ContentSelector string
}

func ListChapter(bookLink string, config SiteConfig) ([]Chapter, error) {
	res, nil := fetchChapter(bookLink, config.ChapterSelector, "")
	return res, nil
}

func fetchChapter(bookLink string, chapterSelector string, alertiveSelector string) ([]Chapter, error) {

	f := DefaultWebAgentFetcher()
	doc, err := f.FetchHtmlAsDoc(bookLink)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	var topics []Chapter

	doc.Find(chapterSelector).Each(func(i int, s *goquery.Selection) { //迭代查询
		s.Find("a").Each(func(j int, a *goquery.Selection) {
			link, ok := a.Attr("href")
			if ok {
				c := Chapter{
					Title: a.Text(),
					Link:  util.RelativePathToAbsolutePath(link, bookLink),
				}
				topics = append(topics, c)
			} else {
				logger.Println("Not find topic with:", a)
			}
		})
	})

	return topics, nil

	// return []Chapter{
	// 	{Link: "some link", Title: "some title 2"},
	// }, nil
}
