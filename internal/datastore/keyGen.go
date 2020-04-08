package datastore

import (
	"github.com/aerospike/aerospike-client-go"
	"github.com/google/uuid"
	"github.com/sajeevany/portfolio-service/internal/config"
	"github.com/sirupsen/logrus"
)

//GetUniqueID - returns a unique key that is not in use as a primary key in the specified namespace and set
func GetUniqueID(logger *logrus.Logger, client *ASClient, setMetadata config.SetMD) (string, *aerospike.Key, error) {

	//TESTING - Won't be unit tested as method is primarily composed to uuid generation and aerospike client code

	var id string
	var exists = true
	var err error
	var key *aerospike.Key

	logger.Debug("Generating unique key")

	//While no ID exists or the given ID exists. ID must be non-empty and the aerospike key must be non-nil.
	//Set the key as exists as true to start to force do-while type flow
	loopCounter := 0
	for exists || id == "" || key == nil {

		//Counter number of times that we've looped. If it exceeds 10, then
		loopCounter++

		id = uuid.New().String()
		logger.Debugf("Generating unique key %v", id)

		exists, key, err = KeyExists(logger, client, id)
		if err != nil {
			logger.Error("Error when checking if key exists", err)
			return "", key, err
		}
	}

	return id, key, err
}

//KeyExists - Checks if the specified key exists. Returns result, corresponding key, and any error. Returns true in the event of an error
func KeyExists(logger *logrus.Logger, client *ASClient, id string)(bool, *aerospike.Key, error){

	//Create aerospike key to check
	key, err := aerospike.NewKey(client.SetMetadata.Namespace, client.SetMetadata.SetName, id)
	if err != nil {
		logger.Errorf("Unexpected error when creating new key <%v>", key)
		return true, key,  err
	}

	//Check if key exists. Use nil policy because no timeout is required
	exists, kerr := client.Client.Exists(nil, key)
	if kerr != nil {
		logger.Error("Error when checking if key exists", kerr)
		return true, key, kerr
	}
	logger.Debugf("key: %v exists:%v", key, exists)

	return exists, key, kerr
}