package message_service

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"github.com/CiaranOtter/command_line_clue/server/clc_services/message"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MessageServer struct {
	message.UnimplementedMessageServiceServer
	DB           *sql.DB
	Sender_chans map[string](chan *message.ReceiveMessage)
}

func (m MessageServer) Broadcast(msg *message.Message) error {
	for name, client := range m.Sender_chans {
		if strings.Compare(name, msg.GetUsername()) == 0 {
			continue
		}
		log.Printf("Send %s to %s", msg.GetMessage(), name)
		client <- &message.ReceiveMessage{
			Username: msg.GetUsername(),
			Message:  msg.GetMessage(),
		}
	}

	return nil
}

func (m MessageServer) SendMessage(ctx context.Context, in *message.Message) (*message.Reply, error) {

	log.Printf("New send message request")

	res, err := m.DB.Exec("INSERT INTO messages (message, sender_id) VALUES ($1, (SELECT id FROM account WHERE username = $2))", in.GetMessage(), in.GetUsername())

	if err != nil {
		log.Printf("Failed to send message", "error", err)
		return nil, status.Error(codes.Canceled, err.Error())
	}

	c, err := res.RowsAffected()

	if err != nil || c < 1 {
		log.Printf("Failed to send message", "error", err)
		return nil, status.Errorf(codes.Canceled, "%d rows affected. %w", err)
	}

	m.Broadcast(in)
	return &message.Reply{
		Success: true,
	}, nil
}

func (m *MessageServer) ReceiveMessages(in *message.JoinChat, stream message.MessageService_ReceiveMessagesServer) error {

	messageChannel := make(chan *message.ReceiveMessage, 10)

	m.Sender_chans[in.GetUsername()] = messageChannel

	defer func() {
		close(m.Sender_chans[in.GetUsername()])
		delete(m.Sender_chans, in.GetUsername())
	}()
	rows, err := m.DB.Query("SELECT account.username as sender, messages.message FROM messages LEFT JOIN account ON messages.sender_id = account.id")

	if err != nil {
		log.Printf("Failed to fetch the messages from the database")
		return status.Error(codes.Aborted, "Failed to fetch message from the database")
	}
	defer rows.Close()

	for rows.Next() {
		mes := message.ReceiveMessage{}

		err := rows.Scan(&mes.Username, &mes.Message)

		if err != nil {
			log.Printf("failed to scan the mesasge data from the database")
			return status.Error(codes.Canceled, "Failed to read message data from the database")
		}

		err = stream.Send(&mes)

		if err != nil {
			log.Printf("Failed to send the message data to the client")
			return status.Error(codes.Canceled, "Failed to send the message across the client connection")
		}

		log.Printf("Message has been sent to %s", in.GetUsername())
	}

	for {
		select {
		case <-stream.Context().Done():
			log.Printf("Client has left the chat")
			return nil
		case msg := <-messageChannel:
			err := stream.Send(msg)

			if err != nil {
				log.Printf("Failed to send a message to the client")
				return status.Error(codes.Canceled, "Failed to send a message to the client")
			}
		}
	}
}
