package storage

import (
	"github.com/sajeevany/portfolio-service/pkg/model"
	"github.com/sirupsen/logrus"
)

type Record struct {
	Inventory map[int]model.StockViewModel `json:"inventory"`
}

func (r Record) GetFields() logrus.Fields {
	return logrus.Fields{
		"inventory": r.Inventory,
	}
}


//NewRecord - creates a new record for insertion from the user defined model.PortfolioCreateModel
func NewRecord(createModel model.PortfolioCreateModel) Record{

	//Create key -> stock mapping. Required to be able to delete or replace entries
	stocks := make(map[int]model.StockViewModel, len(createModel.Stocks))
	for count, val := range createModel.Stocks{
		stocks[count] = val
	}

	return Record{Inventory:stocks}
}