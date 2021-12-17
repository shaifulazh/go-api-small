package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

type Post struct {
	Title int32  `bson:"id_key,omitempty"`
	Body  string `bson:"sadfasdfas,omitempty"`
}

func ConfigDB() *mongo.Database {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	return client.Database("blog")
}

func show() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := ConfigDB().Collection("posts")

	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var posts []Post
	if currErr = cur.All(ctx, &posts); currErr != nil {
		panic(currErr)
	}
	fmt.Println(posts)
}

func main() {
	// insert_some()
	show()
}
