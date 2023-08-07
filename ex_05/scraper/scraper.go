package scraper

import (
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"github.com/hoangtk0100/dc-go-23/ex_05/constant"
	"time"
)

var (
	startPages = map[constant.AnimePage]string{
		constant.Toonily: "https://toonily.com/webtoon-genre/action/",
	}
)

type Scraper interface {
	Scrape() error
}

func NewScraper(page constant.AnimePage, isConcurrent bool, allowRevisit bool, timeout time.Duration, delayTime time.Duration) Scraper {
	collector := setupCollector(isConcurrent, timeout, allowRevisit)

	switch page {
	case constant.Toonily:
		return NewToonilyScraper(collector, startPages[page], delayTime)
	default:
		return nil
	}
}

func setupCollector(isConcurrent bool, timeout time.Duration, allowRevisit bool) *colly.Collector {
	c := colly.NewCollector(
		colly.IgnoreRobotsTxt(),
		colly.Async(isConcurrent),
	)

	extensions.RandomUserAgent(c)
	c.SetRequestTimeout(timeout)
	c.AllowURLRevisit = allowRevisit
	return c
}
