// package main

// import (
// 	"encoding/binary"
// 	"encoding/json"
// 	"fmt"
// 	"net"
// 	"os"
// 	"github.com/name5566/leaf/db/mongodb"
// 	"gopkg.in/mgo.v2"
// )

// var mongoDB_whole *mgo.Session

// func main() {
// 	conn, err := net.Dial("tcp", "192.168.1.110:3563")
// 	if err != nil {
// 		panic(err)
// 	}
// 	temp := LoginJson{
// 		login{
// 			UserId:    "ppqoing",
// 			Password:  "zjhzjh147",
// 			UserIndex: 0,
// 		},
// 	}
// 	jsonStudent, err := json.Marshal(temp)
// 	if err != nil {
// 		fmt.Println("转换为json错误")
// 		os.Exit(-1)
// 	}
// 	data := []byte(string(jsonStudent))
// 	fmt.Println(string(jsonStudent))
// 	m := make([]byte, 2+len(data))
// 	binary.BigEndian.PutUint16(m, uint16(len(data)))
// 	copy(m[2:], data)
// 	conn.Write(m)
// 	fmt.Println("% x", m)
// 	conn.Close()

// }
// func mongodb(){
// mongoDB_c,err:=mongodb.Dial("localhost",27017)
// if err!=nil{
// 	fmt.Println(err)
// 	return
// }
// defer mongoDB_c.Close()
// s:=mongoDB_c.Ref()
// defer mongoDB_c.UnRef()

// }
//
//	type login struct {
//		UserId    string
//		Password  string
//		UserIndex int64
//	}
//
//	type LoginJson struct {
//		Login login
//	}
//
// mgotest project main.go
package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//userCollection := client.Database("golang_server").Collection("Login")
	// var temp Login
	// temp.Password = "pppp"
	// temp.UserId = "oooo"
	// temp.UserIndex = 9876675765

	// log_msg, err_m := json.Marshal(temp)
	// if err_m != nil {
	// 	fmt.Println(err_m)
	// }
	// fmt.Println(string(log_msg))
	// res, err_c := userCollection.InsertOne(context.TODO(), log_msg)
	// if err_c != nil {
	// 	panic(err_c)
	// }
	// fmt.Println(res.InsertedID)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err.Error())
	}
	handle_coll := client.Database("golang_server").Collection("Login")
	fielter := bson.D{{"password", "pppp"}}
	var res Login
	err_ccc := handle_coll.FindOne(context.TODO(), fielter).Decode(&res)
	if err_ccc != nil {
		fmt.Println(err_ccc.Error())
	} else {
		fmt.Println(res.UserId)
		fmt.Println(res.UserIndex)
	}

}

type Login struct {
	UserId    string
	UserIndex uint64
	Password  string
}
