package datastore

import (
	"fmt"
	"github.com/aerospike/aerospike-client-go"
	"github.com/sajeevany/portfolio-service/internal/config"
	"github.com/sirupsen/logrus"
)

type ASClient struct {
	Client      *aerospike.Client
	WritePolicy *aerospike.WritePolicy
	ScanPolicy  *aerospike.ScanPolicy
	SetMetadata config.SetMD
}

func New(conf config.AerospikePortfolioConfig, logger *logrus.Logger) (*ASClient, error) {

	//Create aerospike client
	client, err := aerospike.NewClient(conf.Host, conf.Port)
	if err != nil {
		msg := fmt.Sprintf("Unexpected error when creating aerospike client, <%v> with config.", err)
		logger.WithFields(conf.GetFields()).Error(msg)
		return nil, err
	}
	logger.WithFields(conf.GetFields()).Info("Successful creation of aerospike client")

	//Create policies and define ASClient
	return &ASClient{
		Client:      client,
		WritePolicy: aerospike.NewWritePolicy(0, 0),
		ScanPolicy: aerospike.NewScanPolicy(),
		SetMetadata: conf.SetMD,
	}, nil
}
