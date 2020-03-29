package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type LoggingConf struct {
	Level string `json:"level"`
}

func (c LoggingConf) GetFields() logrus.Fields {
	return logrus.Fields{
		"level": c.Level,
	}
}

//Returns a list of key value pairs of invalid/missing arguments and the reason for their incorrectness
func (c LoggingConf) AddInvalidArgs(logger *logrus.Logger, tag string, invalidArgs map[string]string) {

	//Check if empty
	if c == (LoggingConf{}) {
		invalidArgs[tag] = "Aerospike datastore is empty"
	} else {

		if c.Level == "" {
			invalidArgs[prependTag(tag, "level")] = fmt.Sprintf("Value is empty/unset. <%v>", c.Level)
		}
	}
}
