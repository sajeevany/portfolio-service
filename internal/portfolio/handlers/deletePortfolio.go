package handlers

import (
	"fmt"
	"github.com/aerospike/aerospike-client-go"
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/portfolio-service/internal/datastore"
	"github.com/sajeevany/portfolio-service/pkg/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

//@Summary Deletes a portfolio at the specified ID
//@Description Deletes a portfolio with the specified ID. Returns 200 if the resource did not already exist.
//@Accept json
//@Param id path string true "Portfolio ID"
//@Produce json
//@Success 200 {string} string "ok"
//@Failure 404 {object} model.Error
//@Router /portfolio [delete]
//@Tags portfolio
func DeletePortfolioHandler(logger *logrus.Logger, asClient *datastore.ASClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Validate that id parameter has been set
		portfolioID := ctx.Param("id")
		if portfolioID == "" {
			msg := fmt.Sprintf("Query parameter %v hasn't been set", "id")
			logger.Debug(msg)
			ctx.JSON(http.StatusBadRequest, model.Error{Message: msg})
			return
		}

		//Attempt to delete record
		if err := DeletePortfolio(logger, asClient, portfolioID); err != nil{
			logger.Error(err)
			ctx.JSON(http.StatusInternalServerError, model.Error{Message:err.Error()})
			return
		}else{
			logger.Debug("Deletion operation passed without error")
			ctx.Status(http.StatusOK)
			return
		}
	}
}

//DeletePortfolio - Deletes record with provided key
func DeletePortfolio(logger *logrus.Logger, asClient *datastore.ASClient, portfolioID string) error{

	logger.Debugf("Starting deletion of record <%v>", portfolioID)

	//Build the key needed to delete
	key, err := aerospike.NewKey(asClient.SetMetadata.Namespace, asClient.SetMetadata.SetName, portfolioID)
	if err != nil {
		formattedErr:= fmt.Errorf("unexpected error <%v> when creating new key <%v> in delete operation", err, portfolioID)
		logger.Error(formattedErr.Error())
		return formattedErr
	}

	logger.Debugf("Key created for <%v>. Starting deletion request.", portfolioID)

	//Attempt to delete the entry. Don't pre-optimize by searching for the key because the aerospike client does that and returns a bool representing the answer.
	recordExisted, dErr := asClient.Client.Delete(asClient.WritePolicy, key)
	if dErr != nil{
		fmtErr := fmt.Errorf("encountered error <%v> when deleting portfolio with key <%v>", dErr, portfolioID)
		logger.Error(fmtErr.Error())
		return fmtErr

	}

	logger.Debugf("Deletion operation passed without error. Record was present to delete = <%v>", recordExisted)
	return nil
}