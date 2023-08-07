package main

import (
	"github.com/hoangtk0100/dc-go-23/ex_05/constant"
	"github.com/hoangtk0100/dc-go-23/ex_05/scraper"
	"github.com/hoangtk0100/dc-go-23/ex_05/util"
	"log"
	"time"
)

func main() {
	defer util.Timer("main")()

	s := scraper.NewScraper(constant.Toonily, true, false, time.Second*10, time.Millisecond*2)

	err := s.Scrape()
	if err != nil {
		log.Fatal(err)
	}
}
