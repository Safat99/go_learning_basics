package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoGo/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// ----------------------mongodb helpers -- file
// they are all lowercased

// inserting 1 record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie db with id: ", inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieId string) int64 {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id} // bson.M vs bson.D
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("total records are modified: ", result.ModifiedCount)
	return result.ModifiedCount
}

// delete 1 record
func deleteOneMovie(movieId string) int64 {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted with delete count: ", deleteCount.DeletedCount)
	return deleteCount.DeletedCount
}

// delete all records from mongodb

func deleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("total number of movies deleted: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from database
// we will not find a value...we will find a cursor...we will have to
// loop through over that
func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M // represents bson documents as a map ...more reliable than bson.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

// Actual controllers -- file
// they are usually in this file and db helpers go in the seperate file

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allmovies := getAllMovies()
	json.NewEncoder(w).Encode(allmovies)
}

func GetAllMyMovies(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	allmovies := getAllMovies()
	c.JSON(http.StatusOK, allmovies)
}

func createMovie(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var movie model.Netflix
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}
	insertOneMovie(movie)
	c.JSON(201, gin.H{
		"message": "movie created",
	})

}

func UpdateMovie(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	id := c.Param("id")
	resultCount := updateOneMovie(id)

	c.JSON(201, gin.H{
		"message": "total" + strconv.Itoa(int(resultCount)) + " values updated",
	})

}

func DeleteMovie(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	id := c.Param("id")
	deleteCount := deleteOneMovie(id)

	c.JSON(201, gin.H{
		"message": "total" + strconv.Itoa(int(deleteCount)) + " values deleted",
	})
}
