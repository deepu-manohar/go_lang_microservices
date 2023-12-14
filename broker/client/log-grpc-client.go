package client

import (
	"broker/client/logs"
	"broker/common"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

type LogGRPCClient struct {
	LogServiceClient logs.LogServiceClient
}

func NewLogGRPCClient(grpcConnection *grpc.ClientConn) *LogGRPCClient {
	return &LogGRPCClient{
		LogServiceClient: logs.NewLogServiceClient(grpcConnection),
	}
}

func (l *LogGRPCClient) SendLog(request common.BrokerLogRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logResponse, err := l.LogServiceClient.SendLog(ctx, &logs.LogRequest{
		LogEntry: &logs.Log{
			Name: request.Name,
			Data: request.Data + " from GRPC",
		},
	})
	if err != nil {
		log.Println("failed to make grpc call to send log ", err)
		return err
	} else {
		log.Println("Got response from grpc call %s", logResponse)
		return nil
	}
}
