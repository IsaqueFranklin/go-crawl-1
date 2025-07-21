package webcrawler

import (
	"fmt"
	"sync"
)

// CrawlWebpage crawls websites reading from sitesChannel.
func CrawlWebpage(wg *sync.WaitGroup, sitesChannel chan string, crawledLinksChannel chan string, pendingCountChannel chan int) {
	crawledSites := 0

	for webpageURL := range sitesChannel {
		extractContent(webpageURL, crawledLinksChannel)
		pendingCountChannel <- -1
		crawledSites++
	}

	fmt.Println("Crawled ", crawledSites, " web pages.")

	wg.Done()
}
