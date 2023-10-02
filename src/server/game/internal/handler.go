package internal

import (
	"reflect"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {

}
func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)

}
func handlerHello(args []interface{}) {
	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)
	log.Debug("hell0 %v %v", m.Name, m.Id)
	a.WriteMsg(&msg.Hello{Name: "me", Id: 1})
}
