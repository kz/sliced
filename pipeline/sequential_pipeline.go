package pipeline

import (
	"sync"

	"../scraper"
)

// sequentialPipeline collects all the URLs, then passes them to the scraper by
// spawning a limited number of goroutines.
type sequentialPipeline struct {
	urls           []string    // List of URLs to be processed by the scraper
	urlsMutex      *sync.Mutex // Mutex to prevent races in the urls slice
	maxConcurrency int         // Maximum number of scraper goroutines which can be created
	scraper        scraper.Scraper
}

// Creates a new instance of sequentialPipeline with a scraper goroutine limit
func NewSequentialPipeline(maxConcurrency int, scraper scraper.Scraper) sequentialPipeline {
	return sequentialPipeline{
		urls:           []string{},
		urlsMutex:      &sync.Mutex{},
		maxConcurrency: maxConcurrency,
		scraper:        scraper,
	}
}

func (p *sequentialPipeline) Add(url string) {
	p.urlsMutex.Lock()
	p.urls = append(p.urls, url)
	p.urlsMutex.Unlock()
}

func (p *sequentialPipeline) Append(urls []string) {
	p.urlsMutex.Lock()
	p.urls = append(p.urls, urls...)
	p.urlsMutex.Unlock()
}

func (p *sequentialPipeline) Process() {
	p.urlsMutex.Lock()

	sem := make(chan bool, p.maxConcurrency)
	for _, url := range p.urls {
		sem <- true
		go func(string) {
			defer func() {
				<-sem
			}()

			p.scraper.Process(url)
		}(url)
	}

	p.urlsMutex.Unlock()
}
