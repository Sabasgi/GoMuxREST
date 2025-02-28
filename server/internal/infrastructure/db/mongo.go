package db

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

func ConnectMongo(dsn string) (*mongo.Client, error) {
	Client := options.Client().ApplyURI(dsn)
	c, err := mongo.Connect(context.Background(), Client)
	if err != nil {
		log.Println("DB Connecction : ", c)
		return c, err
	}
	return c, nil
}
