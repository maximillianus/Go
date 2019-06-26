package main

import (
	"context"
	"fmt"
	"time"
)

import "go.mongodb.org/mongo-driver/bson"
import "go.mongodb.org/mongo-driver/mongo"

import "go.mongodb.org/mongo-driver/mongo/options"
import "go.mongodb.org/mongo-driver/mongo/readpref"

func main() {
	businessServiceMongoURIgo := "mongodb://username:password@iad-mongos2.objectrocket.com:15584/some-database"
	// localhost_mongo_uri := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(businessServiceMongoURIgo))
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Error when making connection")
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Cannot ping. Not connected")
		return
	}
	database := client.Database("some-database")

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := database.ListCollections(ctx, bson.D{})
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			fmt.Println("Error decoding result")
			return
		}
		fmt.Println(result["name"])
	}
	if err != nil {
		fmt.Println("Cannot list collections")
		return
	}

	fmt.Println("Connected. End of script")

	err = client.Disconnect(context.TODO())

	if err != nil {
		fmt.Println("Fatal error when closing conenction")
	}
	fmt.Println("Connection to MongoDB closed.")

}
