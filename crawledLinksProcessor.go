package main

// ProcessCrawledLinks reads the crawled links and adds unique links to sitesChannel
func ProcessCrawledLinks(sitesChannel chan string, crawledLinksChannel chan string, pendingCountChannel chan int) {
	foundUrls := make(map[string]bool)

	for cl := range crawledLinksChannel {
		if !foundUrls[cl] {
			foundUrls[cl] = true
			pendingCountChannel <- 1
			sitesChannel <- cl
		}
	}
}
