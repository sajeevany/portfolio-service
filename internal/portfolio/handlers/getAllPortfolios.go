package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/portfolio-service/internal/datastore"
	"github.com/sajeevany/portfolio-service/internal/portfolio/storage"
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

		logger.WithFields(asClient.SetMetadata.GetFields()).Debugf("Scanning aerospike namespace")

		//Scan all objects and add to portfolio
		pChnl := make(chan *storage.Record)
		_, err := asClient.Client.ScanAllObjects(asClient.ScanPolicy, pChnl, asClient.SetMetadata.Namespace, asClient.SetMetadata.SetName)
		if err != nil{
			msg := fmt.Sprintf("Error <%v> when scanning namespace <%v>", err, asClient.SetMetadata.Namespace)
			logger.Error(msg)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err })
			return
		}

		//Assuming that there's probably only 1 portfolio accessible by this user
		logger.Debug("Starting record conversion to portfolio view")
		portfolios :=  make([]model.PortfolioViewModel,1)

		for r := range pChnl{
			logger.Debugf("Writing <%v> to response", r.Inventory)
			portfolios = append(portfolios, model.PortfolioViewModel{
				Metadata: model.MetadataViewModel{},
				Stocks:   r.Inventory,
			})
		}

		////Default response for testing
		response := model.AllPortfoliosViewModel{
			Portfolios: portfolios,
		}
		ctx.JSON(http.StatusOK, response)
	}
}
