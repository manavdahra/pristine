package providers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"os"
	"sync"
	"time"
)

var once sync.Once

type MongoClientConfig interface {
	GetMongoDbUri() string
	GetDbName() string
}

type MongoDAL struct {
	logger *zap.SugaredLogger
	Db     *mongo.Database
}

func InitMongoDBClient(config MongoClientConfig, logger *zap.SugaredLogger) *mongo.Database {
	var db *mongo.Database
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var err error
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetMongoDbUri()))
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}
		err = client.Ping(ctx, nil)
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}
		db = client.Database(config.GetDbName())
	})
	return db
}

func NewMongoDAL(config MongoClientConfig, logger *zap.SugaredLogger) *MongoDAL {
	return &MongoDAL{
		Db:     InitMongoDBClient(config, logger),
		logger: logger,
	}
}
