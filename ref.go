package main

import (
	"command_line_clue/clue"
	comm "command_line_clue/clue/comm"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var Service *ClueService

type Client struct {
	Stream    *comm.ClueService_GameStreamServer
	Name      string
	Player    *clue.Player
	CardIndex []int32
}

func (s *ClueService) ShuffleAndDist() {
	index := make([]int, 0)

	for i := 0; i < len(s.Cards); i++ {
		index = append(index, i)
	}

	for i := range index {
		j := rand.Intn(i + 1)
		index[i], index[j] = index[j], index[i]
	}

	pI := 0

	for _, j := range index {
		s.clients[pI].CardIndex = append(s.clients[pI].CardIndex, int32(j))
		s.clients[pI].Player.Cards = append(s.clients[pI].Player.Cards, s.Cards[j])

		pI = (pI + 1) % s.player_Count
	}

	for _, cl := range s.clients {
		(*cl.Stream).Send(&comm.Message{
			Data: &comm.Message_Cards{
				Cards: &comm.Cards{
					Index: cl.CardIndex,
				},
			},
		})
	}
}

func (c *Client) SetChar(charName string) {
	for _, char := range Service.Characters {
		if strings.Compare(char.CharacterName, charName) == 0 {
			c.Player = clue.NewPlayer(c.Name, char, nil)
			fmt.Printf("Setting %s to character %s\n", c.Name, char.GetString())
			return
		}
	}
}

type ClueService struct {
	Players      []*clue.Player
	clients      []*Client
	Characters   []*clue.Character
	Rooms        []*clue.Room
	Weapons      []*clue.Weapon
	Cards        []clue.Card
	Answer       *clue.Answer
	running      bool
	ready        bool
	player_Count int
	RoomIndex    int
	CharIndex    int
	WeapIndex    int
	comm.UnimplementedClueServiceServer
}

func (s *ClueService) HandleConnect(connection *comm.Connect, stream *comm.ClueService_GameStreamServer) (*Client, error) {

	if strings.Compare(connection.GetPlayerName(), "") != 0 {
		client := &Client{
			Name:   connection.GetPlayerName(),
			Stream: stream,
			Player: nil,
		}

		s.clients = append(s.clients, client)

		fmt.Printf("%s has cnnected to the game.\n", connection.GetPlayerName())

		for _, cl := range s.clients {

			if strings.Compare(cl.Name, client.Name) == 0 {
				continue
			}

			(*cl.Stream).Send(&comm.Message{
				Data: &comm.Message_Opp{
					Opp: &comm.Opponent{
						Data: &comm.Opponent_Con{
							Con: &comm.Connect{
								PlayerName: client.Name,
							},
						},
					},
				},
			})

			(*client.Stream).Send(&comm.Message{
				Data: &comm.Message_Opp{
					Opp: &comm.Opponent{
						Data: &comm.Opponent_Con{
							Con: &comm.Connect{
								PlayerName: cl.Name,
							},
						},
					},
				},
			})

			if cl.Player != nil {
				(*client.Stream).Send(&comm.Message{
					Data: &comm.Message_Opp{
						Opp: &comm.Opponent{
							PlayerName: cl.Name,
							Data: &comm.Opponent_SetChar{
								SetChar: &comm.SetPlayer{
									CharacterName: cl.Player.Char.CharacterName,
								},
							},
						},
					},
				})
			}
		}

		(*stream).Send(&comm.Message{
			Data: &comm.Message_Com{
				&comm.Command{
					Type: comm.CommandType_PICK_CHAR,
				},
			},
		})

		(*stream).Send(&comm.Message{
			Data: &comm.Message_RAns{
				RAns: &comm.RemoveAnswer{
					Char: int32(s.CharIndex),
					Room: int32(s.RoomIndex),
					Weap: int32(s.WeapIndex),
				},
			},
		})

		return client, nil
	}

	fmt.Printf("Invalid\n")

	return nil, status.Error(codes.InvalidArgument, "Connection message is invalid")
}

func (s *ClueService) CheckReady() bool {

	if len(s.clients) < s.player_Count {
		return false
	}

	for _, c := range s.clients {
		if c.Player == nil {
			return false
		}
	}

	return true
}

func (s *ClueService) StartGame() {
	for _, cl := range s.clients {
		(*cl.Stream).Send(&comm.Message{
			Data: &comm.Message_Start{
				Start: &comm.StartGame{},
			},
		})
	}
}

func (s *ClueService) GameStream(stream comm.ClueService_GameStreamServer) error {
	var MyClient *Client

	for {
		mes, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		t := mes.GetData()

		switch t.(type) {
		case *comm.Message_Con:
			MyClient, err = s.HandleConnect(mes.GetCon(), &stream)
			if err != nil {
				return err
			}
			break
		case *comm.Message_SetChar:
			MyClient.SetChar(mes.GetSetChar().GetCharacterName())

			// update all the players of their character choices
			for _, cl := range s.clients {
				if strings.Compare(cl.Name, MyClient.Name) == 0 {
					continue
				}

				(*cl.Stream).Send(&comm.Message{
					Data: &comm.Message_Opp{
						Opp: &comm.Opponent{
							PlayerName: MyClient.Name,
							Data: &comm.Opponent_SetChar{
								SetChar: &comm.SetPlayer{
									CharacterName: MyClient.Player.Char.CharacterName,
								},
							},
						},
					},
				})
			}
			break
		default:
			fmt.Printf("Other type of message\n")
		}

		if s.CheckReady() {
			fmt.Printf("Game Read to Start")

			s.ShuffleAndDist()

			s.StartGame()
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "localhost", 5000))

	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	chars, room, weap := clue.ReadFile("data/characters.csv", "data/rooms.csv", "data/weapons.csv")

	i, j, k := clue.FindClues()

	Service = &ClueService{
		Players:      make([]*clue.Player, 0),
		clients:      make([]*Client, 0),
		Characters:   chars,
		Rooms:        room,
		Weapons:      weap,
		Cards:        clue.Cards,
		Answer:       clue.Solve,
		running:      false,
		ready:        false,
		RoomIndex:    j,
		CharIndex:    i,
		WeapIndex:    k,
		player_Count: 3,
	}

	comm.RegisterClueServiceServer(grpcServer, Service)

	fmt.Printf("Starting Server\n")
	grpcServer.Serve(lis)
}
