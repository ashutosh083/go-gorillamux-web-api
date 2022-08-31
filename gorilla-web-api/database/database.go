package database

import (
	"amazon/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://ashutosh99:ashu1122@ashuuser.v7twhlz.mongodb.net/test"
const dbName = "AmazonShows"
const colName = "shows"

//most important for connection
var Collection *mongo.Collection

//connect with data base
func init() {
	//client options
	clientOptions := options.Client().ApplyURI(connectionString)
	//connect to db
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[ INFO ] Database Connected Successfully!")
	Collection = client.Database(dbName).Collection(colName)
	//if collection refrence is ready
}

//helper methods of data base
func Create(movie model.Amazon) {
	_, err := Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
}

//update one record
func Update(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	_, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

//delete one movie
func Delete(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	_, err := Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
}

//delete all movies
func DeleteAll() {
	filter := bson.D{{}}
	_, err := Collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
}

//get all shows from database
func Find() []primitive.M {
	cur, err := Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
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
