package config

import "github.com/sirupsen/logrus"

type Conf struct {
	//Base attributes
	VersionFile string      `json:"versionFile"`
	Port        int         `json:"port"`
	Cache       DataStore   `json:"cache"`
	PortfolioDS DataStore   `json:"portfolioDS"`
	Logger      LoggingConf `json:"loggingConf"`
}

func (c Conf) GetFields() logrus.Fields {
	return logrus.Fields{
		"version":         c.VersionFile,
		"port":            c.Port,
		"cache datastore": c.Cache.GetFields(),
		"model datastore": c.PortfolioDS.GetFields(),
	}
}

type DataStore struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DbName   string `json:"dbName"`
}

func (c DataStore) GetFields() logrus.Fields {
	return logrus.Fields{
		"type":     c.Type,
		"host":     c.Host,
		"port":     c.Port,
		"password": "Redacted",
		"dbName":   c.DbName,
	}
}

type LoggingConf struct {
	Level string `json:"level"`
}

func (c LoggingConf) GetFields() logrus.Fields {
	return logrus.Fields{
		"level": c.Level,
	}
}
