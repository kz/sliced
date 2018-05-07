package pipeline

// Pipeline is an interface which defines how URLs are passed from the URL
// collector to the scraper. As a result, scraping can occur concurrently
// alongside URL collection.
type Pipeline interface {
	Add(string)      // Adds a URL to the URL store
	Append([]string) // Appends a list of strings to the URL store
	Process()        // Processes the URLs in the URL store
}
