package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Login{})
}

type Login struct {
	UserId    string
	Password  string
	UserIndex int64
}
