package main

import (
	"context"
	"fmt"
	"net"
	"sync"

	"menssenger/server/protos"
	proto "menssenger/server/protos"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedChatServer
}

var clients []proto.Chat_GetMessagesServer

func (s server) SendMessage(context context.Context, message *proto.Message) (*proto.Void, error) {
	resp := proto.Message{Message: message.GetMessage(), User: message.GetUser()}

	for _, client := range clients {
		if err := client.Send(&resp); err != nil {
			fmt.Print(err)
		}
	}

	return &proto.Void{}, nil
}

func (s server) GetMessages(void *protos.Void, server proto.Chat_GetMessagesServer) error {
	clients = append(clients, server)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
	}()

	wg.Wait()

	return nil
}

func main() {
	fmt.Print("\033[H\033[2J")

	listener, err := net.Listen("tcp", ":4444")
	if err != nil {
		fmt.Print("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterChatServer(grpcServer, server{})

	if err := grpcServer.Serve(listener); err != nil {
		fmt.Print("failed to serve: %v", err)
	}
}
