package config_test

import (
	"fmt"
	"github.com/sajeevany/portfolio-service/internal/config"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestAerospikePortfolioConfig_AddInvalidArgs(t *testing.T) {

	t.Parallel()

	var tests = []struct {
		ExpectedInvalidArg string
		AsConfig           config.AerospikePortfolioConfig
	}{
		{"conf.aero", config.AerospikePortfolioConfig{}},                          //Not allowed to be empty
		{"conf.aero.host", config.AerospikePortfolioConfig{Host: "", Port: 1234}}, //Host is not allowed to be empty
		{"conf.aero.port", config.AerospikePortfolioConfig{Port: -8988}},          //Cannot be negative
		{"conf.aero.setMetadata", config.AerospikePortfolioConfig{
			Host:     "asbc.101.com",
			Port:     1234,
			Password: "qwetyu123",
			SetMD:    config.SetMD{},
		}}, //Cannot be empty. This is empty for all above scenarios. Emphasizing this for clarity
	}

	for _, scenario := range tests {
		//redeclare test to prevent parallel run problem
		scenario := scenario
		t.Run(fmt.Sprintf("Invalid argument detection AerospikePortfolioConfig %v", scenario.ExpectedInvalidArg), func(t *testing.T) {
			t.Parallel()

			invalidArgs := make(map[string]string)
			scenario.AsConfig.AddInvalidArgs(logrus.New(), "conf.aero", invalidArgs)

			//Failure if the expected argument does not exist in the set of invalid arguments
			if _, exists := invalidArgs[scenario.ExpectedInvalidArg]; !exists {
				t.Errorf("Could not find expected argument <%v> in map of invalid arguments <%v>", scenario.ExpectedInvalidArg, invalidArgs)
			}

		})
	}
}

func TestSetMD_AddInvalidArgs(t *testing.T) {

	t.Parallel()

	var tests = []struct {
		ExpectedInvalidArg string
		AsConfig           config.SetMD
	}{
		{"conf.aero.setMetadata", config.SetMD{}}, //Not allowed to be empty
		{"conf.aero.setMetadata.namespace", config.SetMD{
			SetName: "portfolio",
		}}, //Namespace cannot be empty
		{"conf.aero.setMetadata.set", config.SetMD{
			Namespace: "portfolio-service",
		}},
	}

	for _, scenario := range tests {
		//redeclare test to prevent parallel run problem
		scenario := scenario
		t.Run(fmt.Sprintf("Invalid argument detection SetMD %v", scenario.ExpectedInvalidArg), func(t *testing.T) {
			t.Parallel()

			invalidArgs := make(map[string]string)
			scenario.AsConfig.AddInvalidArgs(logrus.New(), "conf.aero.setMetadata", invalidArgs)

			//Failure if the expected argument does not exist in the set of invalid arguments
			if _, exists := invalidArgs[scenario.ExpectedInvalidArg]; !exists {
				t.Errorf("Could not find expected argument <%v> in map of invalid arguments <%v>", scenario.ExpectedInvalidArg, invalidArgs)
			}

		})
	}
}
