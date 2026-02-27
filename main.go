package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	n := 5
	for page := 1; page <= n; page++ {
		fmt.Printf("Страница номер %d\n", page)
		url := fmt.Sprintf("https://books.toscrape.com/catalogue/page-%d.html", page)
		ParsePage(url)
	}
}

func ParsePage(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Ошибка Get:", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatal("Ошибка номер ", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	bookSlice := make([]string, 0)
	doc.Find("article.product_pod").Each(func(i int, s *goquery.Selection) {
		fullTitle, ok := s.Find("h3 a").Attr("title")
		if !ok {
			fmt.Printf("у id%d нет полного названия", i)
		}

		price := s.Find(".price_color").Text()

		bookSlice = append(bookSlice, fmt.Sprintf("%d. %s | Цена: %s", i+1, fullTitle, price))
	})
	for _, val := range bookSlice {
		fmt.Println(val)
	}
}
