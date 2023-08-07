package scraper

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/hoangtk0100/dc-go-23/ex_05/util"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"
)

type toonilyScraper struct {
	collector *colly.Collector
	startPage string
	delayTime time.Duration
	mutex     sync.Mutex
}

func NewToonilyScraper(collector *colly.Collector, startPage string, delayTime time.Duration) *toonilyScraper {
	return &toonilyScraper{
		collector: collector,
		startPage: startPage,
		delayTime: delayTime,
		mutex:     sync.Mutex{},
	}
}

func (s *toonilyScraper) Scrape() error {
	file := util.OpenFile("comics.json")

	// Goto categories
	s.collector.OnHTML("div.genres__collapse ul li", func(h *colly.HTMLElement) {
		url := h.Request.AbsoluteURL(h.ChildAttr("a", "href"))
		if err := s.collector.Visit(url); err != nil {
			log.Println(err)
		}
	})

	// Goto the next pages
	s.collector.OnHTML("a.nextpostslink", func(h *colly.HTMLElement) {
		if err := s.collector.Visit(h.Request.AbsoluteURL(h.Attr("href"))); err != nil {
			log.Println(err)
		}
	})

	// Visit comics
	s.collector.OnHTML("div.item-summary", func(h *colly.HTMLElement) {
		url := h.Request.AbsoluteURL(h.ChildAttr("a", "href"))
		if err := s.collector.Visit(url); err != nil {
			log.Println(err)
		}
	})

	// Get comic details
	s.collector.OnHTML("div.site-content", func(h *colly.HTMLElement) {
		time.Sleep(s.delayTime)

		isComicPage, publisher, status := getPublisherStatus(h, "div.post-status > div.post-content_item > div.summary-content")
		if !isComicPage {
			return
		}

		var summary strings.Builder
		h.ForEach("div.summary__content > p", func(_ int, e *colly.HTMLElement) {
			contentOrigin := regexp.MustCompile(`\n`)
			contentConverted := contentOrigin.ReplaceAllString(e.Text, "<br/>")
			summary.WriteString(fmt.Sprintf("<p>%v</p>", contentConverted))
		})

		altName := strings.Split(h.ChildText("div.manga-info-row > div.post-content_item > div.summary-content"), " \n")[0]
		comic := Comic{
			Name:        h.ChildText("div.post-title"),
			Rate:        h.ChildText("span#averagerate"),
			RatingCount: h.ChildText("span#googlecountrate"),
			Views:       getViews(h.ChildText("div.manga-rate-view-comment > div.item")),
			URL:         h.Request.URL.String(),
			AltName:     altName,
			Publisher:   publisher,
			Status:      status,
			Authors:     getMapFromAHrefs(h, "div.author-content > a"),
			Artists:     getMapFromAHrefs(h, "div.artist-content > a"),
			Genres:      getMapFromAHrefs(h, "div.genres-content > a"),
			Tags:        getMapFromAHrefs(h, "div.wp-manga-tags-list > a"),
			Summary:     summary.String(),
			Chapters:    make(map[string]Chapter),
		}

		h.ForEach("div.listing-chapters_wrap.cols-2 > ul > li", func(_ int, e *colly.HTMLElement) {
			chapter := Chapter{
				Name:          e.ChildText("a"),
				URL:           e.ChildAttr("a", "href"),
				PublishedDate: e.ChildText("span > i"),
			}

			comic.Chapters[chapter.Name] = chapter
		})

		if err := util.WriteFile(file, comic); err != nil {
			log.Fatalln(err)
		}
	})

	s.collector.OnError(func(h *colly.Response, err error) {
		log.Printf("error (%s): %v\n", h.Request.URL, err)
	})

	if err := s.visitPage(s.startPage); err != nil {
		return err
	}

	if err := util.CloseFile(file); err != nil {
		return err
	}

	return nil
}

func (s *toonilyScraper) visitPage(url string) error {
	err := s.collector.Visit(url)
	if err != nil {
		log.Println(err)
	}

	s.collector.Wait()

	return err
}

func getViews(input string) string {
	start := strings.LastIndex(input, ")")
	return strings.TrimSpace(input[start+1:])
}

func getMapFromAHrefs(h *colly.HTMLElement, path string) map[string]string {
	result := make(map[string]string)
	h.ForEach(path, func(_ int, m *colly.HTMLElement) {
		result[m.Text] = m.Request.AbsoluteURL(m.Attr("href"))
	})

	return result
}

func getPublisherStatus(h *colly.HTMLElement, path string) (bool, map[string]string, string) {
	publisherStatus := strings.Split(h.ChildText(path), " \n")
	if len(publisherStatus) != 2 {
		return false, nil, ""
	}

	publisher := make(map[string]string)
	if publisherStatus[0] != "Updating" {
		publisher[publisherStatus[0]] = h.ChildAttr(fmt.Sprintf("%s > a", path), "href")
	} else {
		publisher[publisherStatus[0]] = ""
	}

	return true, publisher, publisherStatus[1]
}
