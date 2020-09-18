package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	ctx := context.Background()
	uri := "mongodb://touch:secret@localhost:27017"

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Println("ERROS CONNECT")
		panic(err)
	}
	defer func() { _ = client.Disconnect(ctx) }()

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println("ERROS PING")
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("mydb").Collection("persons")

	ruan := Person{"Ruan", 35, "Cape Town"}

	insertResult, err := collection.InsertOne(context.TODO(), ruan)
	if err != nil {
		log.Println("ERRORS :")
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)

}
