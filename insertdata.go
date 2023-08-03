package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type address struct {
	Street  string `bson:"street" json:"street"`
	City    string `bson:"city" json:"city"`
	State   string `bson:"state" json:"state"`
	Country string `bson:"country" json:"country"`
}

type user struct {
	UserID  int64    `bson:"user_id" json:"user_id"`
	Name    string   `bson:"name" json:"name"`
	Phone   string   `bson:"phone" json:"phone"`
	Address address  `bson:"address" json:"address"`
	Hobbies []string `bson:"hobbies" json:"hobbies"`
}

type httpresponse struct {
	Message string `json:"message"`
}

func insertUser(w http.ResponseWriter, r *http.Request) {

	User, err := validateInsertRequest(r)
	if err != nil {
		response := httpresponse{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)

		return
	}

	client, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	countersCollection := client.Database("testdb").Collection("counters")

	filter := bson.M{"_id": "user_id"}

	update := bson.M{"$inc": bson.M{"seq": 1}}
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)

	result := countersCollection.FindOneAndUpdate(context.Background(), filter, update, opt)

	if result.Err() != nil {
		fmt.Println("Findoneandupdate")
		log.Println(result.Err())
	}

	var counterDoc struct {
		Seq int64 `bson:"seq"`
	}
	err = result.Decode(&counterDoc)
	if err != nil {
		message := err
		jmsg, _ := json.Marshal(message)
		fmt.Println(string(jmsg))
	}

	User.UserID = counterDoc.Seq

	collection := client.Database("testdb").Collection("users")

	_, err = collection.InsertOne(context.TODO(), User)
	if err != nil {
		response := httpresponse{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)

		return
	}

	w.WriteHeader(http.StatusCreated)

	response := httpresponse{
		Message: "user inserted successfully",
	}
	json.NewEncoder(w).Encode(response)
}
