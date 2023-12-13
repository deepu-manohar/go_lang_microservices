package main

import (
	"context"
	"fmt"
	"log"
	"log/data"
	"log/logs"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	grpcPort = "50001"
)

var mongoClient *mongo.Client

type Config struct {
	LogRepo data.LogRepo
}

func main() {
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic("error connecting to mongo ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	logRepo := data.New(mongoClient)
	app := Config{
		LogRepo: logRepo,
	}
	rpcServer := NewRPCServer(logRepo)
	err = rpc.Register(rpcServer)
	if err != nil {
		log.Println("error in registering rpc server", err)
	}
	go app.startGRPC(logRepo)
	go app.startRPC()
	app.Serve()
}

func (app *Config) startGRPC(logRepo data.LogRepo) {
	log.Println("starting grpc server")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Panic("failed to start grpc server", err)
	}
	server := grpc.NewServer()
	logGrpcServer := NewLogsGRPCServer(logRepo)
	logs.RegisterLogServiceServer(server, logGrpcServer)
	log.Println("started grpc server successfully")

	if err = server.Serve(listener); err != nil {
		log.Panic("failed to start grpc server", err)
	}
}

func (app *Config) startRPC() {
	log.Println("starting rpc server")
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		log.Panic("failed to start rpc server", err)
	}
	defer listener.Close()

	for {
		rpcConnection, err := listener.Accept()
		if err != nil {
			log.Println("error accepting request ", err)
		} else {
			go rpc.ServeConn(rpcConnection)
		}
	}
}

func (app *Config) Serve() {
	log.Println("starting server...")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic("failed to start server", err)
	}
}

func connectToMongo() (*mongo.Client, error) {
	log.Println("establishing connection to mongo...")
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})
	conn, error := mongo.Connect(context.TODO(), clientOptions)
	if error != nil {
		log.Println("failed to connect to mongo ", error)
		return nil, error
	}
	log.Println("connection to Mongo is established...")
	return conn, nil
}
