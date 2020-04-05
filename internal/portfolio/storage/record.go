package storage

import (
	"github.com/sajeevany/portfolio-service/pkg/model"
	"github.com/sirupsen/logrus"
	"time"
)

type Record struct {
	Metadata  model.MetadataViewModel      `json:"metadata"`
	Inventory map[int]model.StockViewModel `json:"inventory"`
}

func (r Record) GetFields() logrus.Fields {
	return logrus.Fields{
		"inventory": r.Inventory,
	}
}

//NewRecord - creates a new record for insertion from the user defined model.PortfolioCreateModel
func NewRecord(createModel model.PortfolioCreateModel, id string) Record {

	//Create key -> stock mapping. Required to be able to delete or replace entries
	stocks := make(map[int]model.StockViewModel, len(createModel.Stocks))
	for count, val := range createModel.Stocks {
		stocks[count] = val
	}

	t := time.Now()
	return Record{Metadata: model.MetadataViewModel{
		ID:          id,
		CreateTime:  t.Format("20060102150405"),
		LastUpdated: t.Format("20060102150405"),
	}, Inventory: stocks}
}
