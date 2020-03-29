package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type AerospikePortfolioConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	SetMD    SetMD  `json:"setMetadata"`
}

func (c AerospikePortfolioConfig) GetFields() logrus.Fields {
	return logrus.Fields{
		"host":            c.Host,
		"port":            c.Port,
		"password":        "Redacted",
		"datasetMetadata": c.SetMD.GetFields(),
	}
}

//Returns a list of key value pairs of invalid/missing arguments and the reason for their incorrectness
func (c AerospikePortfolioConfig) AddInvalidArgs(logger *logrus.Logger, tag string, invalidArgs map[string]string) {

	//Check if empty
	if c == (AerospikePortfolioConfig{}) {
		invalidArgs[tag] = "Aerospike datastore is empty"
	} else {

		if c.Host == "" {
			invalidArgs[prependTag(tag, "host")] = fmt.Sprintf("Value is empty/unset. <%v>", c.Host)
		}

		if c.Port <= 0 {
			invalidArgs[prependTag(tag, "port")] = fmt.Sprintf("Value is invalid. 0 and non-negative port values are not allowed. <%v>", c.Port)
		}

		//validate SetMetadata
		c.SetMD.AddInvalidArgs(logger, prependTag(tag, "setMetadata"), invalidArgs)
	}
}

func prependTag(tagPrefix, name string) string {
	return fmt.Sprintf("%v.%v", tagPrefix, name)
}

type SetMD struct {
	Namespace string `json:"namespace"`
	SetName   string `json:"set"`
}

func (set SetMD) GetFields() logrus.Fields {
	return logrus.Fields{
		"namespace": set.Namespace,
		"setName":   set.SetName,
	}
}

func (set SetMD) AddInvalidArgs(logger *logrus.Logger, tag string, invalidArgs map[string]string) {

	if set == (SetMD{}) {
		invalidArgs[tag] = "Aerospike set details are missing/empty"
	} else {

		if set.Namespace == "" {
			invalidArgs[prependTag(tag, "namespace")] = fmt.Sprintf("Value is empty/unset. <%v>", set.Namespace)
		}

		if set.SetName == "" {
			invalidArgs[prependTag(tag, "set")] = fmt.Sprintf("Value is empty/unset. <%v>", set.SetName)
		}
	}
}
