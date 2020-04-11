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

//@Summary Creates portfolio a unique ID
//@Description Insert portfolio. Returns the portfolio ID.
//@Accept json
//@Param portfolio body model.PortfolioCreateModel true "Add account"
//@Produce json
//@Success 200 {object} model.PortfolioID
//@Failure 404 {string} model.Error
//@Router /portfolio [post]
//@Tags portfolio
func PostPortfolioHandler(logger *logrus.Logger, asClient *datastore.ASClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Bind body to portfolio object
		var portfolio model.PortfolioCreateModel
		if bErr := ctx.ShouldBindJSON(&portfolio); bErr != nil {
			msg := fmt.Sprintf("Unable to bind request body to portfolio object %v", bErr)
			logger.Errorf(msg)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}

		//Validate portfolio
		if valid, vErr := portfolio.IsValid(); (vErr != nil) || !valid {
			logger.WithFields(portfolio.GetFields()).Errorf("Input portfolio is invalid")
			ctx.JSON(http.StatusBadRequest, gin.H{"error": vErr})
			return
		}

		//Insert portfolio
		if id, insertErr := insertPortfolio(logger, asClient, portfolio); insertErr != nil{
			err := fmt.Errorf("error storing portfolio in aerospike <%v>", insertErr)
			logger.WithFields(portfolio.GetFields()).Error(err.Error())
			ctx.JSON(http.StatusInternalServerError, model.Error{Message: err.Error()})
			return
		}else {
			ctx.JSON(http.StatusOK, model.PortfolioID{ID:id})
			return
		}
	}
}

//insertRecord - Generates record for input portfolio and writes to aerospike with an unused key. Returns the generated ID. ID is empty if error occurs.
func insertPortfolio(logger *logrus.Logger, client *datastore.ASClient, portfolio model.PortfolioCreateModel)(string, error){

	logger.WithFields(portfolio.GetFields()).Debug("Starting record insertion for portfolio")

	//Get unused ID
	id, key, idErr := datastore.GetUniqueID(logger, client, client.SetMetadata)
	if idErr != nil {
		err := fmt.Errorf("unable to get unique id. Error <%v>", idErr)
		logger.WithFields(client.SetMetadata.GetFields()).Error(err.Error())
		return "", err
	}
	logger.Debugf("Generated ID %v", id)

	//Store in aerospike
	record := storage.NewRecord(portfolio, id)
	logger.WithFields(record.GetFields()).Debug("Created record. Starting put object operation into aerospike")
	if insertErr:= client.Client.PutObject(client.WritePolicy, key, record); insertErr != nil{
		err := fmt.Errorf("aerospike error putting portfolio object with key <%v>. Error <%v>", id, insertErr)
		logger.WithFields(record.GetFields()).Error(err)
		return "", err
	}

	return id, nil
}

