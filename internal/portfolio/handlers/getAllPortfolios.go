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
func GetAllPortfoliosHandler(logger *logrus.Logger, asClient *datastore.ASClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if portfolios, err := GetPortfolios(logger, asClient); err != nil{
			ctx.JSON(http.StatusInternalServerError, model.Error{Message: err.Error()})
			return
		}else{
			//Default response for testing
			response := model.AllPortfoliosViewModel{
				Portfolios: portfolios,
			}
			ctx.JSON(http.StatusOK, response)
		}
	}
}

func GetPortfolios(logger *logrus.Logger, asClient *datastore.ASClient) ([]model.PortfolioViewModel, error){

	logger.WithFields(asClient.SetMetadata.GetFields()).Debug("Scanning aerospike namespace")

	//Scan all objects and add to portfolio
	pChnl := make(chan *storage.Record)
	if _, scanErr := asClient.Client.ScanAllObjects(asClient.ScanPolicy, pChnl, asClient.SetMetadata.Namespace, asClient.SetMetadata.SetName); scanErr != nil{
		err := fmt.Errorf("aerospike error <%v> when scanning namespace <%v>", scanErr, asClient.SetMetadata.Namespace)
		logger.Error(err.Error())
		return nil, err
	}

	//Assuming that there's probably only 1 portfolio accessible by this user
	logger.Debug("Starting record conversion to portfolio view")
	var portfolios []model.PortfolioViewModel

	for r := range pChnl{
		logger.Debugf("Writing <%v> to response", r.Inventory)
		portfolios = append(portfolios, model.PortfolioViewModel{
			Metadata: r.Metadata,
			Stocks:   r.Inventory,
		})
	}

	return portfolios, nil
}
