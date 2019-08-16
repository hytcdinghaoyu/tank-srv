package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Login{})
	Processor.Register(&JoinRoom{})
}

//Login message
type Login struct {
	Name string
	UID  int
}

//JoinRoom message
type JoinRoom struct {
	UID int
}
