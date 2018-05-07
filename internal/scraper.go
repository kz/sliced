package internal

// Scraper processes a single item and adds it to the data store
type Scraper interface {
	Process(url string) // Crawls a single URL and passes data to the data store
}