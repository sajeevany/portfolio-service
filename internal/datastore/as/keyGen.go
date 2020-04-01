package as

import (
	"github.com/aerospike/aerospike-client-go"
	"github.com/google/uuid"
	"github.com/sajeevany/portfolio-service/internal/config"
	"github.com/sirupsen/logrus"
)

//GetUniqueID - returns a unique key that is not in use as a primary key in the specified namespace and set
func GetUniqueID(logger *logrus.Logger, client *aerospike.Client, setMetadata config.SetMD) (string, error) {

	//TESTING - Won't be unit tested as method is primarily composed to uuid generation and aerospike client code

	var id string
	var exists = true
	var err error

	logger.Debug("Generating unique key")

	//While no ID exists or the given ID does not exists
	loopCounter := 0
	for id == "" || exists {

		//Counter number of times that we've looped. If it exceeds 10, then
		loopCounter++

		id = uuid.New().String()
		logger.Debugf("Generating unique key %v", id)

		//Create aerospike key to check
		key, keyError := aerospike.NewKey(setMetadata.Namespace, setMetadata.SetName, id)
		if keyError != nil {
			logger.Error("Unexpected error when creating new key ")
			return "", keyError
		}

		//Check if key exists
		exists, err = client.Exists(aerospike.NewPolicy(), key)
		if err != nil {
			logger.Error("Error check if key exists", err)
			return "", err
		}
		logger.Debugf("key: %v exists:%v", key, exists)
	}

	return id, err
}
