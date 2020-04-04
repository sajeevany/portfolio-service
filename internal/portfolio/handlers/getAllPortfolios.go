package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/portfolio-service/internal/datastore"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/sajeevany/portfolio-service/pkg/model"
)

//@Summary Get portfolios endpoint
//@Description Non-authenticated endpoint that returns array of all stored portfolios.
//@Produce json
//@Success 200 {object} model.AllPortfoliosViewModel
//@Failure 404 {string} model.Error
//@Router /portfolio [get]
//@Tags portfolio
func GetAllPortfolios(logger *logrus.Logger, asClient *datastore.ASClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Default response for testing
		response := model.AllPortfoliosViewModel{
			Portfolios: []model.PortfolioViewModel{{
				Metadata: model.MetadataViewModel{
					ID: "1", CreateTime: "2", LastUpdated: "3"},
				Stocks: []model.StockViewModel{}}},
		}
		ctx.JSON(http.StatusOK, response)
	}
}
