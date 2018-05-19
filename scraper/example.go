package scraper

import "github.com/kz/sliced/drivers/database"

type exampleScraper struct {
	db database.Database
}

func NewExampleScraper(db database.Database) *exampleScraper {
	return &exampleScraper{
		db: db,
	}
}

func (s *exampleScraper) Process(url string) {
	s.db.AddItem(map[string]interface{}{})
}
