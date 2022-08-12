package main

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster0.b6temx2.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("what-the-food").Collection("menu")
	name := "ส้มตำ"
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{Key: "name", Value: name}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the name %s\n", name)
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
