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
var clientMap map[string]*mongo.Collection = make(map[string]*mongo.Collection)

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
func Destroy() {
}
func Add_one[T any](target_class T, db_name string, coll_name string) (*mongo.InsertOneResult, error) {

	temp, err := bson.Marshal(target_class)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	res, err_o := find_cre_handle(db_name, coll_name).InsertOne(context.TODO(), temp)
	if err_o != nil {
		log.Error(err_o.Error())
		return nil, err_o
	}
	return res, err_o
}

func Find_one[T any](db_name string, coll_name string, fieldType string, fieldValue string) (T, bool) {
	fielter := bson.D{{fieldType, fieldValue}}
	var res T
	err := find_cre_handle(db_name, coll_name).FindOne(context.TODO(), fielter).Decode(&res)
	if err != nil {
		log.Error(err.Error())
		return res, false
	} else {
		return res, true
	}
}
func find_cre_handle(db_name string, coll_name string) *mongo.Collection {
	h, contain := clientMap[db_name+coll_name]
	var handle_coll *mongo.Collection
	if contain {
		return h
	} else {
		err_p := MongoClient.Ping(context.TODO(), nil)
		if err_p != nil {
			log.Error(err_p.Error())
		}
		handle_coll = MongoClient.Database(db_name).Collection(coll_name)
		clientMap[db_name+coll_name] = handle_coll
	}
	return handle_coll
}
