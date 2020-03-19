package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type User struct {
	Name     string
	Location Location
}
type Location struct {
	Longitude float64
	Latitude  float64
}

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://192.168.2.28:27017"))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	fmt.Println("mgo server success", " ---- >  mongodb://192.168.2.28:27017")
	dataBase := client.Database("test")
	userCollection := dataBase.Collection("user")
	/*	data := make([]interface{}, 7)
		data[0] = User{Name: "天府广场", Location: Location{
			Longitude: 104.072833,
			Latitude:  30.663422,
		}}
		data[1] = User{Name: "四川大剧院", Location: Location{
			Longitude: 104.074378,
			Latitude:  30.664804,
		}}
		data[2] = User{Name: "新华文轩", Location: Location{
			Longitude: 104.070084,
			Latitude:  30.664649,
		}}
		data[3] = User{Name: "手工茶", Location: Location{
			Longitude: 104.072402,
			Latitude:  30.664121,
		}}
		data[4] = User{Name: "宽窄巷子", Location: Location{
			Longitude: 104.059826,
			Latitude:  30.669883,
		}}
		data[5] = User{Name: "奶茶", Location: Location{
			Longitude: 104.06085,
			Latitude:  30.670054,
		}}
		data[6] = User{Name: "钓鱼台", Location: Location{
			Longitude: 104.058424,
			Latitude:  30.670737,
		}}
		insertMany, err := userCollection.InsertMany(context.Background(), data)
		if err != nil {
			panic(err)
		}
		fmt.Printf("InsertMany插入的消息ID:%v\n", insertMany.InsertedIDs)*/
	userCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bsonx.Doc{{"location", bsonx.String("2dsphere")}},
	})
	//userCollection.Indexes().DropOne(context.Background(), "location_2d")
}
