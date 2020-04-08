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
//@Success 200
//@Failure 404 {string} model.Error
//@Router /portfolio [delete]
//@Tags portfolio
func DeletePortfolio(logger *logrus.Logger, asClient *datastore.ASClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Validate that id parameter has been set
		portfolioID := ctx.Param("id")
		if portfolioID == "" {
			msg := fmt.Sprintf("Query parameter %v hasn't been set", "id")
			logger.Debug(msg)
			ctx.JSON(http.StatusBadRequest, model.Error{Message: msg})
			return
		}

		//Build the key needed to delete
		key, err := aerospike.NewKey(asClient.SetMetadata.Namespace, asClient.SetMetadata.SetName, portfolioID)
		if err != nil {
			msg:= fmt.Sprintf("Unexpected error when creating new key <%v>", portfolioID)
			logger.Error(msg)
			ctx.JSON(http.StatusInternalServerError, model.Error{Message:msg})
			return
		}

		//Attempt to delete the entry. Don't pre-optimize by searching for the key because the aerospike client does that and returns a bool representing the answer.
		if recordExisted, dErr := asClient.Client.Delete(asClient.WritePolicy, key); dErr != nil{
			msg := fmt.Sprintf("Error <%v> when deleting portfolio with key <%v>", dErr, portfolioID)
			logger.Error(msg)
			ctx.JSON(http.StatusInternalServerError, model.Error{Message:msg})
			return
		}else {
			logger.Debugf("Deletion operation passed without error. Record was present to delete = <%v>", recordExisted)
			ctx.Status(http.StatusOK)
			return
		}
	}
}
