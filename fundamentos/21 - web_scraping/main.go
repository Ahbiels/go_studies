package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type Book struct {
	Thumbnail       string  `json:"thumbnail"`
	Rating          uint8   `json:"rating"`
	BookName        string  `json:"book_name"`
	BookPrice       float32 `json:"book_price"`
	Currency        string  `json:"currency"`
	BookAvailable   bool    `json:"book_available"`
	BookDescription string  `json:"book_description"`
}

const site_url string = "https://books.toscrape.com/"

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
		colly.CacheDir("./coursera_cache"),
	)
	detailCollector := c.Clone()
	var books []Book

	c.OnHTML("article[class=product_pod]", func(h *colly.HTMLElement) {
		url := h.Request.AbsoluteURL(h.ChildAttr("h3 a", "href"))
		if strings.Contains(url, "books.toscrape.com/catalogue") {
			detailCollector.Visit(url)
		}
		// bookCollector.Visit(h.Request.AbsoluteURL(h.ChildAttr("h3 a", "href")))

		// bookCollector.Request("GET", h.Request.AbsoluteURL(h.ChildAttr("h3 a", "href")), nil, colly.NewContext(), nil)

	})

	// c.OnHTML("article[class=product_pod] h3", func(h *colly.HTMLElement) {
	// 	url := h.Request.AbsoluteURL(h.ChildAttr("a", "href"))
	// 	if strings.Contains(url, "books.toscrape.com/catalogue") {
	// 		detailCollector.Visit(url)
	// 	}
	// })
	detailCollector.OnHTML("article.product_page", func(h *colly.HTMLElement) {
		fmt.Println(h.ChildText("div.col-sm-6 p.instock"))
		// currency, price := price_clean(h.ChildText("div.col-sm-6 p.price_color"))
		// fmt.Println(currency, price)
		// book := Book{
		// 	Thumbnail:       site_url + h.ChildAttr("div.image_container img", "src"),
		// 	Rating:          rating_to_uint(strings.Split(h.ChildAttr("p.star-rating", "class"), " ")[1]),
		// 	BookName:        h.ChildAttr("h3 a", "title"),
		// 	Currency:        currency,
		// 	BookPrice:       price,
		// 	BookAvailable:   book_is_available(h.ChildText("div.product_price p.instock")),
		// 	BookDescription: h.ChildText("div#product_description + p"),
		// }

		// books = append(books, book)
	})

	// bookCollector.OnHTML("div.page_inner", func(h *colly.HTMLElement) {
	// 	// fmt.Println(h.ChildText("p"))
	// 	// book_detailed := h.Request.AbsoluteURL(h.ChildAttr("h3 a", "title"))
	// 	// bookCollector.Visit(book_detailed)
	// })

	// c.OnHTML("li.next a", func(h *colly.HTMLElement) {
	// 	next_page := h.Request.AbsoluteURL(h.Attr("href"))
	// 	c.Visit(next_page)
	// })

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println(r.URL.String())
	// })
	// bookCollector.OnRequest(func(r *colly.Request) {
	// 	fmt.Println(r.URL.String())
	// })

	c.Visit(site_url)
	encoder, err := json.Marshal(books)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("books.json", encoder, 0644)
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
	fmt.Println(price_str)

	currency = "2"
	price = 23.3
	return
}

func book_is_available(available_str string) (available bool) {
	available_slice := strings.Split(available_str, " ")
	var quantity string
	stage_available := []string{}
	for i,c := range available_slice {
		if c == "(" {
			quantity = available_slice[i+1]
			break
		}
		stage_available = append(stage_available, c)
	}

	if available_str == "In stock" {
		available = true
	} else {
		available = false
	}

	return
}
