package services

import (
	"context"
	"fmt"
	"log"
	"mongodb/config"
	model "mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = config.UserCollection

// insert 1 record
func CreateOneUser(user model.UserModel) {
	inserted, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("user created succesfully with id: ", inserted.InsertedID)
}

func UpdateAUser(userId string,activeStat bool) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{
		"_id": id,
	}
	update := bson.M{"$set": bson.M{"isActive":activeStat }}
	result,err := collection.UpdateOne(context.Background(),filter,update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("user updated modified count: ",result.ModifiedCount)

}

func DeleteOneUser(userId string){
	id,err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	result,err := collection.DeleteOne(context.Background(),filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("user deleted count: ",result.DeletedCount)

}

func DeleteAllUser() int64{

	result,err := collection.DeleteMany(context.Background(),bson.D{{}},nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("users deleted  count: ",result.DeletedCount)
	return result.DeletedCount
}


func GetAllUsers() []primitive.M {
	result,err := collection.Find(context.Background(),bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var users []primitive.M
	for result.Next(context.Background()){
		var user bson.M
		err := result.Decode(&user);
		if err != nil{
			log.Fatal(err)
		}
		users = append(users,user)
	}

	defer result.Close(context.Background())

	return users;

}
