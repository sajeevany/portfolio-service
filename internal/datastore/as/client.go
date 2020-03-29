package as

import (
	"fmt"
	"github.com/aerospike/aerospike-client-go"
	"github.com/sajeevany/portfolio-service/internal/config"
	"github.com/sirupsen/logrus"
)

func New(conf config.AerospikePortfolioConfig, logger *logrus.Logger) (*aerospike.Client, error) {
	client, err := aerospike.NewClient(conf.Host, conf.Port)
	if err != nil {
		msg := fmt.Sprintf("Unexpected error when creating aerospike client, <%v> with config.", err)
		logger.WithFields(conf.GetFields()).Error(msg)
		return nil, err
	}

	logger.WithFields(conf.GetFields()).Info("Successful creation of aerospike client")
	return client, err
}
