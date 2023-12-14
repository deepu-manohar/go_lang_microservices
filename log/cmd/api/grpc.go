package main

import (
	"context"
	"log"
	"log/data"
	"log/logs"
)

type LogGrpcServer struct {
	logs.UnimplementedLogServiceServer
	LogRepo data.LogRepo
}

func NewLogsGRPCServer(logRepo data.LogRepo) *LogGrpcServer {
	return &LogGrpcServer{LogRepo: logRepo}
}

func (l *LogGrpcServer) SendLog(ctx context.Context, request *logs.LogRequest) (*logs.LogResponse, error) {
	req := request.LogEntry
	log.Println("Got request in GRPC ", req)
	logEntry := data.LogEntry{
		Name: req.Name,
		Data: req.Data,
	}
	err := l.LogRepo.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: &logs.Result{Result: "failed"}}
		return res, err
	}
	res := &logs.LogResponse{Result: &logs.Result{Result: "success"}}
	return res, nil
}
