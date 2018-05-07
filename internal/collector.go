package internal

// Collector crawls an online store for list of potentially on-sale items
// and passes crawled URLs to the pipeline
type Collector interface {
	Process() // Crawls URLs and passes them to the pipeline
}