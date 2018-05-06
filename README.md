# Sliced - Scraping Online Store Discounts

Sliced crawls and aggregates discounts from online designer fashion stores. Currently, only stores in the United Kingdom are supported.

Sliced is designed to be modular and platform agnostic. On a higher level, this allows for new stores to be easily supported. On a lower level, the platform agnosticism makes it possible to easily switch the database and cache systems and adapt the application between local and cloud environments. For example, the link between the URL/sitemap scraper and individual page processor can be switched between Go channels and supported third party message queues (Kafka, Google Cloud Pub/Sub, etc.).

**Why do this project?** Sliced is a project for me to interact with different technologies (hence the decision to make it platform agnostic) while building a scalable application which also outputs potentially useful data.