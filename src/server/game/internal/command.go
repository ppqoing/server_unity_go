package internal

import (
	"fmt"
)

func init() {
	skeleton.RegisterCommand("eccho", "echo user inputs", commandEcho)
}
func commandEcho(args []interface{}) interface{} {
	return fmt.Sprintf("%v", args)
}
