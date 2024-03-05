package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	// "github.com/donnior/godemo-vercel/util"
	log "github.com/sirupsen/logrus"
	logger "github.com/sirupsen/logrus"
	// "github.com/donnior/godemo-vercel/internal"
)

func ChapterHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := ListChapter("some book link", SiteConfig{
		Domain:          "some domain",
		ChapterSelector: "some selector",
		ContentSelector: "some selector",
	})
	if err != nil {
		// w.Write("some error")
		fmt.Print("error")
	}
	// w.Write([]byte(res))
	fmt.Print(res)
	json.NewEncoder(w).Encode(res)
}

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
					Link:  RelativePathToAbsolutePath(link, bookLink),
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

type Fetcher struct {
	Header map[string]string
}

var customizedHeader map[string]string

func DefaultWebAgentFetcher() *Fetcher {
	if customizedHeader != nil {
		log.Info("Init fetcher with customized header")
		return &Fetcher{Header: customizedHeader}
	}
	log.Info("Init fetcher with default pc header")
	return &Fetcher{Header: defaultPCAgentHeader()}
}

func DefaultMobileAgentFetcher() *Fetcher {
	if customizedHeader != nil {
		log.Info("Init fetcher with customized header")
		return &Fetcher{Header: customizedHeader}
	}
	log.Info("Init fetcher with default mobile header")
	return &Fetcher{Header: defaultMobileAgentHeader()}
}

func NewAgentFetcher(headers map[string]string) *Fetcher {
	return &Fetcher{Header: headers}
}

func SetHeader(header map[string]string) {
	customizedHeader = header
}

const (
	defaultMobileAgent string = "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) " +
		"AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1"
	defaultPCAgent string = "Mozilla/5.0 (Windows NT 6.1; WOW64)" +
		" AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36"
	defaultAccept string = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"
)

func defaultMobileAgentHeader() map[string]string {
	header := map[string]string{
		"User-Agent": defaultMobileAgent,
		"Accept":     defaultAccept,
		// "Referer":    "https://www.baidu.com",
		// "Referer": "https://www.mumumh.com/",
	}

	return header
}

func defaultPCAgentHeader() map[string]string {
	header := map[string]string{
		"User-Agent": defaultPCAgent,
		"Accept":     defaultAccept,
		// "Referer":    "https://www.baidu.com",
	}

	return header
}

func (fetcher *Fetcher) FetchHtml(url string) string {
	resp, err := fetcher.doFetch(url)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	return string(body)
}

func (fetcher *Fetcher) doFetch(url string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	for key, value := range fetcher.Header {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	return resp, err
}

func (fetcher *Fetcher) FetchHtmlAsDoc(url string) (*goquery.Document, error) {
	resp, err := fetcher.doFetch(url)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	// Convert the designated charset HTML to utf-8 encoded HTML.
	// `charset` being one of the charsets known by the iconv package.
	// utfBody, err := iconv.NewReader(resp.Body, "gbk", "utf-8")
	// if err != nil {
	// 	// handler error
	// 	fmt.Println(err)
	// }
	// fmt.Println("got utf8: ", utfBody)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	return doc, err
}

func RelativePathToAbsolutePath(href string, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return " "
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return " "
	}
	return baseUrl.ResolveReference(uri).String()

}
