package main

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(u *user) (primitive.ObjectID, error) {
	client, context, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(context)
	result, err := client.Database(DATABASE).Collection(USER).InsertOne(context, u)
	if err != nil {
		log.Printf("could not create user: %v\n", err)
		return primitive.NilObjectID, err
	}
	log.Println("User Created successfully")
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

func GetUser(u *user) (*user, error) {
	var u1 user
	client, context, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(context)
	filter := bson.M{"email": u.Email, "password": u.Password}
	err := client.Database(DATABASE).Collection(USER).FindOne(context, filter).Decode(&u1)
	if err != nil {
		log.Printf("Could not find user :%v", err)
		return nil, err
	}
	return &u1, nil
}

func GetUserbyEmail(email string) (*user, error) {
	var u1 user
	client, context, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(context)
	filter := bson.M{"email": email}
	err := client.Database(DATABASE).Collection(USER).FindOne(context, filter).Decode(&u1)
	if err != nil {
		log.Printf("Could not find user :%v", err)
		return nil, err
	}
	return &u1, nil
}
