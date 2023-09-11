package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/tohanilhan/scraper-golang/models"
	"github.com/tohanilhan/scraper-golang/vars"
	"log"
	"strconv"
)

func GetPageCount(c *colly.Collector) int {
	countStr := ""
	c.OnHTML("#main > div > div > div.sidebar-content-block.row > div.col-sm-9.col-sm-push-3 > div > ul.pagination.other > li:nth-child(7)", func(e *colly.HTMLElement) {
		countStr = e.Text
	})

	err := c.Visit(vars.Config.URL + vars.Config.ProductName)
	if err != nil {
		return 0
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		return 0
	}

	return count
}

func GetProducts(collector *colly.Collector, pageCount int) {
	var products []models.Product
	var err error
	// Create another collector to scrape product details
	detailCollector := collector.Clone()

	// Attach callbacks to the collector
	collector.OnHTML(`.col-holder`, func(e *colly.HTMLElement) {
		productURL := e.ChildAttr("div.img-block > a", "href")
		productURL = e.Request.AbsoluteURL(productURL)

		err := detailCollector.Visit(productURL)
		if err != nil {
			return
		}
	})

	// Extract details of the product
	detailCollector.OnHTML(`.holder`, func(e *colly.HTMLElement) {
		product := &models.Product{}

		product.ProductID = e.ChildAttr("div.likes-block > a", "data-product-id")
		//fmt.Println("Getting data for: ", product.ProductID)
		product.ProductProfile = e.ChildAttr("div.container > input", "value")
		product.UserID = e.ChildAttr("div.person-img > span", "data-imageid")
		product.UserName = e.ChildText("div.title-stars-block > a")
		product.UserProfile = e.ChildAttr("div.title-stars-block > a", "href")
		product.ProductName = e.ChildText("div.title-holder > h1")
		product.OldPrice = e.ChildText("div.price-block > div.price-detail > span.disc-price")
		product.NewPrice = e.ChildText("div.price-block > div.price-detail > span.price")
		product.ProductDetail = e.ChildText("div.remarks-block > p")
		product.LikesCount, _ = strconv.Atoi(e.ChildAttr("div.likes-block > a", "data-product-like-count"))
		product.CommentsCount, _ = strconv.Atoi(e.ChildText("div.comments-block > h2 > span.comment-count"))

		// Extract comments
		e.ForEach("ul.comments-list ", func(i int, element *colly.HTMLElement) {
			comment := models.WatchComment{}
			comment.UserName = element.ChildText("ul.comments-list > li > div.comment-detail > div.comment-head > a.name")
			comment.Comment = element.ChildText("ul.comments-list > li > div.comment-detail > div.comment-holder > p")

			// Extract replies
			element.ForEach("ul.replies-list", func(i int, element *colly.HTMLElement) {

				subCommentStr := element.ChildText("ul.replies-list > li > div.comment-detail > div.comment-holder > p")
				if subCommentStr != "" {

					comment.Replies = append(comment.Replies, subCommentStr)
				}
			})
			product.Comments = append(product.Comments, comment)
		})
		products = append(products, *product)
	})

	detailCollector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting : ", request.URL.String())
	})

	for i := 1; i <= pageCount; i++ {
		err := collector.Visit(vars.Config.URL + vars.Config.ProductName + "?sayfa=" + strconv.Itoa(i))
		if err != nil {
			return
		}
	}

	jsonRes, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	// Setting the filename
	fileName := vars.Config.ProductName + ".json"

	// Calling the writeFile function to write the data into the json file
	WriteFile(fileName, string(jsonRes))

}
