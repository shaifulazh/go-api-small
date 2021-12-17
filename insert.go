package main

import (
	"context"
	"fmt"

	"log"
	"time"

	// "go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"  binding:"required"`
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

func insert_some(c *gin.Context) {
	var album album
	err := c.BindJSON(&album)
	if err != nil {
		// log.Fatal(err)
		c.JSON(400, gin.H{"error": err.Error(), "status": "rejected"})
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := ConfigDB().Collection("posts")
	// docs := bson.D{{Key: "id_key", Value: 9},
	// 	{Key: "sadfasdfas", Value: "Nsadfasdfasdfaic Raboy"}}

	res, insertErr := collection.InsertOne(ctx, album)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	fmt.Println(res)
	c.JSON(200, album)
}
