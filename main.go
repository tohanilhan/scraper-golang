package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
	"github.com/tohanilhan/scraper-golang/utils"
	"github.com/tohanilhan/scraper-golang/vars"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	err = env.Parse(&vars.Config)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func main() {

	// Instantiate collector with allowed domains
	c := colly.NewCollector(
		colly.AllowedDomains("dolap.com", "www.dolap.com"),
	)

	// Create another collector to scrape page count detail
	pageCounterCollector := c.Clone()

	// Get page count
	pageCount := utils.GetPageCount(pageCounterCollector)

	// Get all products
	utils.GetProducts(c, pageCount)

	fmt.Println("Done !")
}
