package main

import (
	"log"
	"log/data"
)

type RPCServer struct {
	LogRepo data.LogRepo
}

func NewRPCServer(logRepo data.LogRepo) *RPCServer {
	return &RPCServer{
		LogRepo: logRepo,
	}
}

type RPCRequest struct {
	Name string
	Data string
}

type RPCResponse struct {
	Result string
}

func (r *RPCServer) LogInfo(request RPCRequest, response *RPCResponse) error {
	err := r.LogRepo.Insert(data.LogEntry{
		Name: request.Name,
		Data: request.Data,
	})
	if err != nil {
		log.Println("failed to insert to mongo", err)
	}
	response.Result = "Success!!"
	return nil
}
