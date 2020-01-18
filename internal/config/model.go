package config

type Conf struct {
	//Base attributes
	Version     string    `json:"version"`
	Port        int       `json:"port"`
	Cache       DataStore `json:"cache"`
	PortfolioDS DataStore `json:"portfolioDS"`
	UserDS      DataStore `json:"userDS"`
}

type DataStore struct {
	Type     string `json:"type"`
	Address  string `json:"address"`
	Password string `json:"password"`
	DbName   string `json:"dbName"`
}
