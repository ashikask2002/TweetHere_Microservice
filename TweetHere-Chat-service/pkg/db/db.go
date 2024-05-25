package db

import (
	"context"
	"fmt"
	"tweethere-chat/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDatabase(cfg config.Config) (*mongo.Database, error) {
	ctx := context.TODO()
	mongoConn := options.Client().ApplyURI(cfg.DBUrl)
	mongoclient, err := mongo.Connect(ctx, mongoConn)
	if err != nil {
		return nil, err
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	fmt.Println("mongo connection established")
	return mongoclient.Database(cfg.DBname), nil
}
