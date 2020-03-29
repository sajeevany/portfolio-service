package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type Conf struct {
	//Base attributes
	Port        int                      `json:"port"`
	AerospikeDS AerospikePortfolioConfig `json:"aerospikeDS"` //Aerospike data store
	Logger      LoggingConf              `json:"loggingConf"`
}

func (c Conf) GetFields() logrus.Fields {
	return logrus.Fields{
		"port":                c.Port,
		"portfolio datastore": c.AerospikeDS.GetFields(),
	}
}

type ArgValidator interface {
	AddInvalidArgs(logger *logrus.Logger, tagPrefix string, invalidArgs map[string]string)
}

//AddInvalidArgs - Returns a map of invalid keys and the corresponding problem.
func (c Conf) GetInvalidArgs(logger *logrus.Logger) map[string]string {

	//Opting to not use a fail first approach here. Don't want to
	invalidArgs := make(map[string]string)

	//Validate all local parameters
	if c.Port == 0 || c.Port < 0 {
		invalidArgs["cfg.port"] = fmt.Sprintf("Provided value is 0 or negative %v", c.Port)
	}

	//Validate AerospikeDS
	c.AerospikeDS.AddInvalidArgs(logger, "conf.aerospikeDS", invalidArgs)

	//Validate LoggingConf
	c.Logger.AddInvalidArgs(logger, "conf.loggingConf", invalidArgs)

	return invalidArgs
}
