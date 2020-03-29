package config

import "github.com/sirupsen/logrus"

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
		"password": "REDACTED",
		"dbName":   c.DbName,
	}
}
