package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/portfolio-service/internal/config"
	"github.com/sajeevany/portfolio-service/internal/datastore"
	"github.com/sajeevany/portfolio-service/internal/endpoints"
	"github.com/sajeevany/portfolio-service/internal/logging"
	lm "github.com/sajeevany/portfolio-service/internal/logging/middleware"
	"github.com/sirupsen/logrus"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/sajeevany/portfolio-service/docs"
)

const v1Api = "/api/v1"

//Build injection variables
var (
	GIT_COMMIT  string
	CONFIG_FILE string
)

// @title Portfolio Service API
// @version 1.0
// @description Stores and fetches user and model data
// @license.name MIT License
// @BasePath /api/v1
func main() {

	//Create a universal logger
	logger := logrus.New()

	//Log build values
	msg := fmt.Sprintf("Git commit %v", GIT_COMMIT)
	logger.Info(msg)

	//Read configuration file
	conf, err := config.Read(CONFIG_FILE, logger)
	if err != nil {
		//Log error and use default values returned
		logger.Error(err)
		return
	}

	//Validate configuration
	if invalidArgs := conf.GetInvalidArgs(logger); len(invalidArgs) != 0 {
		logger.WithFields(conf.GetFields()).Fatalf("Config has invalid/missing arguments: <%v>", invalidArgs)
	}
	logger.WithFields(conf.GetFields()).Info("Service config loaded")

	//Update logger with config args
	logger.Debug("Updating services based on service configuration")
	logging.Update(logger, conf.Logger)

	//Create Aerospike client
	asClient, err := datastore.New(conf.AerospikeDS, logger)
	if err != nil {
		logger.Fatal("Failed to create Aerospike client ", err)
	}

	//Initialize router
	router := setupRouter(logger)

	//Setup routes
	setupV1Routes(router, logger, asClient)

	//Add swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Use default route of 8080.
	port := fmt.Sprintf(":%d", conf.Port)
	routerErr := router.Run(port)
	if routerErr != nil {
		logger.Errorf("An error occurred when starting the router. <%v>", routerErr)
	}

}

//setupRouter - Create the router and set middleware
func setupRouter(logger *logrus.Logger) *gin.Engine {

	engine := gin.New()

	//Add middleware
	engine.Use(lm.LogRequest(logger))
	engine.Use(gin.Recovery())

	return engine
}

func setupV1Routes(rtr *gin.Engine, logger *logrus.Logger, asClient *datastore.ASClient) {
	addV1HealthEndpoints(rtr, logger)
	addV1UserEndpoints(rtr, logger)
	addV1PortfolioEndpoints(rtr, logger, asClient)
}

func addV1HealthEndpoints(rtr *gin.Engine, logger *logrus.Logger) {
	v1 := rtr.Group(fmt.Sprintf("%s%s", v1Api, endpoints.HealthGroup))
	{
		hello := endpoints.BuildHelloEndpoint(logger)
		v1.GET(hello.URL, hello.Handlers...)
	}
}

func addV1UserEndpoints(rtr *gin.Engine, logger *logrus.Logger) {
	v1 := rtr.Group(fmt.Sprintf("%s%s", v1Api, endpoints.UserGroup))
	{
		//GET user. Requires id.
		getUser := endpoints.BuildGetUserEndpoint(logger)
		v1.GET(getUser.URL, getUser.Handlers...)

		//GET all users
		getAllUsers := endpoints.BuildGetUsersEndpoint(logger)
		v1.GET(getAllUsers.URL, getAllUsers.Handlers...)

		//POST user
		postUser := endpoints.BuildAddUserEndpoint(logger)
		v1.POST(postUser.URL, postUser.Handlers...)
	}
}

func addV1PortfolioEndpoints(rtr *gin.Engine, logger *logrus.Logger, client *datastore.ASClient) {
	v1 := rtr.Group(fmt.Sprintf("%s%s", v1Api, endpoints.PortfolioGroup))
	{
		//GET portfolios
		getPortfolios := endpoints.BuildGetAllPortfoliosEndpoint(logger, client)
		v1.GET(getPortfolios.URL, getPortfolios.Handlers...)

		//GET portfolio
		getPortfolio := endpoints.BuildGetPortfolioEndpoint(logger, client)
		v1.GET(getPortfolio.URL, getPortfolio.Handlers...)

		//Post portfolio
		postPortfolio := endpoints.BuildPostPortfolioEndpoint(logger, client)
		v1.POST(postPortfolio.URL, postPortfolio.Handlers...)
	}
}
