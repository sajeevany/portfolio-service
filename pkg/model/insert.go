package model

import "github.com/sirupsen/logrus"

//Defines models as required for the portfolio creation routes routes
type PortfolioCreateModel struct {
	Stocks []StockViewModel `json:"stocks" required:"true" description:"Stocks owned by portfolio"`
}

//GetFields - Returns stock view model as a slice of logrus fields
func (model PortfolioCreateModel) GetFields() logrus.Fields {

	stocks := make([]logrus.Fields, len(model.Stocks))
	for index, stock := range model.Stocks {
		stocks[index] = stock.GetFields()
	}

	return logrus.Fields{
		"stocks": stocks,
	}
}

func (model PortfolioCreateModel) IsValid() (bool, error) {
	return true, nil
}

//PortfolioID - portfolio unique identifier
type PortfolioID struct {
	ID string `json:"id" required:"true" description:"Portfolio ID"`
}
