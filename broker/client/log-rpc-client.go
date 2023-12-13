package client

import (
	"broker/common"
	"log"
	"net/rpc"
)

type LogRPCClient struct {
	rpcClient *rpc.Client
}

type RPCRequest struct {
	Name string
	Data string
}

type RPCResponse struct {
	Result string
}

func NewLogRPCClient() LogRPCClient {
	log.Println("starting rpc client on 5001")
	client, err := rpc.Dial("tcp", "log:5001")
	if err != nil {
		log.Panic("failed to initialze rpc client ", err)
		return LogRPCClient{}
	}
	log.Println("started rpc client successfully")
	return LogRPCClient{
		rpcClient: client,
	}
}

func (l *LogRPCClient) SendLog(request common.BrokerLogRequest) error {
	var logRequest = RPCRequest{
		Name: request.Name,
		Data: request.Data + " from RPC",
	}
	var result = RPCResponse{}
	err := l.rpcClient.Call("RPCServer.LogInfo", logRequest, &result)
	if err != nil {
		log.Println("failed in rpc invokation ", err)
	} else {
		log.Println("RPC success!! ", result)
	}
	return nil
}
