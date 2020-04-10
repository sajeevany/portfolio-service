package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/portfolio-service/internal/datastore"
	portfolioHandlers "github.com/sajeevany/portfolio-service/internal/portfolio/handlers"
	"github.com/sirupsen/logrus"
)

const (
	PortfolioGroup      = "/portfolio"
	GetAllPortfolios   = "/"
	PostPortfolio       = "/"
	GetPortfolio        = "/:id"
	DeletePortfolio     = "/:id"
	AddStock            = "/{id}/{tickerID}"
	ReplaceStockEntry   = "/{id}/{inventoryID}"
	ReplaceStockEntries = "/{id}"
	DeleteStockEntry    = "/{id}/{inventoryID}"
	DeleteStockEntries  = "/{id}"
)

//BuildGetAllPortfoliosEndpoint - Returns all portfolios
//swagger:model
func BuildGetAllPortfoliosEndpoint(logger *logrus.Logger, asClient *datastore.ASClient, handlers ...gin.HandlerFunc) Endpoint {
	return Endpoint{
		URL:      GetAllPortfolios,
		Handlers: append(handlers, portfolioHandlers.GetAllPortfoliosHandler(logger, asClient)),
	}
}

//BuildGetPortfolioEndpoint - Return portfolio with matching key
//swagger:model
func BuildGetPortfolioEndpoint(logger *logrus.Logger, asClient *datastore.ASClient, handlers ...gin.HandlerFunc) Endpoint {
	return Endpoint{
		URL:      GetPortfolio,
		Handlers: append(handlers, portfolioHandlers.GetPortfolioHandler(logger, asClient)),
	}
}

//BuildPostPortfolioEndpoint - Inserts new portfolio with a unique key
//swagger:model
func BuildPostPortfolioEndpoint(logger *logrus.Logger, asClient *datastore.ASClient, handlers ...gin.HandlerFunc) Endpoint {
	return Endpoint{
		URL:      PostPortfolio,
		Handlers: append(handlers, portfolioHandlers.PostPortfolioHandler(logger, asClient)),
	}
}

//BuildDeletePortfolioEndpoint - Deletes portfolio with specified id
func BuildDeletePortfolioEndpoint(logger *logrus.Logger, asClient *datastore.ASClient, handlers ...gin.HandlerFunc) Endpoint {
	return Endpoint{
		URL:      DeletePortfolio,
		Handlers: append(handlers, portfolioHandlers.DeletePortfolioHandler(logger, asClient)),
	}
}


