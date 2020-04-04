package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/portfolio-service/internal/datastore"
	"github.com/sirupsen/logrus"
	"net/http"
)

//@Summary Get portfolios endpoint
//@Description Non-authenticated endpoint that returns array of all stored portfolios.
//@Produce json
//@Success 200 {object} model.PortfolioViewModel
//@Failure 404 {string} model.Error
//@Router /portfolio/{:id} [get]
//@Tags portfolio
func GetPortfolio(logger *logrus.Logger, asClient *datastore.ASClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//var portfolios model.AllPortfoliosViewModel
		//
		//err := asClient.Client.GetObject()

		////Default response for testing
		//response := model.AllPortfoliosViewModel{
		//	Portfolios: []model.PortfolioViewModel{{
		//		Metadata: model.MetadataViewModel{
		//			ID: "1", CreateTime: "2", LastUpdated: "3"},
		//		Stocks: []model.StockViewModel{}}},
		//}
		ctx.JSON(http.StatusOK, "responseresponse")
	}
}
