package internal

import (
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
	handleMsg(&msg.Register{}, handleRegister)
	mongodb.Link()
}
func handleLogin(args []interface{}) {
	m := args[0].(*msg.Login)
	a := args[1].(gate.Agent)
	find_res, contain := mongodb.Find_one[msg.Login](mongodb.Database_golang_server, mongodb.Collection_login, "userid", m.UserId)
	if contain {
		if m.Password == find_res.Password {
			a.WriteMsg(&msg.LoginBack{
				LogResault:   true,
				LoginResault: "Success",
			})
		} else {
			a.WriteMsg(&msg.LoginBack{
				LogResault:   false,
				LoginResault: "Password wrong",
			})
		}

	} else {
		_, err := mongodb.Add_one(m, mongodb.Database_golang_server, mongodb.Collection_login)
		if err != nil {
			log.Error(err.Error())
		}
		a.WriteMsg(&msg.LoginBack{
			LogResault:   false,
			LoginResault: "Dont have this user",
		})
	}

}
func handleRegister(args []interface{}) {
	r := args[0].(*msg.Register)
	a := args[1].(gate.Agent)
	_, contain := mongodb.Find_one[msg.Login](mongodb.Database_golang_server, mongodb.Collection_login, "userid", r.UserId)
	if contain {
		a.WriteMsg(&msg.RegisterBack{
			RegisterRes: false,
			RegisterMsg: "UserId Registed",
		})

	} else {
		_, err := mongodb.Add_one(r, mongodb.Database_golang_server, mongodb.Collection_login)
		if err != nil {
			log.Error(err.Error())
		} else {
			a.WriteMsg(&msg.RegisterBack{
				RegisterRes: true,
				RegisterMsg: "success",
			})
		}
	}
}
