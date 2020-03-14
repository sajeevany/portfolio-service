package logging

import (
	"fmt"
	"github.com/sajeevany/portfolio-service/internal/config"
	"github.com/sirupsen/logrus"
)

func Init() *logrus.Logger {
	return logrus.New()
}

func Update(logger *logrus.Logger, config config.LoggingConf) error{

	//Update level
	lvl, err := logrus.ParseLevel(config.Level)
	if err != nil{
		msg := fmt.Sprintf("Unable to parse level <%v> into logrus.Level", config.Level)
		logger.Warn(msg)
		return err
	}
	logger.SetLevel(lvl)

	return nil
}