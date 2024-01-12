package parser

import (
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/kdrkrgz/stockme/model"
)

func BorsaDovizWebParser(resp *http.Response) (*model.Stock, error) {
	// parse borsa.doviz.com data
	// goquery

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	var symbolAndName []string
	var price string
	stock := &model.Stock{}
	doc.Find(".currency-card").Each(func(i int, s *goquery.Selection) {

		symbolAndName = strings.Split(s.Find(".page-title").Text(), "-")
		price = s.Find(".text-xl").Text()
		// fmt.Printf("Data %d: %s ---- %s: %s\n", i, symbolAndName[0], symbolAndName[1], price)
	})
	defer resp.Body.Close()
	stock.Date = time.Now().Format("2006-01-02 15:04")
	stock.Name = symbolAndName[1]
	stock.Symbol = symbolAndName[0]
	stock.BuyPrice = price
	stock.SellPrice = price

	return stock, nil
}
