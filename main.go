package webcrawler

import (
	"sync"
)

// CrawlerMain creates all the channels and triggers BO to start crawling.
func CrawlerMain() {
	sitesChannel := make(chan string)
	crawledLinksChannel := make(chan string)
	pendingCountChannel := make(chan int)

	siteToCrawl := "https://theuselessweb.com/"

	go func() {
		crawledLinksChannel <- siteToCrawl
	}()

	var wg sync.WaitGroup

	go ProcessCrawledLinks(sitesChannel, crawledLinksChannel, pendingCountChannel)
	go MonitorCrawling(sitesChannel, crawledLinksChannel, pendingCountChannel)

	var numCrawlerThreads = 50
	for i := 0; i < numCrawlerThreads; i++ {
		wg.Add(1)
		go CrawlWebpage(&wg, sitesChannel, crawledLinksChannel, pendingCountChannel)
	}

	wg.Wait()
}
