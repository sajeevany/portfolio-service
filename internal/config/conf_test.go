package config_test

import (
	"github.com/sajeevany/portfolio-service/internal/config"
	"github.com/sirupsen/logrus"
	"testing"
)

//Tests base config with 0 port value
func TestConf_GetInvalidArgs(t *testing.T) {

	t.Parallel()

	//Define params. Conf skips all nested structs, those structs should define tests to cover invalid arg testing
	conf := config.Conf{
		Port:        0,
		AerospikeDS: config.AerospikePortfolioConfig{},
		Logger:      config.LoggingConf{},
	}

	//Execute testing method
	invalidArgs := conf.GetInvalidArgs(logrus.New())

	//Verification
	//Expect non-empty invalid args map
	if len(invalidArgs) == 0 {
		t.Errorf("Expected number of args to non-zero")
	}

	//Expect port argument in invalid args map
	if _, ok := invalidArgs["cfg.port"]; !ok {
		t.Errorf("Expected to see cfg.port in map of invalid args")
	}
}
