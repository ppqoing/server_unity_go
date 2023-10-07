package internal

import (
	"fmt"
	"reflect"
	"server/msg"

	"github.com/name5566/leaf/log"

	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/gate"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.Login{}, handleLogin)
	mongodb.Link()
}
func handleLogin(args []interface{}) {
	m := args[0].(*msg.Login)
	a := args[1].(gate.Agent)
	log.Debug("userId: %v \tpassword: %v\t userindex: %v", m.UserId, m.Password, m.UserIndex)
	res, err := mongodb.Add_one(m, mongodb.Database_golang_server, mongodb.Collection_login)
	if err != nil {
		log.Error(err.Error())
	}
	fmt.Println(res)
	a.WriteMsg(&msg.Login{
		UserId:    "ppqoing",
		Password:  "zjhzjh147",
		UserIndex: 1256321,
	})
}
