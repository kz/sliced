# Sliced - Scraping Online Store Discounts

Sliced crawls and aggregates discounts from online designer fashion stores. Work in progress.

## Components

The components of the scraper are as follows:

- The **URL collector** collects all (potentially on-sale) item URLs from an online store (ideally from a sitemap or a page listing all sale items) and returns a list of URLs.
- The **item scraper** takes in a URL, scrapes its item page to check if it is on sale, and passes the sale item to the data store.