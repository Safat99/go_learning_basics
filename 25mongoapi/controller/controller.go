package controller

import (
	"context"
	"fmt"
	"log"
	"mongoGo/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017/netflix"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to  mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("Collection reference is ready")

}

// mongodb helpers -- file

//inserting 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie db with id: ", inserted.InsertedID)
}
