package main

import (
	"context"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
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
		panic(err)
	}
	defer func() { _ = client.Disconnect(ctx) }()

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("mydb").Collection("Person")

	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		
		muhlan := Person{"Muhlan", 21, "Bangkok"}
		insertResult, err := collection.InsertOne(context.TODO(), muhlan)
		if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
	})
		
	router.Run()
}
