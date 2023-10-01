package utils

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cli *mongo.Client

func CreateMongoClient() {

	var err error = nil

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	cli, err = mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	var result bson.M
	if err = cli.Database("admin").RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		log.Fatal(err)
	}
}

func GetMongoData() *mongo.Database {
	return cli.Database(os.Getenv("MONGO_DB"))
}

func CloseMongoClient() error {

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	return cli.Disconnect(ctx)
}
