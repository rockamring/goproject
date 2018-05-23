package module

import "github.com/rockamring/goproject/server/chanrpc"

type Skeleton struct {
	GoLen				int
	TimerDispatcherLen 	int
	AsynCallLen			int
	ChanRPCServer		*chanrpc.Server


	client				*chanrpc.Server
	server				*chanrpc.Server
	commandServer		*chanrpc.Server
}

