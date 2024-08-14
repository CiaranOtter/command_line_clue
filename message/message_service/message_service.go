package message_service

import (
	"clc_services/message"
	"context"
	"database/sql"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MessageServer struct {
	message.UnimplementedMessageServiceServer
	DB *sql.DB
}

func (m MessageServer) SendMessage(ctx context.Context, in *message.Message) (*message.Reply, error) {

	log.Printf("New send message request")

	res, err := m.DB.Exec("INSERT INTO messges (message, sender_id) VALUES ($1, SELECT id FROM account WHERE username = $2)", in.GetMessage(), in.GetUsername())

	if err != nil {
		log.Printf("Failed to send message", "error", err)
		return nil, status.Error(codes.Canceled, err.Error())
	}

	c, err := res.RowsAffected()

	if err != nil || c < 1 {
		log.Printf("Failed to send message", "error", err)
		return nil, status.Errorf(codes.Canceled, "%d rows affected. %w", err)
	}

	return &message.Reply{
		Success: true,
	}, nil
}

func (m MessageServer) ReceiveMessages(in *message.JoinChat, stream grpc.ServerStreamingServer[message.ReceiveMessage]) error {
	return nil
}
