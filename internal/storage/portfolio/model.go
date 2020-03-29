package portfolio

//StorageModel - defines how data is to be stored in persistent storage
type StorageModel struct {
	MD        MetadataStorageModel `json:"MetadataStorageModel"`
	Inventory InventoryModel       `json:Inventory`
}

type MetadataStorageModel struct {
	ID          string
	CreateTime  string
	LastUpdated string
}

type InventoryModel struct {
	Stocks map[string]StockStorageModel
}

type StockStorageModel struct {
	Ticker   string
	Quantity int
	Price    int
	Currency string
}
