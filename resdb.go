package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {

	connectionURI := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to MongoDB!")

	return client, ctx, cancel
}

func Create(rest *resturant) (primitive.ObjectID, error) {
	// rest.ID = primitive.NewObjectID()
	client, context, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(context)
	result, err := client.Database(DATABASE).Collection(COLLECTION).InsertOne(context, rest)
	if err != nil {
		log.Printf("Could not create Resturant: %v", err)
		return primitive.NilObjectID, err
	}
	oid := (result.InsertedID.(primitive.ObjectID))
	return oid, nil
}

func Get() ([]*resturant, error) {
	var restaurants []*resturant
	client, context, cancel := getConnection()
	filter := bson.M{}
	defer cancel()
	defer client.Disconnect(context)
	cur, err := client.Database(DATABASE).Collection(COLLECTION).Find(context, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context)
	err = cur.All(context, &restaurants)
	if err != nil {
		log.Printf("Failed while marshaling : %v", err)
		return nil, err
	}
	return restaurants, nil
}

func GetById(id string) (*resturant, error) {
	var rest resturant
	objectid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Object id could not be created")
		return nil, err
	}
	client, context, cancel := getConnection()
	filter := bson.M{"_id": objectid}
	defer cancel()
	defer client.Disconnect(context)
	err = client.Database(DATABASE).Collection(COLLECTION).FindOne(context, filter).Decode(&rest)
	if err != nil {
		return nil, errors.New("could not find a task")
	}
	return &rest, nil

}

func Update(rest *resturant) (*resturant, error) {
	var updatedrest resturant
	client, context, cancel := getConnection()
	update := bson.M{"$set": rest}
	filter := bson.M{"_id": rest.ID}
	defer cancel()
	defer client.Disconnect(context)
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &after,
	}
	err := client.Database(DATABASE).Collection(COLLECTION).FindOneAndUpdate(context, filter, update, &opts).Decode(&updatedrest)
	if err != nil {
		log.Printf("Could not update details : %v", err)
		return nil, err
	}
	return &updatedrest, nil
}

func Delete(id string) (int64, error) {
	client, context, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(context)
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error when objectid:%v", err.Error())
		return 0, err
	}
	filter := bson.M{"_id": objid}
	result, err := client.Database(DATABASE).Collection(COLLECTION).DeleteOne(context, filter)
	if err != nil {
		log.Printf("err when query:%v", err)
		return 0, err
	}
	fmt.Printf("result: %v\n", result)
	return result.DeletedCount, nil
}
