package endpoints

import (
	"github.com/gin-gonic/gin"
	portfolio "github.com/sajeevany/portfolio-service/internal/portfolio/handlers"
	"github.com/sirupsen/logrus"
)

const (
	PortfolioGroup = "/portfolio"
	GetAllPortfolios_= "/"
	GetPortfolio = "/:portfolioID"
	DeletePortfolio = "/:portfolioID"
	AddStock = "/:portfolioID/:ticker"
	ReplaceStockEntry = "/:portfolioID/:inventoryID"
	ReplaceStockEntries = "/:portfolioID"
	DeleteStockEntry = "/:portfolioID/:inventoryID"
	DeleteStockEntries = "/:portfolioID"
)

//BuildGetAllPortfoliosEndpoint - Returns all portfolios
func BuildGetAllPortfoliosEndpoint(logger *logrus.Logger, handlers ...gin.HandlerFunc) Endpoint {
	return Endpoint{
		URL:      GetAllPortfolios_,
		Handlers: append(handlers, portfolio.GetAllPortfolios(logger)),
	}
}