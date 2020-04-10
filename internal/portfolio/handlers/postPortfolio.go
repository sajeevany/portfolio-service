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
func PostPortfolioHandler(logger *logrus.Logger, client *datastore.ASClient) gin.HandlerFunc {
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

		//Get unused ID
		id, key, err := datastore.GetUniqueID(logger, client, client.SetMetadata)
		if err != nil {
			logger.WithFields(client.SetMetadata.GetFields()).Errorf("Unable to get unique id %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		logger.Debugf("Generated ID %v", id)

		//Store in aerospike
		record := storage.NewRecord(portfolio, id)
		if insertErr:= client.Client.PutObject(client.WritePolicy, key, record); insertErr != nil{
			msg := fmt.Sprintf("Error <%v> inserting record with key <%v>", insertErr, id)
			logger.WithFields(record.GetFields()).Errorf(msg)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}

		//Insert operation completed. Format response and return
		response := model.PortfolioID{
			ID: id,
		}
		ctx.JSON(http.StatusOK, response)
	}
}
