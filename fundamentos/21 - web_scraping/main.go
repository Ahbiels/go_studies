package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Book struct {
	Thumbnail       string  `json:"thumbnail"`
	Rating          uint8   `json:"rating"`
	BookName        string  `json:"book_name"`
	BookPrice       float32 `json:"book_price"`
	Currency        string  `json:"currency"`
	BookAvailable   bool    `json:"book_available"`
	BookQuantity    int     `json:"book_quantity"`
	BookDescription string  `json:"book_description"`
}

const site_url string = "https://books.toscrape.com/"

func main() {
	start := time.Now()
	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
		colly.MaxDepth(2),
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 30})
	
	detailCollector := c.Clone()
	detailCollector.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})
	var books []Book

	c.OnHTML("article[class=product_pod]", func(h *colly.HTMLElement) {
		url := h.Request.AbsoluteURL(h.ChildAttr("h3 a", "href"))
		if strings.Contains(url, "books.toscrape.com/catalogue") {
			detailCollector.Visit(url)
		}
	})
	detailCollector.OnHTML("article.product_page", func(h *colly.HTMLElement) {
		available, quantity := book_is_available(h.ChildText("div.col-sm-6 p.instock"))
		currency, price := price_clean(h.ChildText("div.col-sm-6 p.price_color"))
		book := Book{
			Thumbnail:       site_url + h.ChildAttr("div.image_container img", "src"),
			Rating:          rating_to_uint(strings.Split(h.ChildAttr("p.star-rating", "class"), " ")[1]),
			BookName:        h.ChildAttr("h3 a", "title"),
			Currency:        currency,
			BookPrice:       price,
			BookAvailable:   available,
			BookQuantity:    quantity,
			BookDescription: h.ChildText("div#product_description + p"),
		}

		books = append(books, book)
	})

	// c.OnHTML("li.next a", func(h *colly.HTMLElement) {
	// 	next_page := h.Request.AbsoluteURL(h.Attr("href"))
	// 	c.Visit(next_page)
	// })

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})
	c.Visit(site_url)
	c.Wait()
	detailCollector.Wait()
	encoder, err := json.Marshal(books)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("books.json", encoder, 0644)
	end := time.Since(start)
	fmt.Printf("\n O programa levou cerca de %v para ser executado", end)
}

func rating_to_uint(rating_str string) (rating uint8) {
	switch rating_str {
	case "One":
		rating = 1
	case "Two":
		rating = 2
	case "Three":
		rating = 3
	case "Four":
		rating = 4
	case "Five":
		rating = 5
	default:
		rating = 5
	}

	return
}

func price_clean(price_str string) (currency string, price float32) {
	currency = string(price_str[:2])
	price_stage, _ := strconv.ParseFloat(string(price_str[2:]), 64)
	price = float32(price_stage)
	return
}

func book_is_available(available_str string) (available bool, quantity int) {
	var available_slice []string
	if strings.Contains(available_str, "(") {
		available_slice = strings.Split(available_str, "(")
		if strings.TrimSpace(available_slice[0]) == "In stock" {
			available = true
		} else {
			available = false
		}
		total_book, _ := strconv.Atoi(strings.Split(strings.ReplaceAll(available_slice[1], ")", ""), " ")[0])
		quantity = total_book
	}
	return
}
