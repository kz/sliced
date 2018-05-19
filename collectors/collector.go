package collectors

// Collector crawls an online store for list of potentially on-sale items
// and returns them as a list of URLs.
type Collector interface {
	Process() []string
}