package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strconv"
	"strings"
)

type Watch struct {
	UserID				string	`json:"User_Id"`
	UserName 			string	`json:"User-Name"`
	UserProfile			string	`json:"User-Profile"`
	ProductID			string	`json:"Product-ID"`
	ProductProfile		string	`json:"Product_Link"`
	WatchName 			string	`json:"Product-Name"`
	ProductDetail 		string	`json:"Product-Detail"`
	OldPrice 			string	`json:"Old-Price"`
	NewPrice 			string	`json:"New-Price"`
	LikesCount			int 	`json:"Like-Count"`
	CommentsCount		int 	`json:"Comments-Count"`
	Comments			[]WatchComment `json:"Comments"`

}
type WatchComment struct {
	UserName 			string	`json:"User-Name"`
	Comment 			  string
	Replies 			[]SubComment `json:"Replies"`
}
type SubComment struct {
	SubComment 			  	[]string
}

func main() {

	allWatches := []Watch{}
	c:= colly.NewCollector(
		colly.AllowedDomains("dolap.com","www.dolap.com"),
	)
	detailCollector := c.Clone()

	c.OnHTML(`.col-holder`, func(e *colly.HTMLElement) {
		productURL := e.ChildAttr("div.img-block > a", "href")
		productURL = e.Request.AbsoluteURL(productURL)

		detailCollector.Visit(productURL)
	})

	detailCollector.OnHTML(`.holder`, func(e *colly.HTMLElement) {
		tempProduct := Watch{}
		var err error
		tempProduct.ProductID = e.ChildAttr("div.likes-block > a", "data-product-id")
		//fmt.Println("Getting data for: ", tempProduct.ProductID)
		tempProduct.ProductProfile = e.ChildAttr("div.container > input", "value")
		tempProduct.UserID = e.ChildAttr("div.person-img > span", "data-imageid")
		tempProduct.UserName = e.ChildText("div.title-stars-block > a")
		tempProduct.UserProfile = e.ChildAttr("div.title-stars-block > a", "href")
		tempProduct.WatchName = e.ChildText("div.title-holder > h1")
		tempProduct.OldPrice = e.ChildText("div.price-block > div.price-detail > span.disc-price")
		tempProduct.NewPrice = e.ChildText("div.price-block > div.price-detail > span.price")
		tempProduct.ProductDetail = e.ChildText("div.remarks-block > p")
		tempProduct.LikesCount, err = strconv.Atoi(e.ChildAttr("div.likes-block > a", "data-product-like-count"))
		tempProduct.CommentsCount, err = strconv.Atoi(e.ChildText("div.comments-block > h2 > span.comment-count"))
		if err != nil {
			log.Println("No comments found")
		}

		e.ForEach("ul.comments-list ", func(i int, element *colly.HTMLElement) {
			tempComment := WatchComment{}
			tempComment.UserName = element.ChildText("ul.comments-list > li > div.comment-detail > div.comment-head > a.name")
			tempComment.Comment = element.ChildText("ul.comments-list > li > div.comment-detail > div.comment-holder > p")

			element.ForEach("ul.replies-list", func(i int, element *colly.HTMLElement) {
				tempReply := SubComment{}
				a := element.ChildText("ul.replies-list > li > div.comment-detail > div.comment-holder > p")
				if a != "" {
					b := strings.Split(a,"@")
					copy(b[i:], b[i+1:]) // Shift a[i+1:] left one index.
					b[len(b)-1] = ""     // Erase last element (write zero value).
					b = b[:len(b)-1]     // Truncate slice.

					tempReply.SubComment = b
					tempComment.Replies = append(tempComment.Replies, tempReply)
				}
			})

			tempProduct.Comments = append(tempProduct.Comments, tempComment)

		})

		allWatches = append(allWatches, tempProduct)
	})

	detailCollector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting : ", request.URL.String())
	})

	for i := 1; i <= 277; i++ {
		c.Visit("https://dolap.com/saat?sayfa=" + strconv.Itoa(i))
	}


	jsonRes, err := json.MarshalIndent(allWatches, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	// Setting the filename
	fileName := "dolap_new.json"

	// Calling the writeFile function to write the data into the json file
	writeFile(fileName, string(jsonRes))

	fmt.Println("Done !")
}



// Function for writing the json file with the data
func writeFile(fileName string, data string) {
	f, openerr := os.Create(fileName)

	if openerr != nil {
		log.Fatal(openerr)
	}

	defer f.Close()

	_, err2 := f.WriteString(data)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Writing file:", fileName)
}
