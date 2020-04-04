package model

import "github.com/sirupsen/logrus"

//AllPortfoliosViewModel - view of all portfolios
type AllPortfoliosViewModel struct {
	Portfolios []PortfolioViewModel `json:"portfolios" required:"true" description:"List of all portfolios"`
}

//PortfolioViewModel - Detailed view of portfolio and stock purchases
type PortfolioViewModel struct {
	Metadata MetadataViewModel      `json:"metadata" required:"true" description:"Portfolio metadata"`
	Stocks   map[int]StockViewModel `json:"stocks" required:"true" description:"List of all stocks held by portfolio"`
}

//StockViewModel - Detailed view of stock
type StockViewModel struct {
	Ticker        string  `json:"ticker" required:"true" description:"Ticker symbol" example:"CP.TO"`
	Name          string  `json:"name" required:"true" description:"Stock name" example:"Canadian Pacific Railway Limited"`
	Quantity      int     `json:"quantity" required:"true" description:"Number of shares purchased" example:"100"`
	PurchasePrice float64 `json:"purchasePrice" required:"true" description:"Price at time of purchase" example:"10000.00"`
	CurrentPrice  float64 `json:"currentPrice" required:"true" description:"Latest available price per share" example:"105.00"`
	PurchaseDate  string  `json:"purchaseDate" required:"true" description:"Date of purchase" example:"02/03/2020"`
	Currency      string  `json:"currency" required:"true" description:"Purchase currency" example:"CAD"`
}

func (model StockViewModel) GetFields() logrus.Fields {
	return logrus.Fields{
		"ticker":        model.Ticker,
		"name":          model.Name,
		"quantity":      model.Quantity,
		"purchasePrice": model.PurchasePrice,
		"currentPrice":  model.CurrentPrice,
		"purchaseDate":  model.PurchaseDate,
		"currency":      model.Currency,
	}
}

type MetadataViewModel struct {
	ID          string `json:"id" required:"true" description:"Portfolio ID" example:"123884"`
	CreateTime  string `json:"createTime" required:"true" description:"Time stamp when portfolio was created. Provided in UTC." example:"02/01/2020 11:12:00"`
	LastUpdated string `json:"lastUpdated" required:"true" description:"Time in which portfolio was last updated" example:"02/01/2020 11:12:00"`
}
