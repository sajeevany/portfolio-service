package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/portfolio-service/internal/datastore"
	"github.com/sajeevany/portfolio-service/internal/portfolio/storage"
	"github.com/sajeevany/portfolio-service/pkg/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

//@Summary Get portfolios endpoint
//@Description Non-authenticated endpoint that returns a portfolio with matching key.
//@Produce json
//@Param id path string true "Portfolio ID"
//@Success 200 {object} model.PortfolioViewModel
//@Failure 404 {object} model.Error
//@Router /portfolio/{id} [get]
//@Tags portfolio
func GetPortfolio(logger *logrus.Logger, asClient *datastore.ASClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Validate that id parameter has been set
		portfolioID := ctx.Param("id")
		if portfolioID == "" {
			msg := fmt.Sprintf("Query parameter %v hasn't been set", "id")
			logger.Debug(msg)
			ctx.JSON(http.StatusBadRequest, model.Error{Message: msg})
			return
		}

		//Check if key exists
		keyExists, key, err := datastore.KeyExists(logger, asClient, portfolioID)
		if err != nil{
			msg := fmt.Sprintf("Error <%v>. Unexpected internal error when checking if id <%v> exists", err, portfolioID)
			logger.Debug(msg)
			ctx.JSON(http.StatusInternalServerError, model.Error{Message: msg})
			return
		}

		if keyExists {
			//key exists so query aerospike and return record
			var r storage.Record
			if readErr := asClient.Client.GetObject(nil, key, &r); readErr != nil{
				msg := fmt.Sprintf("Error <%v>. Unable to read object for key <%v>", readErr, portfolioID)
				logger.Debug(msg)
				ctx.JSON(http.StatusBadRequest, model.Error{Message: msg})
				return
			}

			response := model.PortfolioViewModel{
				Metadata: r.Metadata,
				Stocks:   r.Inventory,
			}
			ctx.JSON(http.StatusOK, response)
			return
		}else{
			ctx.Status(http.StatusNotFound)
			return
		}
	}
}
