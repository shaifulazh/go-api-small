package main

import (
	"context"
	// "fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"

	"github.com/gin-gonic/gin"
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

func show() []album {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := ConfigDB().Collection("posts")

	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var albumw []album
	if currErr = cur.All(ctx, &albumw); currErr != nil {
		panic(currErr)
	}
	return albumw
}

type Result struct {
	Results []album
}

func index(c *gin.Context) {
	// object_json := Result{show()}
	c.JSON(200, gin.H{"result": show()})
}
func main() {
	// insert_some()
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/**/*")

	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/", index)
	router.POST("/post", insert_some)
	router.Run("localhost:8080")
}
