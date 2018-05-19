package collectors

import (
	"log"
	"net/http"
)

import "gopkg.in/xmlpath.v2"

const tedBakerProductSitemapURL = "http://www.tedbaker.com/uk/uk-googleSitemapProductDestinationMedia/sitemap.xml"
const matchingPath = "/urlset/url/loc"

type tedBakerCollector struct{}

func NewTedBakerCollector() *tedBakerCollector {
	return &tedBakerCollector{}
}

func (c *tedBakerCollector) Process() []string {
	response, err := http.Get(tedBakerProductSitemapURL)
	if err != nil {
		log.Fatal("Unable to get Ted Baker sitemap: ", err.Error())
	}
	defer response.Body.Close()

	path := xmlpath.MustCompile(matchingPath)
	root, err := xmlpath.Parse(response.Body)
	if err != nil {
		log.Fatal("Unable to parse Ted Baker sitemap: ", err.Error())
	}

	urls := make([]string, 0)
	iter := path.Iter(root)
	for iter.Next() {
		urls = append(urls, iter.Node().String())
	}

	return urls
}
