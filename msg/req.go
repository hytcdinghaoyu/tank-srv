package msg

func init() {
	// 这里我们注册了一个 JSON 消息 Hello
	Processor.Register(&Login{})
}

type Login struct {
	Name string
	UID  int
}

type JoinRoomReq struct {
	UID int
}
