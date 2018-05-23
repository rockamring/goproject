package module

import "github.com/rockamring/goproject/server/chanrpc"

type Skeleton struct {
	GoLen				int
	TimerDispatcherLen 	int
	AsynCallLen			int
	ChanRPCServer		*chanrpc.Server


	client				*chanrpc.Client
	server				*chanrpc.Server
	commandServer		*chanrpc.Server
}

func (s *Skeleton) AsynCall(server *chanrpc.Server, id interface{}, args ...interface{}) {
	if s.AsynCallLen == 0 {
		panic("invalid AsynCallLen")
	}

	s.client.Attach(server)// TODO
	s.client.AsynCall(id, args...)
}

func (s *Skeleton) RegisterChanRPC(id interface{}, f interface{}) {
	if s.ChanRPCServer == nil {
		panic("invalid ChanRPCServer")
	}

	s.server.Register(id, f)
}