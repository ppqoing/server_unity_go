package mongodb

import (
	"context"
	"fmt"

	"github.com/name5566/leaf/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

const (
	Database_URL           string = "mongodb://localhost:27017"
	Database_golang_server string = "golang_server"
	Collection_login       string = "Login"
)

func Link() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Database_URL))
	if err != nil {
		log.Error(err.Error())
		fmt.Println("mongo link err")
	}
	MongoClient = client
}
func Add_one[T any](target_class T, db_name string, coll_name string) (*mongo.InsertOneResult, error) {
	temp, err := bson.Marshal(target_class)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	err_p := MongoClient.Ping(context.TODO(), nil)
	if err_p != nil {
		log.Error(err_p.Error())
		return nil, err_p
	}
	tar_client := MongoClient.Database(db_name).Collection(coll_name)
	res, err_o := tar_client.InsertOne(context.TODO(), temp)
	if err_o != nil {
		log.Error(err_o.Error())
		return nil, err_o
	}
	return res, err_o
}
