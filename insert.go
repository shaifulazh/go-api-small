package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

func insert_some() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := ConfigDB().Collection("posts")
	docs := bson.D{{Key: "id_key", Value: 9},
		{Key: "sadfasdfas", Value: "Nsadfasdfasdfaic Raboy"}}

	res, insertErr := collection.InsertOne(ctx, docs)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	fmt.Println(res)
}
