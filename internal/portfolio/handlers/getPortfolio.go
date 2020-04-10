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
func GetPortfolioHandler(logger *logrus.Logger, asClient *datastore.ASClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Validate that id parameter has been set
		portfolioID := ctx.Param("id")
		if portfolioID == "" {
			msg := fmt.Sprintf("Query parameter %v hasn't been set", "id")
			logger.Debug(msg)
			ctx.JSON(http.StatusBadRequest, model.Error{Message: msg})
			return
		}

		//Get portfolio
		recordExists, portfolio, err := GetPortfolio(logger, asClient, portfolioID)
		if err != nil {
			logger.Errorf("unexpected error when fetching portfolio record for ID <%v>", portfolioID)
			ctx.JSON(http.StatusInternalServerError, model.Error{Message:err.Error()})
			return
		}

		//Shape response based on record existence
		if recordExists{
			logger.Debugf("Record exists for key <%v>. portfolio: <%v>", portfolioID, portfolio)
			ctx.JSON(http.StatusOK, portfolio)
		}else{
			ctx.Status(http.StatusNotFound)
			return
		}
	}
}

//GetPortfolio - Fetches portfolio with matching key. Returns key existence check, portfolio and error. Returns false whenever an error occurs and when a record is detected not to be found.
func GetPortfolio(logger *logrus.Logger, asClient *datastore.ASClient, portfolioID string) (bool, model.PortfolioViewModel, error){

	logger.Debugf("Starting portfolio lookup for <%v>", portfolioID)

	//Check if key exists
	keyExists, key, keyErr := datastore.KeyExists(logger, asClient, portfolioID)
	if keyErr != nil{
		err := fmt.Errorf("unexpected internal error when checking if id <%v> exists. error <%v> ", portfolioID, keyErr)
		logger.Debug(err)
		return false, model.PortfolioViewModel{}, err
	}
	logger.Debugf("Key built for id <%v>. Key existence check = <%v>", portfolioID, keyExists)

	//If the key exists then attempt to fetch the record
	if keyExists {
		var r storage.Record
		if readErr := asClient.Client.GetObject(nil, key, &r); readErr != nil{
			err := fmt.Errorf("unable to read object for key <%v>. Error <%v>", portfolioID, readErr)
			logger.Debug(err)
			return true, model.PortfolioViewModel{}, err
		}

		response := model.PortfolioViewModel{
			Metadata: r.Metadata,
			Stocks:   r.Inventory,
		}

		return true, response, nil
	}

	//No record matching input key. Return false for record existance check
	return false, model.PortfolioViewModel{}, nil
}