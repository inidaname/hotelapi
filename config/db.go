package config

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var MI MongoInstance

func ConnectDB() {
	_ = mgm.SetDefaultConfig(nil, "hotels", options.Client().ApplyURI("mongodb://ocalhost:27017"))
}
