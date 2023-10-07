package internal

import "reflect"

func init() {
	// handler(&msg.Hello{}, handlerHello)
}
func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)

}
func handlerHello(args []interface{}) {
	// m := args[0].(*msg.Hello)
	// a := args[1].(gate.Agent)
	// log.Debug("hell0 %v", m.Name)
	// log.Debug("id: %v", m.Id)
	// a.WriteMsg(&msg.Hello{
	// 	Name: "me",
	// 	Id:   12,
	// })
}
