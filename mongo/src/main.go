package main

import (
	"context"

	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	insert()
	retrieve()
}

const (
	database   = "golang"
	collection = "users"
)

func retrieve() {
	db, client, ctx, err := Connect()
	defer client.Disconnect(ctx)
	if err != nil {
		log.Fatalln("Connection Error: " + err.Error())
	}
	coll := db.Collection(collection)
	//filter := bson.D{{Key: "hello", Value: "world"}}
	filter := bson.D{{Key: "davide"},{Key: "hello"}}
	cur, err := coll.Find(ctx, filter)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		log.Println("record does not exist")
	} else if err != nil {
		log.Fatal("Error in Find by filter: " + err.Error())
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(result.Map()["hello"])
		log.Println(result.Map()["davide"])
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}

func insert() {
	db, client, ctx, err := Connect()
	defer client.Disconnect(ctx)
	if err != nil {
		log.Fatalln("Connection Error: " + err.Error())
	}
	coll := db.Collection(collection)
	coll.InsertOne(context.Background(), bson.M{"hello": "world"})
	coll.InsertOne(context.Background(), bson.D{{Key: "davide", Value: "dinnocente"}})
	log.Println("Inserted")
}

func Connect() (*mongo.Database, *mongo.Client, context.Context, error) {
	log.Println("Connecting With Mongo")
	uri := "mongodb://mongoadmin:secret@localhost:27071/?timeoutMS=5000"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln("Connection Error: " + err.Error())
		return nil, nil, nil, err
	}
	log.Println("Client created")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln("Connection Error: " + err.Error())
		return nil, nil, nil, err
	}
	log.Println("Client connected")
	db := client.Database(database)
	log.Println("Database retrieved")
	return db, client, ctx, nil
}

/*
log.Panicln(err.Error()) //https://github.com/vitessio/vitess/issues/2842
log.Fatalln(err.Error())
*/
