package config

import "github.com/sirupsen/logrus"

type Conf struct {
	//Base attributes
	VersionFile string    `json:"versionFile"`
	Port        int       `json:"port"`
	Cache       DataStore `json:"cache"`
	PortfolioDS DataStore `json:"portfolioDS"`
}

type DataStore struct {
	Type     string `json:"type"`
	Address  string `json:"address"`
	Password string `json:"password"`
	DbName   string `json:"dbName"`
}

func (c Conf) GetFields() logrus.Fields {
	return logrus.Fields{"version": c.VersionFile, "port": c.Port}
}
