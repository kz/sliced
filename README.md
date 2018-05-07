# Sliced - Scraping Online Store Discounts

Sliced crawls and aggregates discounts from online designer fashion stores. Currently, only stores in the United Kingdom are supported.

Sliced is designed to be modular and platform agnostic. On a higher level, this allows for new stores to be easily supported. On a lower level, the platform agnosticism makes it possible to easily switch the database and cache systems and adapt the application between local and cloud environments. For example, the link between the URL/sitemap scraper and individual page processor can be switched between Go channels and supported third party message queues (Kafka, Google Cloud Pub/Sub, etc.).

**Why do this project?** Sliced is a project for me to interact with different technologies (hence the decision to make it platform agnostic) while building a scalable application which also outputs potentially useful data.

## Components

The components of the scraper are as follows:

- The **pipeline** manages the flow of data of collected URLs so that scraping items is not compulsory after collecting URLs.
- The **URL collector** collects all (potentially on-sale) item URLs from an online store (ideally from a sitemap or a page listing all sale items) and exports them via the pipeline.
- The **item scraper** takes in a URL, scrapes its item page to check if it is on sale, and passes the sale item to the data store. 
- The **data store** stores the sale items.