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
