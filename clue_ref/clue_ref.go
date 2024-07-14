package clue_ref

import "command_line_clue/clue/comm"

type Client struct {
	Stream *comm.ClueService_GameStreamServer
	Name   *comm.ClueServiceClient
}
