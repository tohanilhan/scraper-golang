## About

This is a simple scraper that scrapes data from [dolap.com](https://dolap.com) and stores that information under products folder as a json file. It is written in Go and uses [Colly](www.github.com/gocolly/colly) as a scraping framework.
It uses [godotenv](www.github.com/joho/godotenv) to load environment variables from .env file and [carrlos0/env]("github.com/caarlos0/env/v6") to load environment variables into structs.

## Usage

You can run the project directly by running the below script from the cart-api directory inside the project:

    go run main.go

## Environment Variables

You can use the .env file inside this project. You can change the values of the variables inside the .env file. The variables are:

- **URL:** The url of the website that you want to scrape. (In our case, it is dolap.com because this scraper configured to scrape data from dolap.com)


- **PRODUCT_NAME:** The name of the product that you want to scrape. (gozluk,saat etc.)