package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func db() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(getEnvValue("MONGO")))
	if err != nil {
		log.Panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Panic(err)
	} else {
		log.Println("Mongo connected!")
	}
}
