package collectors

type exampleCollector struct {
	maxConcurrency int // Maximum number of scraper goroutines which can be created
}

func NewExampleCollector(maxConcurrency int) *exampleCollector {
	return &exampleCollector{
		maxConcurrency: maxConcurrency,
	}
}

func (c *exampleCollector) Process() []string {
	return []string{}
}
