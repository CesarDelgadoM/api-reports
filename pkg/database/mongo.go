package database

import (
	"context"
	"fmt"

	"github.com/CesarDelgadoM/api-reports/config"
	"github.com/CesarDelgadoM/api-reports/pkg/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
	Config config.MongoConfig
}

func ConnectMongoDB(config config.MongoConfig) *MongoDB {
	uri := fmt.Sprintf(config.URI, config.User, config.Password)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		zap.Log.Fatal("Connect to mongo db failed: ", err)
	}
	zap.Log.Info("Connect to mongo db success")

	return &MongoDB{
		Client: client,
		Config: config,
	}
}

func (m *MongoDB) Disconnect() error {
	return m.Client.Disconnect(context.TODO())
}

func (m *MongoDB) CollectionRestaurant() *mongo.Collection {
	return m.Client.Database(m.Config.DBName).Collection("restaurant")
}
